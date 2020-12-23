package service

import (
	"fmt"
	"github.com/aaronchen2k/openstc/src/libs/common"
	redisUtils "github.com/aaronchen2k/openstc/src/libs/redis"
	"github.com/aaronchen2k/openstc/src/middleware"
	"github.com/aaronchen2k/openstc/src/models"
	"github.com/aaronchen2k/openstc/src/repo"
	"github.com/fatih/color"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/jameskeane/bcrypt"
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
func (s *UserService) CheckLogin(u *models.User, password string) (*models.Token, int64, string) {

	if u.ID == 0 {
		return nil, 400, "用户不存在"
	} else {
		uid := strconv.FormatUint(uint64(u.ID), 10)
		if s.TokenRepo.IsUserTokenOver(uid) {
			return nil, 400, "以达到同时登录设备上限"
		}
		if ok := bcrypt.Match(password, u.Password); ok {
			token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
				"iat": time.Now().Unix(),
			})
			tokenString, _ := token.SignedString([]byte("HS2JDFKhu7Y1av7b"))

			rsv2 := models.RedisSessionV2{
				UserId:       uid,
				LoginType:    models.LoginTypeWeb,
				AuthType:     models.AuthPwd,
				CreationDate: time.Now().Unix(),
				Scope:        s.TokenRepo.GetUserScope("admin"),
			}
			conn := redisUtils.GetRedisClusterClient()
			defer conn.Close()

			if err := s.TokenRepo.ToCache(conn, rsv2, tokenString); err != nil {
				return nil, 400, err.Error()
			}

			if err := s.TokenRepo.SyncUserTokenCache(conn, rsv2, tokenString); err != nil {
				return nil, 400, err.Error()
			}

			return &models.Token{Token: tokenString}, 200, "登陆成功"
		} else {
			return nil, 400, "用户名或密码错误"
		}
	}
}

// CreateUser create user
func (s *UserService) CreateUser(u *models.User) error {
	u.Password = common.HashPassword(u.Password)
	if err := s.UserRepo.DB.Create(u).Error; err != nil {
		return err
	}

	s.addRoles(u)

	return nil
}

// UpdateUserById update user by id
func (s *UserService) UpdateUserById(id uint, nu *models.User) error {
	if len(nu.Password) > 0 {
		nu.Password = common.HashPassword(nu.Password)
	}
	if err := s.UserRepo.Update(&models.User{}, nu, id); err != nil {
		return err
	}

	s.addRoles(nu)
	return nil
}

// addRoles add roles for user
func (s *UserService) addRoles(user *models.User) {
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
