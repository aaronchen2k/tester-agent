package handler

import (
	"github.com/aaronchen2k/tester/internal/pkg/utils"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
	"strconv"
)

type HostCtrl struct {
	Ctx         iris.Context
	HostService *service.ClusterService `inject:""`
}

func NewHostCtrl() *HostCtrl {
	return &HostCtrl{}
}

func (c *HostCtrl) List(ctx iris.Context) {
	keywords := ctx.FormValue("keywords")
	pageNoStr := ctx.FormValue("pageNo")
	pageSizeStr := ctx.FormValue("pageSize")

	pageNo, _ := strconv.Atoi(pageNoStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	hosts, total := c.HostService.ListAll(keywords, pageNo, pageSize)

	_, _ = ctx.JSON(_utils.ApiResPage(200, "请求成功",
		hosts, pageNo, pageSize, total))
}

func (c *HostCtrl) Get(ctx iris.Context) {

	_, _ = ctx.JSON(_utils.ApiRes(200, "请求成功", nil))
}
