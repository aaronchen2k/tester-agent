package repo

import (
	"fmt"
	"github.com/aaronchen2k/openstc/src/models"
	"time"

	"github.com/fatih/color"
	"gorm.io/gorm"
)

type RoleRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewRoleRepo() *RoleRepo {
	return &RoleRepo{}
}

func (r *RoleRepo) NewRole() *models.Role {
	return &models.Role{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

// GetRole get role
func (r *RoleRepo) GetRole(search *models.Search) (*models.Role, error) {
	t := r.NewRole()
	//err := r.Found(search).First(t).Error
	//if !r.IsNotFound(err) {
	//	return t, err
	//}
	return t, nil
}

/**
 * 通过 id 获取 role 记录
 * @method GetRoleById
 * @param  {[type]}       role  *Role [description]
 */
//func GetRolesByIds(ids []int) ([]*Role, error) {
//	var roles []*Role
//	err := IsNotFound(libs.DB.Find(&roles, ids).Error)
//	if err != nil {
//		return nil, err
//	}
//	return roles, nil
//}

// DeleteRoleById del role by id
func (r *RoleRepo) DeleteRoleById(id uint) error {
	role := r.NewRole()
	role.ID = id
	err := r.DB.Delete(role).Error
	if err != nil {
		color.Red(fmt.Sprintf("DeleteRoleErr:%s \n", err))
		return err
	}

	return nil
}

// GetAllRoles get all roles
func (r *RoleRepo) GetAllRoles(s *models.Search) ([]*models.Role, int64, error) {
	var roles []*models.Role
	var count int64
	//all := r.GetAll(&models.Role{}, s)
	//all = all.Scopes(r.Relation(s.Relations))
	//if err := all.Count(&count).Error; err != nil {
	//	return nil, count, err
	//}
	//all = all.Scopes(r.Paginate(s.Offset, s.Limit))
	//if err := all.Find(&roles).Error; err != nil {
	//	return nil, count, err
	//}
	return roles, count, nil
}
