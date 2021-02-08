package handler

import (
	"fmt"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
)

type VmCtrl struct {
	Ctx       iris.Context
	VmService *service.VmService `inject:""`
}

func NewVmCtrl() *VmCtrl {
	return &VmCtrl{}
}

func (c *VmCtrl) Register() (result _domain.RpcResult) {
	var vm _domain.Vm
	if err := c.Ctx.ReadJSON(&vm); err != nil {
		_logUtils.Error(err.Error())
		result.Fail("wrong request data")
		return
	}

	po := c.VmService.Register(vm)

	result.Success(fmt.Sprintf("succes to register host %d.", po))
	return result
}
