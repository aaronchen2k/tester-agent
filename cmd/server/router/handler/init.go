package handler

import (
	"github.com/aaronchen2k/openstc/internal/server/service"
	"github.com/aaronchen2k/openstc/internal/server/utils"
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

	_, _ = ctx.JSON(agentUtils.ApiRes(200, "请求成功", nil))
}
