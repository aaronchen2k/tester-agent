package routes

import (
	"github.com/aaronchen2k/openstc/src/controllers"
	"github.com/aaronchen2k/openstc/src/libs/casbin"
	"github.com/aaronchen2k/openstc/src/libs/common"
	"github.com/aaronchen2k/openstc/src/middleware"
	"github.com/aaronchen2k/openstc/src/repo"
	"github.com/kataras/iris/v12"
)

func App(api *iris.Application,
	accountCtrl *controllers.AccountController, userCtrl *controllers.UserController,
	roleCtrl *controllers.RoleController, permCtrl *controllers.PermController,
	repo *repo.TokenRepo) {

	api.UseRouter(middleware.CrsAuth())
	app := api.Party("/api").AllowMethods(iris.MethodOptions)
	{
		// 二进制模式 ， 启用项目入口
		if common.Config.BinData {
			app.Get("/", func(ctx iris.Context) { // 首页模块
				_ = ctx.View("index.html")
			})
		}

		v1 := app.Party("/v1")
		{
			v1.Post("/admin/login", accountCtrl.UserLogin)

			v1.PartyFunc("/admin", func(admin iris.Party) {
				casbinMiddleware := middleware.New(casbinUtils.Enforcer, repo)       //casbin for gorm                                                   // <- IMPORTANT, register the middleware.
				admin.Use(middleware.JwtHandler().Serve, casbinMiddleware.ServeHTTP) //登录验证
				admin.Post("/logout", accountCtrl.UserLogout).Name = "退出"
				admin.Get("/expire", accountCtrl.UserExpire).Name = "刷新 token"
				admin.Get("/profile", userCtrl.GetProfile).Name = "个人信息"

				admin.PartyFunc("/users", func(users iris.Party) {
					users.Get("/", userCtrl.GetAllUsers).Name = "用户列表"
					users.Get("/{id:uint}", userCtrl.GetUser).Name = "用户详情"
					users.Post("/", userCtrl.CreateUser).Name = "创建用户"
					users.Put("/{id:uint}", userCtrl.UpdateUser).Name = "编辑用户"
					users.Delete("/{id:uint}", userCtrl.DeleteUser).Name = "删除用户"
				})
				admin.PartyFunc("/roles", func(roles iris.Party) {
					roles.Get("/", roleCtrl.GetAllRoles).Name = "角色列表"
					roles.Get("/{id:uint}", roleCtrl.GetRole).Name = "角色详情"
					roles.Post("/", roleCtrl.CreateRole).Name = "创建角色"
					roles.Put("/{id:uint}", roleCtrl.UpdateRole).Name = "编辑角色"
					roles.Delete("/{id:uint}", roleCtrl.DeleteRole).Name = "删除角色"
				})
				admin.PartyFunc("/permissions", func(permissions iris.Party) {
					permissions.Get("/", permCtrl.GetAllPermissions).Name = "权限列表"
					permissions.Get("/{id:uint}", permCtrl.GetPermission).Name = "权限详情"
					permissions.Post("/", permCtrl.CreatePermission).Name = "创建权限"
					permissions.Put("/{id:uint}", permCtrl.UpdatePermission).Name = "编辑权限"
					permissions.Delete("/{id:uint}", permCtrl.DeletePermission).Name = "删除权限"
				})
			})
		}
	}

}
