package server

import (
	"fmt"
	"github.com/aaronchen2k/openstc/src/controllers"
	"github.com/aaronchen2k/openstc/src/libs/casbin"
	"github.com/aaronchen2k/openstc/src/libs/common"
	"github.com/aaronchen2k/openstc/src/libs/db"
	redisUtils "github.com/aaronchen2k/openstc/src/libs/redis"
	"github.com/aaronchen2k/openstc/src/models"
	"github.com/aaronchen2k/openstc/src/repo"
	"github.com/aaronchen2k/openstc/src/routes"
	"github.com/aaronchen2k/openstc/src/service"
	"github.com/kataras/iris/v12"
	"net/http"
	"strings"
	"time"

	"github.com/kataras/iris/v12/context"
)

type Server struct {
	App       *iris.Application
	AssetFile http.FileSystem
	baseRepo  *repo.BaseRepo
}

func NewServer(assetFile http.FileSystem) *Server {
	app := iris.Default()
	return &Server{
		App:       app,
		AssetFile: assetFile,
	}
}

func Init(version string, printVersion, seederData, syncPerms, printRouter *bool) {
	// gen objects
	permRepo := repo.NewPermRepo()
	roleRepo := repo.NewRoleRepo(permRepo)
	tokenRepo := repo.NewTokenRepo()
	userRepo := repo.NewUserRepo(tokenRepo)

	seederService := service.NewSeeder(userRepo, roleRepo, permRepo)

	accountCtrl := controllers.NewAccountController(userRepo, roleRepo, permRepo, tokenRepo)
	permCtrl := controllers.NewPermController(userRepo, permRepo)
	roleCtrl := controllers.NewRoleController(userRepo, roleRepo, permRepo)
	userCtrl := controllers.NewUserController(userRepo, roleRepo)

	// irisServer := server.NewServer(AssetFile()) // 加载前端文件
	irisServer := NewServer(nil)
	if irisServer == nil {
		panic("Http 初始化失败")
	}

	irisServer.App.Logger().SetLevel(common.Config.LogLevel)

	// init

	//if libs.Config.BinData {
	//	s.App.RegisterView(iris.HTML(s.AssetFile, ".html"))
	//	s.App.HandleDir("/", s.AssetFile)
	//}

	db.InitDb()
	casbinUtils.InitCasbin()
	redisUtils.InitRedisCluster(common.GetRedisUris(), common.Config.Redis.Pwd)
	irisServer.baseRepo.Migrate()

	//iris.RegisterOnInterrupt(func() {
	//	_ = libs.Db
	//})

	routes.App(irisServer.App, accountCtrl, userCtrl, roleCtrl, permCtrl, tokenRepo)

	// deal with the command
	if *printVersion {
		fmt.Println(fmt.Sprintf("版本号：%s", version))
	}

	if *seederData {
		fmt.Println("填充数据：")
		fmt.Println()
		seederService.Run()
	}

	if *syncPerms {
		fmt.Println("同步权限：")
		fmt.Println()
		seederService.AddPerm()
	}

	if *printRouter {
		fmt.Println("系统权限：")
		fmt.Println()
		routes := irisServer.GetRoutes()
		for _, route := range routes {
			fmt.Println("+++++++++++++++")
			fmt.Println(fmt.Sprintf("名称 ：%s ", route.DisplayName))
			fmt.Println(fmt.Sprintf("路由地址 ：%s ", route.Name))
			fmt.Println(fmt.Sprintf("请求方式 ：%s", route.Act))
			fmt.Println()
		}
	}

	if common.IsPortInUse(common.Config.Port) {
		panic(fmt.Sprintf("端口 %d 已被使用", common.Config.Port))
	}

	// start the service
	err := irisServer.Serve()
	if err != nil {
		panic(err)
	}
}

func (s *Server) Serve() error {
	if common.Config.Https {
		host := fmt.Sprintf("%s:%d", common.Config.Host, 443)
		if err := s.App.Run(iris.TLS(host, common.Config.CertPath, common.Config.CertKey)); err != nil {
			return err
		}
	} else {
		if err := s.App.Run(
			iris.Addr(fmt.Sprintf("%s:%d", common.Config.Host, common.Config.Port)),
			iris.WithoutServerError(iris.ErrServerClosed),
			iris.WithOptimizations,
			iris.WithTimeFormat(time.RFC3339),
		); err != nil {
			return err
		}
	}

	return nil
}

type PathName struct {
	Name   string
	Path   string
	Method string
}

// 获取路由信息
func (s *Server) GetRoutes() []*models.Permission {
	var rrs []*models.Permission
	names := getPathNames(s.App.GetRoutesReadOnly())
	if common.Config.Debug {
		fmt.Println(fmt.Sprintf("路由权限集合：%v", names))
		fmt.Println(fmt.Sprintf("Iris App ：%v", s.App))
	}
	for _, pathName := range names {
		if !isPermRoute(pathName.Name) {
			rr := &models.Permission{Name: pathName.Path, DisplayName: pathName.Name, Description: pathName.Name, Act: pathName.Method}
			rrs = append(rrs, rr)
		}
	}
	return rrs
}

func getPathNames(routeReadOnly []context.RouteReadOnly) []*PathName {
	var pns []*PathName
	if common.Config.Debug {
		fmt.Println(fmt.Sprintf("routeReadOnly：%v", routeReadOnly))
	}
	for _, s := range routeReadOnly {
		pn := &PathName{
			Name:   s.Name(),
			Path:   s.Path(),
			Method: s.Method(),
		}
		pns = append(pns, pn)
	}

	return pns
}

// 过滤非必要权限
func isPermRoute(name string) bool {
	exceptRouteName := []string{"OPTIONS", "GET", "POST", "HEAD", "PUT", "PATCH", "payload"}
	for _, er := range exceptRouteName {
		if strings.Contains(name, er) {
			return true
		}
	}
	return false
}
