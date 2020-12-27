package controller

import (
	_domain "github.com/aaronchen2k/openstc-common/src/domain"
	_logUtils "github.com/aaronchen2k/openstc-common/src/libs/log"
	"github.com/aaronchen2k/openstc/src/libs/common"
	"github.com/aaronchen2k/openstc/src/service"
	"github.com/kataras/iris/v12"
	"strconv"
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
	keywords := ctx.FormValue("keywords")
	pageNoStr := ctx.FormValue("pageNo")
	pageSizeStr := ctx.FormValue("pageSize")

	pageNo, _ := strconv.Atoi(pageNoStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	hosts, total := c.HostService.ListAll(keywords, pageNo, pageSize)

	_, _ = ctx.JSON(common.ApiResPage(200, "请求成功",
		hosts, pageNo, pageSize, total))
}

func (c *HostController) Get(ctx iris.Context) {

	_, _ = ctx.JSON(common.ApiRes(200, "请求成功", nil))
}
