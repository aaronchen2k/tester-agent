package repo

import (
	"fmt"
	"github.com/aaronchen2k/openstc/src/libs/casbin"
	"github.com/aaronchen2k/openstc/src/libs/db"
	"github.com/aaronchen2k/openstc/src/models"
	"github.com/aaronchen2k/openstc/src/transformer"
	gf "github.com/snowlyg/gotransformer"
	"strconv"
	"time"

	"github.com/fatih/color"
	"gorm.io/gorm"
)

type RoleRepo struct {
	BaseRepo
	permRepo *PermRepo
}

func NewRoleRepo(permRepo *PermRepo) *RoleRepo {
	return &RoleRepo{permRepo: permRepo}
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
	err := r.Found(search).First(t).Error
	if !r.IsNotFound(err) {
		return t, err
	}
	return t, nil
}

/**
 * 通过 id 获取 role 记录
 * @method GetRoleById
 * @param  {[type]}       role  *Role [description]
 */
//func GetRolesByIds(ids []int) ([]*Role, error) {
//	var roles []*Role
//	err := IsNotFound(libs.Db.Find(&roles, ids).Error)
//	if err != nil {
//		return nil, err
//	}
//	return roles, nil
//}

// DeleteRoleById del role by id
func (r *RoleRepo) DeleteRoleById(id uint) error {
	role := r.NewRole()
	role.ID = id
	err := db.Db.Delete(role).Error
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
	all := r.GetAll(&models.Role{}, s)
	all = all.Scopes(r.Relation(s.Relations))
	if err := all.Count(&count).Error; err != nil {
		return nil, count, err
	}
	all = all.Scopes(r.Paginate(s.Offset, s.Limit))
	if err := all.Find(&roles).Error; err != nil {
		return nil, count, err
	}
	return roles, count, nil
}

// CreateRole create role
func (r *RoleRepo) CreateRole(role *models.Role) error {
	if err := db.Db.Create(role).Error; err != nil {
		return err
	}

	r.addPerms(role.PermIds, role)

	return nil
}

// addPerms add perms
func (r *RoleRepo) addPerms(permIds []uint, role *models.Role) {
	if len(permIds) > 0 {
		roleId := strconv.FormatUint(uint64(role.ID), 10)
		if _, err := casbinUtils.Enforcer.DeletePermissionsForUser(roleId); err != nil {
			color.Red(fmt.Sprintf("AppendPermsErr:%s \n", err))
		}
		var perms []models.Permission
		db.Db.Where("id in (?)", permIds).Find(&perms)
		for _, perm := range perms {
			if _, err := casbinUtils.Enforcer.AddPolicy(roleId, perm.Name, perm.Act); err != nil {
				color.Red(fmt.Sprintf("AddPolicy:%s \n", err))
			}
		}
	} else {
		color.Yellow(fmt.Sprintf("没有角色：%s 权限为空 \n", role.Name))
	}
}

// UpdateRole update role
func (r *RoleRepo) UpdateRole(id uint, nr *models.Role) error {
	if err := r.Update(&models.Role{}, nr, id); err != nil {
		return err
	}

	r.addPerms(nr.PermIds, nr)

	return nil
}

// RolePermissions get role's permissions
func (r *RoleRepo) RolePermissions(role *models.Role) []*models.Permission {
	perms := r.GetPermissionsForUser(role.ID)
	var ps []*models.Permission
	for _, perm := range perms {
		if len(perm) >= 3 && len(perm[1]) > 0 && len(perm[2]) > 0 {
			s := &models.Search{
				Fields: []*models.Filed{
					{
						Key:       "name",
						Condition: "=",
						Value:     perm[1],
					},
					{
						Key:       "act",
						Condition: "=",
						Value:     perm[2],
					},
				},
			}
			p, err := r.permRepo.GetPermission(s)
			if err == nil && p.ID > 0 {
				ps = append(ps, p)
			}
		}
	}
	return ps
}

func (r *RoleRepo) RolesTransform(roles []*models.Role) []*transformer.Role {
	var rs []*transformer.Role
	for _, role := range roles {
		r := r.RoleTransform(role)
		rs = append(rs, r)
	}
	return rs
}
func (r *RoleRepo) RoleTransform(role *models.Role) *transformer.Role {
	transformerRole := &transformer.Role{}
	g := gf.NewTransform(r, role, time.RFC3339)
	_ = g.Transformer()
	transformerRole.Perms = r.permRepo.PermsTransform(r.RolePermissions(role))
	return transformerRole
}
