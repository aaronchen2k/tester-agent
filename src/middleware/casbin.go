package middleware

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/openstc/src/libs/common"
	redisUtils "github.com/aaronchen2k/openstc/src/libs/redis"
	"github.com/aaronchen2k/openstc/src/repo"
	"github.com/casbin/casbin/v2"
	"github.com/fatih/color"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"net/http"
)

type CasbinService struct {
	Enforcer  *casbin.Enforcer `inject:""`
	TokenRepo *repo.TokenRepo  `inject:""`
}

func NewCasbinService() *CasbinService {
	return &CasbinService{}
}

func (c *CasbinService) ServeHTTP(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
	value := ctx.Values().Get("jwt").(*jwt.Token)

	conn := redisUtils.GetRedisClusterClient()
	defer conn.Close()

	sess, err := c.TokenRepo.GetRedisSessionV2(conn, value.Raw)
	if err != nil {
		c.TokenRepo.UserTokenExpired(value.Raw)
		_, _ = ctx.JSON(common.ApiResource(401, nil, ""))
		ctx.StopExecution()
		return
	}
	if sess == nil {
		ctx.StopExecution()
		_, _ = ctx.JSON(common.ApiResource(401, nil, ""))
		ctx.StopExecution()
		return
	} else {
		check, err := c.Check(ctx.Request(), sess.UserId)
		if !check {
			_, _ = ctx.JSON(common.ApiResource(403, nil, err.Error()))
			ctx.StopExecution()
			return
		} else {
			ctx.Values().Set("sess", sess)
		}
	}

	ctx.Next()
}

// Check checks the username, request's method and path and
// returns true if permission grandted otherwise false.
func (c *CasbinService) Check(r *http.Request, userId string) (bool, error) {
	method := r.Method
	path := r.URL.Path
	ok, err := c.Enforcer.Enforce(userId, path, method)
	if err != nil {
		color.Red("验证权限报错：%v;%s-%s-%s", err.Error(), userId, path, method)
		return false, err
	}
	if !ok {
		msg := fmt.Sprintf("你未拥有 %s:%s 操作权限", method, path)
		color.Red(msg)
		return ok, errors.New(msg)
	}
	return ok, nil
}
