package server

import (
	"fmt"
	"github.com/aaronchen2k/openstc/src/controllers"
	"github.com/aaronchen2k/openstc/src/libs/common"
	"github.com/aaronchen2k/openstc/src/libs/db"
	redisUtils "github.com/aaronchen2k/openstc/src/libs/redis"
	"github.com/aaronchen2k/openstc/src/middleware"
	"github.com/aaronchen2k/openstc/src/models"
	"github.com/aaronchen2k/openstc/src/repo"
	"github.com/aaronchen2k/openstc/src/routes"
	"github.com/aaronchen2k/openstc/src/service"
	"github.com/facebookgo/inject"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"

	"github.com/kataras/iris/v12/context"
)

type Server struct {
	App       *iris.Application
	AssetFile http.FileSystem
}

func NewServer(assetFile http.FileSystem) *Server {
	app := iris.Default()
	return &Server{
		App:       app,
		AssetFile: assetFile,
	}
}

func Init(version string, printVersion, printRouter *bool) {

	db.InitDB()
	db.GetInst().Migrate()

	// irisServer := server.NewServer(AssetFile()) // 加载前端文件
	irisServer := NewServer(nil)
	if irisServer == nil {
		panic("Http 初始化失败")
	}
	irisServer.App.Logger().SetLevel(common.Config.LogLevel)

	//if common.Config.BinData {
	//	irisServer.App.RegisterView(iris.HTML(irisServer.AssetFile, ".html"))
	//	irisServer.App.HandleDir("/", irisServer.AssetFile)
	//}

	router := routes.NewRouter(irisServer.App)
	injectObj(router)
	router.App()

	//casbinUtils.InitCasbin()
	redisUtils.InitRedisCluster(common.GetRedisUris(), common.Config.Redis.Pwd)

	iris.RegisterOnInterrupt(func() {
		defer db.GetInst().Close()
	})

	// deal with the command
	if *printVersion {
		fmt.Println(fmt.Sprintf("版本号：%s", version))
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

func injectObj(router *routes.Router) {
	// inject
	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	if err := g.Provide(
		// db
		&inject.Object{Value: db.GetInst().DB()},

		&inject.Object{Value: middleware.NewEnforcer()},
		&inject.Object{Value: middleware.NewCasbinService()},

		// repo
		&inject.Object{Value: repo.NewCommonRepo()},
		&inject.Object{Value: repo.NewPermRepo()},
		&inject.Object{Value: repo.NewTokenRepo()},
		&inject.Object{Value: repo.NewRoleRepo()},
		&inject.Object{Value: repo.NewUserRepo()},

		// service
		&inject.Object{Value: service.NewSeeder()},

		// controller
		&inject.Object{Value: controllers.NewInitController()},
		&inject.Object{Value: controllers.NewAccountController()},
		&inject.Object{Value: controllers.NewPermController()},
		&inject.Object{Value: controllers.NewRoleController()},
		&inject.Object{Value: controllers.NewUserController()},

		// router
		&inject.Object{Value: router},
	); err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}

	err := g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
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
