package handler

import (
	"github.com/aaronchen2k/tester/internal/pkg/utils"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
)

type MachineCtrl struct {
	Ctx            iris.Context
	MachineService *service.VirtualService `inject:""`
}

func NewMachineCtrl() *MachineCtrl {
	return &MachineCtrl{}
}

func (c *MachineCtrl) ListVm(ctx iris.Context) {
	rootNode := c.MachineService.ListVm()
	_, _ = ctx.JSON(_utils.ApiRes(200, "请求成功", rootNode))
}

func (c *MachineCtrl) ListContainers(ctx iris.Context) {
	rootNode := c.MachineService.ListContainers()
	_, _ = ctx.JSON(_utils.ApiRes(200, "请求成功", rootNode))
}

func (c *MachineCtrl) Get(ctx iris.Context) {

	_, _ = ctx.JSON(_utils.ApiRes(200, "请求成功", nil))
}
