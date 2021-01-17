package service

import (
	"fmt"
	"github.com/aaronchen2k/tester/internal/server/biz/domain"
	"github.com/aaronchen2k/tester/internal/server/biz/middleware"
	"github.com/aaronchen2k/tester/internal/server/biz/redis"
	sessionUtils "github.com/aaronchen2k/tester/internal/server/biz/session"
	"github.com/aaronchen2k/tester/internal/server/cfg"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/repo"
	"github.com/aaronchen2k/tester/internal/server/utils"
	"github.com/fatih/color"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/jameskeane/bcrypt"
	"github.com/kataras/iris/v12"
	"strconv"
	"time"
)

type UserService struct {
	UserRepo  *repo.UserRepo  `inject:""`
	TokenRepo *repo.TokenRepo `inject:""`

	CasbinService *middleware.CasbinService `inject:""`
}

func NewUserService() *UserService {
	return &UserService{}
}

// CheckLogin check login user
func (s *UserService) CheckLogin(ctx iris.Context, u *model.User, password string) (*model.Token, int64, string) {

	if u.ID == 0 {
		return nil, 400, "用户不存在"
	} else {
		uid := strconv.FormatUint(uint64(u.ID), 10)
		if serverConf.Config.Redis.Enable && s.TokenRepo.IsUserTokenOver(uid) {
			return nil, 400, "已达到同时登录设备上限"
		}
		if ok := bcrypt.Match(password, u.Password); ok {
			token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
				"iat": time.Now().Unix(),
			})
			tokenStr, _ := token.SignedString([]byte("HS2JDFKhu7Y1av7b"))

			cred := domain.UserCredentials{
				UserId:       uid,
				LoginType:    domain.LoginTypeWeb,
				AuthType:     domain.AuthPwd,
				CreationDate: time.Now().Unix(),
				Scope:        s.TokenRepo.GetUserScope("admin"),
				Token:        tokenStr,
			}

			if serverConf.Config.Redis.Enable {
				conn := redisUtils.GetRedisClusterClient()
				defer conn.Close()

				if err := s.TokenRepo.CacheToRedis(conn, cred, tokenStr); err != nil {
					return nil, 400, err.Error()
				}
				if err := s.TokenRepo.SyncUserTokenCache(conn, cred, tokenStr); err != nil {
					return nil, 400, err.Error()
				}
			} else {
				sessionUtils.SaveCredentials(ctx, &cred)
			}

			return &model.Token{Token: tokenStr}, 200, "登陆成功"
		} else {
			return nil, 400, "用户名或密码错误"
		}
	}
}

// CreateUser create user
func (s *UserService) CreateUser(u *model.User) error {
	u.Password = agentUtils.HashPassword(u.Password)
	if err := s.UserRepo.DB.Create(u).Error; err != nil {
		return err
	}

	s.addRoles(u)

	return nil
}

// UpdateUserById update user by id
func (s *UserService) UpdateUserById(id uint, nu *model.User) error {
	if len(nu.Password) > 0 {
		nu.Password = agentUtils.HashPassword(nu.Password)
	}
	if err := s.UserRepo.Update(&model.User{}, nu, id); err != nil {
		return err
	}

	s.addRoles(nu)
	return nil
}

// addRoles add roles for user
func (s *UserService) addRoles(user *model.User) {
	if len(user.RoleIds) > 0 {
		userId := strconv.FormatUint(uint64(user.ID), 10)
		if _, err := s.CasbinService.Enforcer.DeleteRolesForUser(userId); err != nil {
			color.Red(fmt.Sprintf("CreateUserErr:%s \n ", err))
		}

		for _, roleId := range user.RoleIds {
			roleId := strconv.FormatUint(uint64(roleId), 10)
			if _, err := s.CasbinService.Enforcer.AddRoleForUser(userId, roleId); err != nil {
				color.Red(fmt.Sprintf("CreateUserErr:%s \n ", err))
			}
		}
	}
}

func (s *UserService) UpdateRefreshToken(id uint, token string) {
	s.UserRepo.UpdateRefreshToken(id, token)
}
