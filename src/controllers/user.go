package controllers

import (
	"github.com/aaronchen2k/openstc/src/libs/common"
	"github.com/aaronchen2k/openstc/src/models"
	"github.com/aaronchen2k/openstc/src/repo"
	"github.com/aaronchen2k/openstc/src/transformer"
	"github.com/aaronchen2k/openstc/src/validates"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	gf "github.com/snowlyg/gotransformer"
)

type UserController struct {
	BaseController
	userRepo *repo.UserRepo
	roleRepo *repo.RoleRepo
}

func NewUserController(userRepo *repo.UserRepo, roleRepo *repo.RoleRepo) *UserController {
	return &UserController{userRepo: userRepo, roleRepo: roleRepo}
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
	sess := ctx.Values().Get("sess").(*models.RedisSessionV2)
	id := uint(common.ParseInt(sess.UserId, 10))
	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "id",
				Condition: "=",
				Value:     id,
			},
		},
	}
	user, err := c.userRepo.GetUser(s)
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
	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "username",
				Condition: "=",
				Value:     "username",
			},
		},
	}
	user, err := c.userRepo.GetUser(s)
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
	sess := ctx.Values().Get("sess").(*models.RedisSessionV2)
	id := uint(common.ParseInt(sess.UserId, 10))

	avatar := new(models.Avatar)
	if err := ctx.ReadJSON(avatar); err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}

	err := validates.Validate.Struct(*avatar)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validates.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(common.ApiResource(400, nil, e))
				return
			}
		}
	}

	user := c.userRepo.NewUser()
	user.ID = id
	user.Avatar = avatar.Avatar
	err = c.userRepo.UpdateUserById(id, user)
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
	id, _ := ctx.Params().GetUint("id")
	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "id",
				Condition: "=",
				Value:     id,
			},
		},
	}
	user, err := c.userRepo.GetUser(s)
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
	user := new(models.User)
	if err := ctx.ReadJSON(user); err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}

	err := validates.Validate.Struct(*user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validates.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(common.ApiResource(400, nil, e))
				return
			}
		}
	}

	err = c.userRepo.CreateUser(user)
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
	user := new(models.User)

	if err := ctx.ReadJSON(user); err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
	}

	err := validates.Validate.Struct(*user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validates.ValidateTrans) {
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

	err = c.userRepo.UpdateUserById(id, user)
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

	err := c.userRepo.DeleteUser(id)
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
	s := c.GetCommonListSearch(ctx)
	name := ctx.FormValue("name")

	s.Fields = append(s.Fields, c.userRepo.GetSearch("name", name))
	users, count, err := c.userRepo.GetAllUsers(s)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}

	transform := c.usersTransform(users)
	_, _ = ctx.JSON(common.ApiResource(200, map[string]interface{}{"items": transform, "total": count, "limit": s.Limit}, "操作成功"))

}

func (c *UserController) usersTransform(users []*models.User) []*transformer.User {
	var us []*transformer.User
	for _, user := range users {
		u := c.userTransform(user)
		us = append(us, u)
	}
	return us
}

func (c *UserController) userTransform(user *models.User) *transformer.User {
	u := &transformer.User{}
	g := gf.NewTransform(u, user, time.RFC3339)
	_ = g.Transformer()

	roleIds := c.userRepo.GetRolesForUser(user.ID)
	var ris []int
	for _, roleId := range roleIds {
		ri, _ := strconv.Atoi(roleId)
		ris = append(ris, ri)
	}
	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "id",
				Condition: "in",
				Value:     ris,
			},
		},
	}
	roles, _, err := c.roleRepo.GetAllRoles(s)
	if err == nil {
		u.Roles = c.roleRepo.RolesTransform(roles)
	}
	return u
}
