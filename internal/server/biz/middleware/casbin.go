package middleware

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/openstc/internal/server/biz/domain"
	"github.com/aaronchen2k/openstc/internal/server/biz/redis"
	"github.com/aaronchen2k/openstc/internal/server/biz/session"
	"github.com/aaronchen2k/openstc/internal/server/cfg"
	"github.com/aaronchen2k/openstc/internal/server/repo"
	"github.com/aaronchen2k/openstc/internal/server/utils"
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

func (m *CasbinService) ServeHTTP(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
	value := ctx.Values().Get("jwt").(*jwt.Token)

	var (
		credentials *domain.UserCredentials
		err         error
	)

	if serverConf.Config.Redis.Enable {
		conn := redisUtils.GetRedisClusterClient()
		defer conn.Close()

		credentials, err = m.TokenRepo.GetRedisSession(conn, value.Raw)
		if err != nil || credentials == nil {
			m.TokenRepo.UserTokenExpired(value.Raw)
			_, _ = ctx.JSON(agentUtils.ApiRes(401, "", nil))
			ctx.StopExecution()
			return
		}
	} else {
		credentials = sessionUtils.GetCredentials(ctx)
	}

	if credentials == nil {
		ctx.StopExecution()
		_, _ = ctx.JSON(agentUtils.ApiRes(401, "", nil))
		ctx.StopExecution()
		return
	} else {
		check, err := m.Check(ctx.Request(), credentials.UserId)
		if !check {
			_, _ = ctx.JSON(agentUtils.ApiRes(403, err.Error(), nil))
			ctx.StopExecution()
			return
		} else {
			ctx.Values().Set("sess", credentials)
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
