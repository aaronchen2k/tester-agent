package handler

import (
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
)

type DeviceCtrl struct {
	Ctx     iris.Context
	Service *service.DeviceService `inject:""`
}

func NewDeviceCtrl() *DeviceCtrl {
	return &DeviceCtrl{}
}
func (g *DeviceCtrl) PostRegister() (result _domain.RpcResult) {
	var devices []_domain.DeviceInst
	if err := g.Ctx.ReadJSON(&devices); err != nil {
		_logUtils.Error(err.Error())
		result.Fail("wrong request data")
		return
	}

	result = g.Service.Register(devices)
	return result
}
