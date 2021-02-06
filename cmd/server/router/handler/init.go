package handler

import (
	"github.com/aaronchen2k/tester/internal/pkg/utils"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
)

type InitCtrl struct {
	SeederService *service.SeederService `inject:""`
}

func NewInitCtrl() *InitCtrl {
	return &InitCtrl{}
}

func (c *InitCtrl) InitData(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	c.SeederService.Run()

	_, _ = ctx.JSON(_utils.ApiRes(200, "请求成功", nil))
}
