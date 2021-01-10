package handler

import (
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/aaronchen2k/tester/internal/server/utils"
	"github.com/kataras/iris/v12"
)

type MachineController struct {
	Ctx            iris.Context
	MachineService *service.MachineService `inject:""`
}

func NewMachineController() *MachineController {
	return &MachineController{}
}

func (c *MachineController) ListVm(ctx iris.Context) {
	rootNode := c.MachineService.ListVm()
	_, _ = ctx.JSON(agentUtils.ApiRes(200, "请求成功", rootNode))
}

func (c *MachineController) ListContainers(ctx iris.Context) {
	rootNode := c.MachineService.ListContainers()
	_, _ = ctx.JSON(agentUtils.ApiRes(200, "请求成功", rootNode))
}

func (c *MachineController) Get(ctx iris.Context) {

	_, _ = ctx.JSON(agentUtils.ApiRes(200, "请求成功", nil))
}
