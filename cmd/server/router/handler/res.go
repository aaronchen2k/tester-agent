package handler

import (
	"github.com/aaronchen2k/tester/internal/pkg/utils"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
)

type ResCtrl struct {
	Ctx        iris.Context
	ResService *service.ResService `inject:""`
}

func NewMachineCtrl() *ResCtrl {
	return &ResCtrl{}
}

func (c *ResCtrl) ListVm(ctx iris.Context) {
	rootNode := c.ResService.ListVm()
	_, _ = ctx.JSON(_utils.ApiRes(200, "请求成功", rootNode))
}

func (c *ResCtrl) ListContainer(ctx iris.Context) {
	rootNode := c.ResService.ListContainers()
	_, _ = ctx.JSON(_utils.ApiRes(200, "请求成功", rootNode))
}

func (c *ResCtrl) Get(ctx iris.Context) {

	_, _ = ctx.JSON(_utils.ApiRes(200, "请求成功", nil))
}
