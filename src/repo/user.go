package repo

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/openstc/src/libs/casbin"
	"github.com/aaronchen2k/openstc/src/libs/common"
	"github.com/aaronchen2k/openstc/src/libs/db"
	redisUtils "github.com/aaronchen2k/openstc/src/libs/redis"
	"github.com/aaronchen2k/openstc/src/models"
	"gorm.io/gorm"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/jameskeane/bcrypt"
)

type UserRepo struct {
	BaseRepo
	tokenRepo *TokenRepo
}

func NewUserRepo(tokenRepo *TokenRepo) *UserRepo {
	return &UserRepo{tokenRepo: tokenRepo}
}

func (role *UserRepo) NewUser() *models.User {
	return &models.User{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

// GetUser get user
func (r *UserRepo) GetUser(search *models.Search) (*models.User, error) {
	t := r.NewUser()
	err := r.Found(search).First(t).Error
	if !r.IsNotFound(err) {
		return t, err
	}
	return t, nil
}

// DeleteUser del user . if user's username is username ,can't del it.
func (r *UserRepo) DeleteUser(id uint) error {
	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "id",
				Condition: "=",
				Value:     id,
			},
		},
	}
	u, err := r.GetUser(s)
	if err != nil {
		return err
	}
	if u.Username == "username" {
		return errors.New(fmt.Sprintf("不能删除管理员 : %s \n ", u.Username))
	}

	if err := db.Db.Delete(u, id).Error; err != nil {
		color.Red(fmt.Sprintf("DeleteUserByIdErr:%s \n ", err))
		return err
	}
	return nil
}

// GetAllUsers get all users
func (r *UserRepo) GetAllUsers(s *models.Search) ([]*models.User, int64, error) {
	var users []*models.User
	var count int64
	q := r.GetAll(&models.User{}, s)
	if err := q.Count(&count).Error; err != nil {
		return nil, count, err
	}
	q = q.Scopes(r.Paginate(s.Offset, s.Limit), r.Relation(s.Relations))
	if err := q.Find(&users).Error; err != nil {
		color.Red(fmt.Sprintf("GetAllUserErr:%s \n ", err))
		return nil, count, err
	}
	return users, count, nil
}

// CreateUser create user
func (r *UserRepo) CreateUser(u *models.User) error {
	u.Password = common.HashPassword(u.Password)
	if err := db.Db.Create(u).Error; err != nil {
		return err
	}

	r.addRoles(u)

	return nil
}

// UpdateUserById update user by id
func (r *UserRepo) UpdateUserById(id uint, nu *models.User) error {
	if len(nu.Password) > 0 {
		nu.Password = common.HashPassword(nu.Password)
	}
	if err := r.Update(&models.User{}, nu, id); err != nil {
		return err
	}

	r.addRoles(nu)
	return nil
}

// addRoles add roles for user
func (r *UserRepo) addRoles(user *models.User) {
	if len(user.RoleIds) > 0 {
		userId := strconv.FormatUint(uint64(user.ID), 10)
		if _, err := casbinUtils.Enforcer.DeleteRolesForUser(userId); err != nil {
			color.Red(fmt.Sprintf("CreateUserErr:%s \n ", err))
		}

		for _, roleId := range user.RoleIds {
			roleId := strconv.FormatUint(uint64(roleId), 10)
			if _, err := casbinUtils.Enforcer.AddRoleForUser(userId, roleId); err != nil {
				color.Red(fmt.Sprintf("CreateUserErr:%s \n ", err))
			}
		}
	}
}

// CheckLogin check login user
func (r *UserRepo) CheckLogin(u *models.User, password string) (*models.Token, int64, string) {

	if u.ID == 0 {
		return nil, 400, "用户不存在"
	} else {
		uid := strconv.FormatUint(uint64(u.ID), 10)
		if r.tokenRepo.isUserTokenOver(uid) {
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
				Scope:        r.tokenRepo.getUserScope("admin"),
			}
			conn := redisUtils.GetRedisClusterClient()
			defer conn.Close()

			if err := r.tokenRepo.ToCache(conn, rsv2, tokenString); err != nil {
				return nil, 400, err.Error()
			}

			if err := r.tokenRepo.SyncUserTokenCache(conn, rsv2, tokenString); err != nil {
				return nil, 400, err.Error()
			}

			return &models.Token{tokenString}, 200, "登陆成功"
		} else {
			return nil, 400, "用户名或密码错误"
		}
	}
}
