package router

import (
	"github.com/aaronchen2k/openstc/cmd/server/router/handler"
	"github.com/aaronchen2k/openstc/internal/server/biz/middleware"
	"github.com/aaronchen2k/openstc/internal/server/cfg"
	"github.com/aaronchen2k/openstc/internal/server/repo"
	"github.com/kataras/iris/v12"
)

type Router struct {
	api *iris.Application

	CasbinService *middleware.CasbinService `inject:""`

	AccountCtrl *handler.AccountController `inject:""`
	AppiumCtrl  *handler.AppiumController  `inject:""`
	DeviceCtrl  *handler.DeviceController  `inject:""`
	FileCtrl    *handler.FileController    `inject:""`
	HostCtrl    *handler.HostController    `inject:""`
	ImageCtrl   *handler.ImageController   `inject:""`
	InitCtrl    *handler.InitController    `inject:""`
	PermCtrl    *handler.PermController    `inject:""`
	RoleCtrl    *handler.RoleController    `inject:""`
	TaskCtrl    *handler.TaskController    `inject:""`
	UserCtrl    *handler.UserController    `inject:""`
	VmCtrl      *handler.VmController      `inject:""`

	TokenRepo *repo.TokenRepo `inject:""`
}

func NewRouter(app *iris.Application) *Router {
	router := &Router{}
	router.api = app

	return router
}

func (r *Router) App() {
	r.api.UseRouter(middleware.CrsAuth())

	app := r.api.Party("/api").AllowMethods(iris.MethodOptions)
	{
		// 二进制模式 ， 启用项目入口
		if serverConf.Config.BinData {
			app.Get("/", func(ctx iris.Context) { // 首页模块
				_ = ctx.View("index.html")
			})
		}

		v1 := app.Party("/v1")
		{
			v1.Get("/admin/init", r.InitCtrl.InitData)
			v1.Post("/admin/login", r.AccountCtrl.UserLogin)

			v1.PartyFunc("/admin", func(admin iris.Party) { // <- IMPORTANT, register the middleware.
				admin.Use(middleware.JwtHandler().Serve, r.CasbinService.ServeHTTP) //登录验证

				admin.Post("/logout", r.AccountCtrl.UserLogout).Name = "退出"
				admin.Get("/expire", r.AccountCtrl.UserExpire).Name = "刷新 token"
				admin.Get("/profile", r.UserCtrl.GetProfile).Name = "个人信息"

				admin.PartyFunc("/hosts", func(hosts iris.Party) {
					hosts.Get("/", r.HostCtrl.List).Name = "PVE节点列表"
					hosts.Get("/{id:uint}", r.HostCtrl.Get).Name = "用户详情"
				})

				admin.PartyFunc("/users", func(users iris.Party) {
					users.Get("/", r.UserCtrl.GetAllUsers).Name = "用户列表"
					users.Get("/{id:uint}", r.UserCtrl.GetUser).Name = "用户详情"
					users.Post("/", r.UserCtrl.CreateUser).Name = "创建用户"
					users.Put("/{id:uint}", r.UserCtrl.UpdateUser).Name = "编辑用户"
					users.Delete("/{id:uint}", r.UserCtrl.DeleteUser).Name = "删除用户"
				})
				admin.PartyFunc("/roles", func(roles iris.Party) {
					roles.Get("/", r.RoleCtrl.GetAllRoles).Name = "角色列表"
					roles.Get("/{id:uint}", r.RoleCtrl.GetRole).Name = "角色详情"
					roles.Post("/", r.RoleCtrl.CreateRole).Name = "创建角色"
					roles.Put("/{id:uint}", r.RoleCtrl.UpdateRole).Name = "编辑角色"
					roles.Delete("/{id:uint}", r.RoleCtrl.DeleteRole).Name = "删除角色"
				})
				admin.PartyFunc("/permissions", func(permissions iris.Party) {
					permissions.Get("/", r.PermCtrl.GetAllPermissions).Name = "权限列表"
					permissions.Get("/{id:uint}", r.PermCtrl.GetPermission).Name = "权限详情"
					permissions.Post("/", r.PermCtrl.CreatePermission).Name = "创建权限"
					permissions.Put("/{id:uint}", r.PermCtrl.UpdatePermission).Name = "编辑权限"
					permissions.Delete("/{id:uint}", r.PermCtrl.DeletePermission).Name = "删除权限"
				})
			})
		}
	}

}
