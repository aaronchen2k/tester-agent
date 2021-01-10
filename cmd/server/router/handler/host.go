package handler

import (
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/aaronchen2k/tester/internal/server/utils"
	"github.com/kataras/iris/v12"
	"strconv"
)

type HostController struct {
	Ctx         iris.Context
	HostService *service.ClusterService `inject:""`
}

func NewHostController() *HostController {
	return &HostController{}
}

func (c *HostController) List(ctx iris.Context) {
	keywords := ctx.FormValue("keywords")
	pageNoStr := ctx.FormValue("pageNo")
	pageSizeStr := ctx.FormValue("pageSize")

	pageNo, _ := strconv.Atoi(pageNoStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	hosts, total := c.HostService.ListAll(keywords, pageNo, pageSize)

	_, _ = ctx.JSON(agentUtils.ApiResPage(200, "请求成功",
		hosts, pageNo, pageSize, total))
}

func (c *HostController) Get(ctx iris.Context) {

	_, _ = ctx.JSON(agentUtils.ApiRes(200, "请求成功", nil))
}
