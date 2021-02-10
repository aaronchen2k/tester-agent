package handler

import (
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	_utils "github.com/aaronchen2k/tester/internal/pkg/utils"
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

func (c *VmCtrl) Register(ctx iris.Context) {
	var vm _domain.Vm
	if err := c.Ctx.ReadJSON(&vm); err != nil {
		_logUtils.Error(err.Error())
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), ""))
		return
	}

	c.VmService.Register(vm)

	_, _ = ctx.JSON(_utils.ApiRes(200, "", ""))
	return
}
