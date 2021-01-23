package handler

import (
	"github.com/aaronchen2k/tester/internal/pkg/utils"
	"github.com/aaronchen2k/tester/internal/server/service"
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

	_, _ = ctx.JSON(utils.ApiRes(200, "请求成功", nil))
}
