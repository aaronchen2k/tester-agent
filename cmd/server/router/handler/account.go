package handler

import (
	"github.com/aaronchen2k/tester/internal/pkg/utils"
	"github.com/aaronchen2k/tester/internal/server/biz/domain"
	"github.com/aaronchen2k/tester/internal/server/biz/redis"
	"github.com/aaronchen2k/tester/internal/server/biz/session"
	"github.com/aaronchen2k/tester/internal/server/biz/validate"
	"github.com/aaronchen2k/tester/internal/server/cfg"
	"github.com/aaronchen2k/tester/internal/server/repo"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/go-playground/validator/v10"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

type AccountController struct {
	UserService *service.UserService `inject:""`

	UserRepo  *repo.UserRepo  `inject:""`
	TokenRepo *repo.TokenRepo `inject:""`
	RoleRepo  *repo.RoleRepo  `inject:""`
	PermRepo  *repo.PermRepo  `inject:""`
}

func NewAccountController() *AccountController {
	return &AccountController{}
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
	aul := new(validate.LoginRequest)

	if err := ctx.ReadJSON(aul); err != nil {
		_, _ = ctx.JSON(utils.ApiRes(400, err.Error(), nil))
		return
	}

	err := validate.Validate.Struct(*aul)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(utils.ApiRes(400, e, nil))
				return
			}
		}
	}

	ctx.Application().Logger().Infof("%s 登录系统", aul.Username)

	search := &domain.Search{
		Fields: []*domain.Filed{
			{
				Key:       "username",
				Condition: "=",
				Value:     aul.Username,
			},
		},
	}
	user, err := c.UserRepo.GetUser(search)
	if err != nil {
		_, _ = ctx.JSON(utils.ApiRes(400, err.Error(), nil))
		return
	}

	response, code, msg := c.UserService.CheckLogin(ctx, user, aul.Password)
	response.RememberMe = aul.RememberMe

	refreshToken := ""
	if code == 200 && aul.RememberMe {
		refreshToken = response.Token
	}

	c.UserService.UpdateRefreshToken(user.ID, refreshToken)

	_, _ = ctx.JSON(utils.ApiRes(code, msg, response))
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

	var (
		credentials *domain.UserCredentials
		err         error
	)
	if serverConf.Config.Redis.Enable {
		conn := redisUtils.GetRedisClusterClient()
		defer conn.Close()

		credentials, err = c.TokenRepo.GetRedisSession(conn, value.Raw)
		if err != nil {
			_, _ = ctx.JSON(utils.ApiRes(400, err.Error(), nil))
			return
		}
		if credentials != nil {
			if err := c.TokenRepo.DelUserTokenCache(conn, *credentials, value.Raw); err != nil {
				_, _ = ctx.JSON(utils.ApiRes(400, err.Error(), nil))
				return
			}
		}
	} else {
		credentials = sessionUtils.GetCredentials(ctx)
		if credentials == nil {
			_, _ = ctx.JSON(utils.ApiRes(400, err.Error(), nil))
			return
		} else {
			sessionUtils.RemoveCredentials(ctx)
		}
	}

	ctx.Application().Logger().Infof("%d 退出系统", credentials.UserId)
	_, _ = ctx.JSON(utils.ApiRes(200, "退出", nil))
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
	sess, err := c.TokenRepo.GetRedisSession(conn, value.Raw)
	if err != nil {
		_, _ = ctx.JSON(utils.ApiRes(400, err.Error(), nil))
		return
	}
	if sess != nil {
		if err := c.TokenRepo.UpdateUserTokenCacheExpire(conn, *sess, value.Raw); err != nil {
			_, _ = ctx.JSON(utils.ApiRes(400, err.Error(), nil))
			return
		}
	}

	_, _ = ctx.JSON(utils.ApiRes(200, "", nil))
}
