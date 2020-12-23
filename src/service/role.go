package service

import (
	"fmt"
	"github.com/aaronchen2k/openstc/src/middleware"
	"github.com/aaronchen2k/openstc/src/models"
	"github.com/aaronchen2k/openstc/src/repo"
	"github.com/aaronchen2k/openstc/src/transformer"
	"github.com/fatih/color"
	gf "github.com/snowlyg/gotransformer"
	"strconv"
	"time"
)

type RoleService struct {
	CommonService

	RoleRepo *repo.RoleRepo `inject:""`
	PermRepo *repo.PermRepo `inject:""`

	CasbinService *middleware.CasbinService `inject:""`
}

func NewRoleService() *RoleService {
	return &RoleService{}
}

// RolePermissions get role's permissions
func (s *RoleService) RolePermissions(role *models.Role) []*models.Permission {
	perms := s.GetPermissionsForUser(role.ID)
	var ps []*models.Permission
	for _, perm := range perms {
		if len(perm) >= 3 && len(perm[1]) > 0 && len(perm[2]) > 0 {
			search := &models.Search{
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
			p, err := s.PermRepo.GetPermission(search)
			if err == nil && p.ID > 0 {
				ps = append(ps, p)
			}
		}
	}
	return ps
}

func (s *RoleService) RolesTransform(roles []*models.Role) []*transformer.Role {
	var rs []*transformer.Role
	for _, role := range roles {
		r := s.RoleTransform(role)
		rs = append(rs, r)
	}
	return rs
}
func (s *RoleService) RoleTransform(role *models.Role) *transformer.Role {
	transformerRole := &transformer.Role{}
	g := gf.NewTransform(s, role, time.RFC3339)
	_ = g.Transformer()
	transformerRole.Perms = s.PermRepo.PermsTransform(s.RolePermissions(role))
	return transformerRole
}

// CreateRole create role
func (s *RoleService) CreateRole(role *models.Role) error {
	if err := s.RoleRepo.DB.Create(role).Error; err != nil {
		return err
	}

	s.addPerms(role.PermIds, role)

	return nil
}

// UpdateRole update role
func (s *RoleService) UpdateRole(id uint, nr *models.Role) error {
	if err := s.RoleRepo.Update(&models.Role{}, nr, id); err != nil {
		return err
	}

	s.addPerms(nr.PermIds, nr)

	return nil
}

func (s *RoleService) GetRolesForUser(uid uint) []string {
	uids, err := s.CasbinService.Enforcer.GetRolesForUser(strconv.FormatUint(uint64(uid), 10))
	if err != nil {
		color.Red(fmt.Sprintf("GetRolesForUser 错误: %v", err))
		return []string{}
	}

	return uids
}

// addPerms add perms
func (s *RoleService) addPerms(permIds []uint, role *models.Role) {
	if len(permIds) > 0 {
		roleId := strconv.FormatUint(uint64(role.ID), 10)
		if _, err := s.CasbinService.Enforcer.DeletePermissionsForUser(roleId); err != nil {
			color.Red(fmt.Sprintf("AppendPermsErr:%s \n", err))
		}
		var perms []models.Permission
		s.RoleRepo.DB.Where("id in (?)", permIds).Find(&perms)
		for _, perm := range perms {
			if _, err := s.CasbinService.Enforcer.AddPolicy(roleId, perm.Name, perm.Act); err != nil {
				color.Red(fmt.Sprintf("AddPolicy:%s \n", err))
			}
		}
	} else {
		color.Yellow(fmt.Sprintf("没有角色：%s 权限为空 \n", role.Name))
	}
}
