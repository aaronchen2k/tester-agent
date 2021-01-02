package handler

import (
	_domain "github.com/aaronchen2k/openstc/internal/pkg/domain"
	_logUtils "github.com/aaronchen2k/openstc/internal/pkg/libs/log"
	"github.com/aaronchen2k/openstc/internal/server/service"
	"github.com/kataras/iris/v12"
)

type DeviceController struct {
	Ctx     iris.Context
	Service *service.DeviceService `inject:""`
}

func NewDeviceController() *DeviceController {
	return &DeviceController{}
}
func (g *DeviceController) PostRegister() (result _domain.RpcResult) {
	var devices []_domain.DeviceInst
	if err := g.Ctx.ReadJSON(&devices); err != nil {
		_logUtils.Error(err.Error())
		result.Fail("wrong request data")
		return
	}

	result = g.Service.Register(devices)
	return result
}
