package controller

import (
	"github.com/aaronchen2k/openstc-common/src/libs/convertor"
	"github.com/aaronchen2k/openstc/src/domain"
	"github.com/aaronchen2k/openstc/src/libs/common"
	sessionUtils "github.com/aaronchen2k/openstc/src/libs/session"
	"github.com/aaronchen2k/openstc/src/model"
	"github.com/aaronchen2k/openstc/src/repo"
	"github.com/aaronchen2k/openstc/src/service"
	"github.com/aaronchen2k/openstc/src/transformer"
	"github.com/aaronchen2k/openstc/src/validate"
	"github.com/go-playground/validator/v10"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
)

type UserController struct {
	UserService *service.UserService `inject:""`
	RoleService *service.RoleService `inject:""`
	UserRepo    *repo.UserRepo       `inject:""`
	RoleRepo    *repo.RoleRepo       `inject:""`
}

func NewUserController() *UserController {
	return &UserController{}
}

/**
* @api {get} /admin/profile 获取登陆用户信息
* @apiName 获取登陆用户信息
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 获取登陆用户信息
* @apiSampleRequest /admin/profile
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func (c *UserController) GetProfile(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	cred := sessionUtils.GetCredentials(ctx)
	if cred == nil {
		_, _ = ctx.JSON(common.ApiResource(401, nil, "not login"))
		return
	}

	id := uint(common.ParseInt(cred.UserId, 10))
	s := &domain.Search{
		Fields: []*domain.Filed{
			{
				Key:       "id",
				Condition: "=",
				Value:     id,
			},
		},
	}
	user, err := c.UserRepo.GetUser(s)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}
	_, _ = ctx.JSON(common.ApiResource(200, c.userTransform(user), "请求成功"))
}

/**
* @api {get} /profile 管理员信息
* @apiName 管理员信息
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 管理员信息
* @apiSampleRequest /profile
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func (c *UserController) GetAdminInfo(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	user, err := c.UserRepo.GetUser(nil)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}
	_, _ = ctx.JSON(common.ApiResource(200, map[string]string{"avatar": user.Avatar}, "请求成功"))
}

/**
* @api {get} /admin/change_avatar 修改头像
* @apiName 修改头像
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 修改头像
* @apiSampleRequest /admin/change_avatar
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func (c *UserController) ChangeAvatar(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	sess := sessionUtils.GetCredentials(ctx)
	id := uint(common.ParseInt(sess.UserId, 10))

	avatar := new(model.Avatar)
	if err := ctx.ReadJSON(avatar); err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}

	err := validate.Validate.Struct(*avatar)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(common.ApiResource(400, nil, e))
				return
			}
		}
	}

	user := c.UserRepo.NewUser()
	user.ID = id
	user.Avatar = avatar.Avatar
	err = c.UserService.UpdateUserById(id, user)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}
	_, _ = ctx.JSON(common.ApiResource(200, c.userTransform(user), "请求成功"))
}

/**
* @api {get} /admin/users/:id 根据id获取用户信息
* @apiName 根据id获取用户信息
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 根据id获取用户信息
* @apiSampleRequest /admin/users/:id
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func (c *UserController) GetUser(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	//id, _ := ctx.Params().GetUint("id")

	user, err := c.UserRepo.GetUser(nil)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}
	_, _ = ctx.JSON(common.ApiResource(200, c.userTransform(user), "操作成功"))
}

/**
* @api {post} /admin/users/ 新建账号
* @apiName 新建账号
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 新建账号
* @apiSampleRequest /admin/users/
* @apiParam {string} username 用户名
* @apiParam {string} password 密码
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *UserController) CreateUser(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	user := new(model.User)
	if err := ctx.ReadJSON(user); err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}

	err := validate.Validate.Struct(*user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(common.ApiResource(400, nil, e))
				return
			}
		}
	}

	err = c.UserService.CreateUser(user)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}

	if user.ID == 0 {
		_, _ = ctx.JSON(common.ApiResource(400, nil, "操作失败"))
		return
	}
	_, _ = ctx.JSON(common.ApiResource(200, c.userTransform(user), "操作成功"))
	return

}

/**
* @api {post} /admin/users/:id/update 更新账号
* @apiName 更新账号
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 更新账号
* @apiSampleRequest /admin/users/:id/update
* @apiParam {string} username 用户名
* @apiParam {string} password 密码
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *UserController) UpdateUser(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	user := new(model.User)

	if err := ctx.ReadJSON(user); err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
	}

	err := validate.Validate.Struct(*user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(common.ApiResource(400, nil, e))
				return
			}
		}
	}

	id, _ := ctx.Params().GetUint("id")
	if user.Username == "username" {
		_, _ = ctx.JSON(common.ApiResource(400, nil, "不能编辑管理员"))
		return
	}

	err = c.UserService.UpdateUserById(id, user)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}
	_, _ = ctx.JSON(common.ApiResource(200, c.userTransform(user), "操作成功"))
}

/**
* @api {delete} /admin/users/:id/delete 删除用户
* @apiName 删除用户
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 删除用户
* @apiSampleRequest /admin/users/:id/delete
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *UserController) DeleteUser(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	id, _ := ctx.Params().GetUint("id")

	err := c.UserRepo.DeleteUser(id)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}
	_, _ = ctx.JSON(common.ApiResource(200, nil, "删除成功"))
}

/**
* @api {get} /users 获取所有的账号
* @apiName 获取所有的账号
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 获取所有的账号
* @apiSampleRequest /users
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *UserController) GetAllUsers(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	//name := ctx.FormValue("name")

	users, count, err := c.UserRepo.GetAllUsers(nil)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}

	transform := c.usersTransform(users)
	_, _ = ctx.JSON(common.ApiResource(200,
		map[string]interface{}{"items": transform, "total": count, "limit": "s.Limit"}, "操作成功"))

}

func (c *UserController) usersTransform(users []*model.User) []*transformer.User {
	var us []*transformer.User
	for _, user := range users {
		u := c.userTransform(user)
		us = append(us, u)
	}
	return us
}

func (c *UserController) userTransform(user *model.User) *transformer.User {
	u := &transformer.User{}
	g := convertor.NewTransform(u, user, time.RFC3339)
	_ = g.Transformer()

	roleIds := c.RoleService.GetRolesForUser(user.ID)
	var ris []int
	for _, roleId := range roleIds {
		ri, _ := strconv.Atoi(roleId)
		ris = append(ris, ri)
	}
	roles, _, err := c.RoleRepo.GetAllRoles(nil)
	if err == nil {
		u.Roles = c.RoleService.RolesTransform(roles)
	}
	return u
}
