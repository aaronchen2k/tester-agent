package controller

import (
	"github.com/aaronchen2k/openstc/src/libs/common"
	"github.com/aaronchen2k/openstc/src/service"
	"github.com/kataras/iris/v12"
)

type InitController struct {
	SeederService *service.SeederService `inject:""`
}

func NewInitController() *InitController {
	return &InitController{}
}

func (c *InitController) InitData(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	c.SeederService.Run()
	c.SeederService.AddPerm()

	_, _ = ctx.JSON(common.ApiRes(200, "请求成功", nil))
}
