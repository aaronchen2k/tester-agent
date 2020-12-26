package controller

import (
	_domain "github.com/aaronchen2k/openstc-common/src/domain"
	_logUtils "github.com/aaronchen2k/openstc-common/src/libs/log"
	"github.com/aaronchen2k/openstc/src/service"
	"github.com/kataras/iris/v12"
)

type HostController struct {
	Ctx         iris.Context
	HostService *service.HostService `inject:""`
}

func NewHostController() *HostController {
	return &HostController{HostService: service.NewHostService()}
}
func (g *HostController) PostRegister() (result _domain.RpcResult) {
	var host _domain.Host
	if err := g.Ctx.ReadJSON(&host); err != nil {
		_logUtils.Error(err.Error())
		result.Fail("wrong request data")
		return
	}

	result = g.HostService.Register(host)
	return result
}
