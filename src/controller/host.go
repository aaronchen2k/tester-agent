package controller

import (
	_domain "github.com/aaronchen2k/openstc-common/src/domain"
	_logUtils "github.com/aaronchen2k/openstc-common/src/libs/log"
	"github.com/aaronchen2k/openstc/src/libs/common"
	"github.com/aaronchen2k/openstc/src/service"
	"github.com/kataras/iris/v12"
)

type HostController struct {
	Ctx         iris.Context
	HostService *service.HostService `inject:""`
}

func NewHostController() *HostController {
	return &HostController{}
}
func (c *HostController) PostRegister() (result _domain.RpcResult) {
	var host _domain.Host
	if err := c.Ctx.ReadJSON(&host); err != nil {
		_logUtils.Error(err.Error())
		result.Fail("wrong request data")
		return
	}

	result = c.HostService.Register(host)
	return result
}

func (c *HostController) List(ctx iris.Context) {
	hosts := c.HostService.ListAll()

	_, _ = ctx.JSON(common.ApiResource(200, hosts, "请求成功"))
}

func (c *HostController) Get(ctx iris.Context) {

	_, _ = ctx.JSON(common.ApiResource(200, nil, "请求成功"))
}
