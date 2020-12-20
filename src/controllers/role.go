package controllers

import (
	"github.com/aaronchen2k/openstc/src/libs/common"
	"github.com/aaronchen2k/openstc/src/models"
	"github.com/aaronchen2k/openstc/src/repo"
	"github.com/aaronchen2k/openstc/src/validates"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type RoleController struct {
	BaseController

	userRepo *repo.UserRepo
	roleRepo *repo.RoleRepo
	permRepo *repo.PermRepo
}

func NewRoleController(userRepo *repo.UserRepo, roleRepo *repo.RoleRepo, permRepo *repo.PermRepo) *RoleController {
	return &RoleController{userRepo: userRepo, roleRepo: roleRepo, permRepo: permRepo}
}

/**
* @api {get} /admin/roles/:id 根据id获取角色信息
* @apiName 根据id获取角色信息
* @apiGroup Roles
* @apiVersion 1.0.0
* @apiDescription 根据id获取角色信息
* @apiSampleRequest /admin/roles/:id
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission
 */
func (c *RoleController) GetRole(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	id, _ := ctx.Params().GetUint("id")
	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "id",
				Condition: "=",
				Value:     id,
			},
		},
	}
	role, err := c.roleRepo.GetRole(s)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}

	rr := c.roleRepo.RoleTransform(role)
	rr.Perms = c.permRepo.PermsTransform(c.roleRepo.RolePermissions(role))
	_, _ = ctx.JSON(common.ApiResource(200, rr, "操作成功"))
}

/**
* @api {post} /admin/roles/ 新建角色
* @apiName 新建角色
* @apiGroup Roles
* @apiVersion 1.0.0
* @apiDescription 新建角色
* @apiSampleRequest /admin/roles/
* @apiParam {string} name 角色名
* @apiParam {string} display_name
* @apiParam {string} description
* @apiParam {string} level
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *RoleController) CreateRole(ctx iris.Context) {

	ctx.StatusCode(iris.StatusOK)
	role := new(models.Role)

	if err := ctx.ReadJSON(role); err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}

	err := validates.Validate.Struct(*role)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validates.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(common.ApiResource(400, nil, e))
				return
			}
		}
	}

	err = c.roleRepo.CreateRole(role)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}
	if role.ID == 0 {
		_, _ = ctx.JSON(common.ApiResource(400, nil, "操作失败"))
		return
	}
	_, _ = ctx.JSON(common.ApiResource(200, c.roleRepo.RoleTransform(role), "操作成功"))

}

/**
* @api {post} /admin/roles/:id/update 更新角色
* @apiName 更新角色
* @apiGroup Roles
* @apiVersion 1.0.0
* @apiDescription 更新角色
* @apiSampleRequest /admin/roles/:id/update
* @apiParam {string} name 角色名
* @apiParam {string} display_name
* @apiParam {string} description
* @apiParam {string} level
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *RoleController) UpdateRole(ctx iris.Context) {

	ctx.StatusCode(iris.StatusOK)
	role := new(models.Role)
	if err := ctx.ReadJSON(role); err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}

	err := validates.Validate.Struct(*role)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validates.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(common.ApiResource(400, nil, e))
				return
			}
		}
	}

	id, _ := ctx.Params().GetUint("id")
	err = c.roleRepo.UpdateRole(id, role)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}
	_, _ = ctx.JSON(common.ApiResource(200, c.roleRepo.RoleTransform(role), "操作成功"))

}

/**
* @api {delete} /admin/roles/:id/delete 删除角色
* @apiName 删除角色
* @apiGroup Roles
* @apiVersion 1.0.0
* @apiDescription 删除角色
* @apiSampleRequest /admin/roles/:id/delete
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *RoleController) DeleteRole(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	id, _ := ctx.Params().GetUint("id")

	err := c.roleRepo.DeleteRoleById(id)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}

	_, _ = ctx.JSON(common.ApiResource(200, nil, "删除成功"))
}

/**
* @api {get} /roles 获取所有的角色
* @apiName 获取所有的角色
* @apiGroup Roles
* @apiVersion 1.0.0
* @apiDescription 获取所有的角色
* @apiSampleRequest /roles
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *RoleController) GetAllRoles(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	s := c.GetCommonListSearch(ctx)
	roles, count, err := c.roleRepo.GetAllRoles(s)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}

	transform := c.roleRepo.RolesTransform(roles)
	_, _ = ctx.JSON(common.ApiResource(200, map[string]interface{}{"items": transform, "total": count, "limit": s.Limit}, "操作成功"))
}
