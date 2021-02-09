package handler

import (
	"github.com/aaronchen2k/tester/internal/pkg/utils"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
	"strconv"
)

type ClusterCtrl struct {
	Ctx            iris.Context
	ClusterService *service.ClusterService `inject:""`
}

func NewClusterCtrl() *ClusterCtrl {
	return &ClusterCtrl{}
}

func (c *ClusterCtrl) List(ctx iris.Context) {
	keywords := ctx.FormValue("keywords")
	pageNoStr := ctx.FormValue("pageNo")
	pageSizeStr := ctx.FormValue("pageSize")

	pageNo, _ := strconv.Atoi(pageNoStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	clusters, total := c.ClusterService.ListAll(keywords, pageNo, pageSize)

	_, _ = ctx.JSON(_utils.ApiResPage(200, "请求成功",
		clusters, pageNo, pageSize, total))
}

func (c *ClusterCtrl) Get(ctx iris.Context) {
	_, _ = ctx.JSON(_utils.ApiRes(200, "请求成功", nil))
}
