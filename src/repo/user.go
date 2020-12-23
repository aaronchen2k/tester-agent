package repo

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/openstc/src/models"
	"gorm.io/gorm"
	"time"

	"github.com/fatih/color"
)

type UserRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) NewUser() *models.User {
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

	if err := r.DB.Delete(u, id).Error; err != nil {
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
