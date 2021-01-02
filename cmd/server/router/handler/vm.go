package handler

import (
	"fmt"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
)

type VmController struct {
	Ctx          iris.Context
	VmService    *service.VmService    `inject:""`
	ImageService *service.ImageService `inject:""`
	HostService  *service.HostService  `inject:""`
}

func NewVmController() *VmController {
	return &VmController{}
}

func (g *VmController) PostRegister() (result _domain.RpcResult) {
	var vm _domain.Vm
	if err := g.Ctx.ReadJSON(&vm); err != nil {
		_logUtils.Error(err.Error())
		result.Fail("wrong request data")
		return
	}

	po := g.VmService.Register(vm)

	result.Success(fmt.Sprintf("succes to register host %d.", po))
	return result
}
