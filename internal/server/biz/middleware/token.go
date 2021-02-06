package middleware

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	"github.com/aaronchen2k/tester/internal/pkg/utils"
	"github.com/aaronchen2k/tester/internal/server/biz/domain"
	"github.com/aaronchen2k/tester/internal/server/biz/redis"
	"github.com/aaronchen2k/tester/internal/server/biz/session"
	"github.com/aaronchen2k/tester/internal/server/cfg"
	"github.com/aaronchen2k/tester/internal/server/repo"
	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"net/http"
	"strconv"
	"time"
)

type TokenService struct {
	Middleware *JwtService      `inject:""`
	Enforcer   *casbin.Enforcer `inject:""`
	UserRepo   *repo.UserRepo   `inject:""`
	TokenRepo  *repo.TokenRepo  `inject:""`
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

// Get returns the user (&token) information for this client/request
func (m *TokenService) Get(ctx iris.Context) *jwt.Token {
	v := ctx.Values().Get(m.Middleware.Config.ContextKey)
	if v == nil {
		return nil
	}
	return v.(*jwt.Token)
}

func (m *TokenService) Serve(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
	value := m.Get(ctx)

	credentials, _ := m.GetCredentials(value, ctx)
	if credentials == nil { // jwt token expired, try to refresh
		tokenStr := value.Raw

		// find user by token
		user, _ := m.UserRepo.GetByToken(tokenStr)

		if user.ID != 0 && // user exist and token not expire
			time.Now().Unix()-user.TokenUpdatedTime.Unix() < _const.UserTokenExpireTime {

			// refresh the credentials
			uid := strconv.FormatUint(uint64(user.ID), 10)
			cred := domain.UserCredentials{
				UserId:       uid,
				LoginType:    domain.LoginTypeWeb,
				AuthType:     domain.AuthPwd,
				CreationDate: time.Now().Unix(),
				Scope:        domain.AdminScope,
				Token:        tokenStr,
			}

			if serverConf.Config.Redis.Enable {
				conn := redisUtils.GetRedisClusterClient()
				defer conn.Close()

				if err := m.TokenRepo.CacheToRedis(conn, cred, tokenStr); err != nil {
					m.Middleware.Config.ErrorHandler(ctx, err)
					return
				}
				if err := m.TokenRepo.SyncUserTokenCache(conn, cred, tokenStr); err != nil {
					m.Middleware.Config.ErrorHandler(ctx, err)
					return
				}
			} else {
				sessionUtils.SaveCredentials(ctx, &cred)
			}
		}
	}

	// load again
	credentials, _ = m.GetCredentials(value, ctx)
	if credentials == nil {
		ctx.StopExecution()
		_, _ = ctx.JSON(_utils.ApiRes(401, "", nil))
		ctx.StopExecution()
		return
	}

	ctx.Next()
}

func (m *TokenService) GetCredentials(value *jwt.Token, ctx iris.Context) (
	credentials *domain.UserCredentials, err error) {
	if serverConf.Config.Redis.Enable {
		conn := redisUtils.GetRedisClusterClient()
		defer conn.Close()

		credentials, err = m.TokenRepo.GetRedisSession(conn, value.Raw)
		if err != nil || credentials == nil {
			m.TokenRepo.UserTokenExpired(value.Raw)
			_, _ = ctx.JSON(_utils.ApiRes(401, "", nil))
			ctx.StopExecution()
			return
		}
	} else {
		credentials = sessionUtils.GetCredentials(ctx)
	}

	return
}
