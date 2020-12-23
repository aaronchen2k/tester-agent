package controllers

import (
	"fmt"
	"github.com/aaronchen2k/openstc/src/libs/common"
	"github.com/aaronchen2k/openstc/src/models"
	"github.com/aaronchen2k/openstc/src/repo"
	"github.com/aaronchen2k/openstc/src/validates"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type PermController struct {
	UserRepo *repo.UserRepo `inject:""`
	PermRepo *repo.PermRepo `inject:""`
}

func NewPermController() *PermController {
	return &PermController{}
}

/**
* @api {get} /admin/permissions/:id 根据id获取权限信息
* @apiName 根据id获取权限信息
* @apiGroup Permissions
* @apiVersion 1.0.0
* @apiDescription 根据id获取权限信息
* @apiSampleRequest /admin/permissions/:id
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission
 */
func (c *PermController) GetPermission(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	//id, _ := ctx.Params().GetUint("id")

	perm, err := c.PermRepo.GetPermission(nil)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}

	_, _ = ctx.JSON(common.ApiResource(200, c.PermRepo.PermTransform(perm), "操作成功"))
}

/**
* @api {post} /admin/permissions/ 新建权限
* @apiName 新建权限
* @apiGroup Permissions
* @apiVersion 1.0.0
* @apiDescription 新建权限
* @apiSampleRequest /admin/permissions/
* @apiParam {string} name 权限名
* @apiParam {string} display_name
* @apiParam {string} description
* @apiParam {string} level
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *PermController) CreatePermission(ctx iris.Context) {

	ctx.StatusCode(iris.StatusOK)
	perm := new(models.Permission)
	if err := ctx.ReadJSON(perm); err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}
	err := validates.Validate.Struct(*perm)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validates.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(common.ApiResource(400, nil, e))
				return
			}
		}
	}

	err = c.PermRepo.CreatePermission(perm)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, fmt.Sprintf("Error create prem: %s", err.Error())))
		return
	}

	if perm.ID == 0 {
		_, _ = ctx.JSON(common.ApiResource(400, perm, "操作失败"))
		return
	}
	_, _ = ctx.JSON(common.ApiResource(200, c.PermRepo.PermTransform(perm), "操作成功"))

}

/**
* @api {post} /admin/permissions/:id/update 更新权限
* @apiName 更新权限
* @apiGroup Permissions
* @apiVersion 1.0.0
* @apiDescription 更新权限
* @apiSampleRequest /admin/permissions/:id/update
* @apiParam {string} name 权限名
* @apiParam {string} display_name
* @apiParam {string} description
* @apiParam {string} level
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *PermController) UpdatePermission(ctx iris.Context) {

	ctx.StatusCode(iris.StatusOK)
	aul := new(models.Permission)

	if err := ctx.ReadJSON(aul); err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}
	err := validates.Validate.Struct(*aul)
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
	err = c.PermRepo.UpdatePermission(id, aul)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, fmt.Sprintf("Error update prem: %s", err.Error())))
		return
	}

	_, _ = ctx.JSON(common.ApiResource(200, c.PermRepo.PermTransform(aul), "操作成功"))

}

/**
* @api {delete} /admin/permissions/:id/delete 删除权限
* @apiName 删除权限
* @apiGroup Permissions
* @apiVersion 1.0.0
* @apiDescription 删除权限
* @apiSampleRequest /admin/permissions/:id/delete
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *PermController) DeletePermission(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	id, _ := ctx.Params().GetUint("id")
	err := c.PermRepo.DeletePermissionById(id)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}
	_, _ = ctx.JSON(common.ApiResource(200, nil, "删除成功"))
}

/**
* @api {get} /permissions 获取所有的权限
* @apiName 获取所有的权限
* @apiGroup Permissions
* @apiVersion 1.0.0
* @apiDescription 获取所有的权限
* @apiSampleRequest /permissions
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *PermController) GetAllPermissions(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	permissions, count, err := c.PermRepo.GetAllPermissions(nil)
	if err != nil {
		_, _ = ctx.JSON(common.ApiResource(400, nil, err.Error()))
		return
	}

	transform := c.PermRepo.PermsTransform(permissions)
	_, _ = ctx.JSON(common.ApiResource(200, map[string]interface{}{"items": transform, "total": count, "limit": "s.Limit"}, "操作成功"))

}
