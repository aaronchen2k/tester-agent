package handler

import (
	"github.com/aaronchen2k/openstc/internal/pkg/libs/convertor"
	"github.com/aaronchen2k/openstc/internal/server/biz/domain"
	sessionUtils "github.com/aaronchen2k/openstc/internal/server/biz/session"
	"github.com/aaronchen2k/openstc/internal/server/biz/transformer"
	"github.com/aaronchen2k/openstc/internal/server/biz/validate"
	"github.com/aaronchen2k/openstc/internal/server/model"
	"github.com/aaronchen2k/openstc/internal/server/repo"
	"github.com/aaronchen2k/openstc/internal/server/service"
	"github.com/aaronchen2k/openstc/internal/server/utils"
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
		_, _ = ctx.JSON(agentUtils.ApiRes(401, "not login", nil))
		return
	}

	id := uint(agentUtils.ParseInt(cred.UserId, 10))
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
		_, _ = ctx.JSON(agentUtils.ApiRes(400, err.Error(), nil))
		return
	}
	_, _ = ctx.JSON(agentUtils.ApiRes(200, "请求成功", c.userTransform(user)))
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
		_, _ = ctx.JSON(agentUtils.ApiRes(400, err.Error(), nil))
		return
	}
	_, _ = ctx.JSON(agentUtils.ApiRes(200, "请求成功", map[string]string{"avatar": user.Avatar}))
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
	id := uint(agentUtils.ParseInt(sess.UserId, 10))

	avatar := new(model.Avatar)
	if err := ctx.ReadJSON(avatar); err != nil {
		_, _ = ctx.JSON(agentUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := validate.Validate.Struct(*avatar)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(agentUtils.ApiRes(400, e, nil))
				return
			}
		}
	}

	user := c.UserRepo.NewUser()
	user.ID = id
	user.Avatar = avatar.Avatar
	err = c.UserService.UpdateUserById(id, user)
	if err != nil {
		_, _ = ctx.JSON(agentUtils.ApiRes(400, err.Error(), nil))
		return
	}
	_, _ = ctx.JSON(agentUtils.ApiRes(200, "请求成功", c.userTransform(user)))
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
		_, _ = ctx.JSON(agentUtils.ApiRes(400, err.Error(), nil))
		return
	}
	_, _ = ctx.JSON(agentUtils.ApiRes(200, "操作成功", c.userTransform(user)))
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
		_, _ = ctx.JSON(agentUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := validate.Validate.Struct(*user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(agentUtils.ApiRes(400, e, nil))
				return
			}
		}
	}

	err = c.UserService.CreateUser(user)
	if err != nil {
		_, _ = ctx.JSON(agentUtils.ApiRes(400, err.Error(), nil))
		return
	}

	if user.ID == 0 {
		_, _ = ctx.JSON(agentUtils.ApiRes(400, "操作失败", nil))
		return
	}
	_, _ = ctx.JSON(agentUtils.ApiRes(200, "操作成功", c.userTransform(user)))
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
		_, _ = ctx.JSON(agentUtils.ApiRes(400, err.Error(), nil))
	}

	err := validate.Validate.Struct(*user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(agentUtils.ApiRes(400, e, nil))
				return
			}
		}
	}

	id, _ := ctx.Params().GetUint("id")
	if user.Username == "username" {
		_, _ = ctx.JSON(agentUtils.ApiRes(400, "不能编辑管理员", nil))
		return
	}

	err = c.UserService.UpdateUserById(id, user)
	if err != nil {
		_, _ = ctx.JSON(agentUtils.ApiRes(400, err.Error(), nil))
		return
	}
	_, _ = ctx.JSON(agentUtils.ApiRes(200, "操作成功", c.userTransform(user)))
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
		_, _ = ctx.JSON(agentUtils.ApiRes(400, err.Error(), nil))
		return
	}
	_, _ = ctx.JSON(agentUtils.ApiRes(200, "删除成功", nil))
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
		_, _ = ctx.JSON(agentUtils.ApiRes(400, err.Error(), nil))
		return
	}

	transform := c.usersTransform(users)
	_, _ = ctx.JSON(agentUtils.ApiRes(200, "操作成功", map[string]interface{}{"items": transform, "total": count, "limit": "s.Limit"}))

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
