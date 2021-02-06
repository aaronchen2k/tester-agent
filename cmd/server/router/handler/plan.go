package handler

import (
	"github.com/aaronchen2k/tester/internal/pkg/utils"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/service"
	serverConst "github.com/aaronchen2k/tester/internal/server/utils/const"
	"github.com/kataras/iris/v12"
	"strconv"
)

type PlanCtrl struct {
	BaseCtrl

	PlanService *service.PlanService `inject:""`
	TaskService *service.TaskService `inject:""`
}

func NewPlanCtrl() *PlanCtrl {
	return &PlanCtrl{}
}

func (c *PlanCtrl) List(ctx iris.Context) {
	keywords := ctx.FormValue("keywords")
	pageNoStr := ctx.FormValue("pageNo")
	pageSizeStr := ctx.FormValue("pageSize")

	pageNo, _ := strconv.Atoi(pageNoStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	if pageSize == 0 {
		pageSize = serverConst.PageSize
	}

	plans, total := c.PlanService.List(keywords, pageNo, pageSize)

	_, _ = ctx.JSON(_utils.ApiResPage(200, "请求成功",
		plans, pageNo, pageSize, total))
}

func (c *PlanCtrl) Get(ctx iris.Context) {

}

func (c *PlanCtrl) Create(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	plan := new(model.Plan)
	if err := ctx.ReadJSON(plan); err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	if c.Validate(*plan, ctx) {
		return
	}

	err := c.PlanService.Save(plan)
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", plan))
	return
}

func (c *PlanCtrl) Update(ctx iris.Context) {

}

func (c *PlanCtrl) Delete(ctx iris.Context) {

}
