package controllers

import (
	"github.com/aaronchen2k/openstc/src/libs/common"
	"github.com/aaronchen2k/openstc/src/libs/redis"
	"github.com/aaronchen2k/openstc/src/models"
	"github.com/aaronchen2k/openstc/src/repo"
	"github.com/aaronchen2k/openstc/src/validates"
	"github.com/go-playground/validator/v10"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

type AccountController struct {
	userRepo  *repo.UserRepo
	tokenRepo *repo.TokenRepo
	roleRepo  *repo.RoleRepo
	permRepo  *repo.PermRepo
}

func NewAccountController(userRepo *repo.UserRepo,
	roleRepo *repo.RoleRepo, permRepo *repo.PermRepo, tokenRepo *repo.TokenRepo) *AccountController {

	return &AccountController{userRepo: userRepo, tokenRepo: tokenRepo, roleRepo: roleRepo, permRepo: permRepo}
}

/**
* @api {post} /admin/login 用户登陆
* @apiName 用户登陆
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 用户登陆
* @apiSampleRequest /admin/login
* @apiParam {string} username 用户名
* @apiParam {string} password 密码
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *AccountController) UserLogin(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	aul := new(validates.LoginRequest)

	if err := ctx.ReadJSON(aul); err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}

	err := validates.Validate.Struct(*aul)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validates.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(common.ApiResource(400, nil, e))
				return
			}
		}
	}

	ctx.Application().Logger().Infof("%s 登录系统", aul.Username)

	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "username",
				Condition: "=",
				Value:     aul.Username,
			},
		},
	}

	user, err := c.userRepo.GetUser(s)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}

	response, code, msg := c.userRepo.CheckLogin(user, aul.Password)

	_, _ = ctx.JSON(common.ApiResource(code, response, msg))
}

/**
* @api {get} /logout 用户退出登陆
* @apiName 用户退出登陆
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 用户退出登陆
* @apiSampleRequest /logout
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *AccountController) UserLogout(ctx iris.Context) {

	ctx.StatusCode(iris.StatusOK)
	value := ctx.Values().Get("jwt").(*jwt.Token)
	conn := redisUtils.GetRedisClusterClient()
	defer conn.Close()
	sess, err := c.tokenRepo.GetRedisSessionV2(conn, value.Raw)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}
	if sess != nil {
		if err := c.tokenRepo.DelUserTokenCache(conn, *sess, value.Raw); err != nil {
			_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
			return
		}
	}

	ctx.Application().Logger().Infof("%d 退出系统", sess.UserId)
	_, _ = ctx.JSON(common.ApiResource(200, nil, "退出"))
}

/**
* @api {get} /expire 刷新token
* @apiName 刷新token
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 刷新token
* @apiSampleRequest /expire
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *AccountController) UserExpire(ctx iris.Context) {

	ctx.StatusCode(iris.StatusOK)
	value := ctx.Values().Get("jwt").(*jwt.Token)
	conn := redisUtils.GetRedisClusterClient()
	defer conn.Close()
	sess, err := c.tokenRepo.GetRedisSessionV2(conn, value.Raw)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}
	if sess != nil {
		if err := c.tokenRepo.UpdateUserTokenCacheExpire(conn, *sess, value.Raw); err != nil {
			_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
			return
		}
	}

	_, _ = ctx.JSON(common.ApiResource(200, nil, ""))
}
