package handler

import (
	"fmt"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
)

type PlanController struct {
	Ctx         iris.Context
	PlanService *service.PlanService `inject:""`
	TaskService *service.TaskService `inject:""`
}

func NewPlanController() *PlanController {
	return &PlanController{}
}
func (g *PlanController) PostCreate() (result _domain.RpcResult) {
	var plan model.Plan
	if err := g.Ctx.ReadJSON(&plan); err != nil {
		_logUtils.Error(err.Error())
		result.Fail("wrong request data")
		return
	}

	po := g.PlanService.Save(plan)
	count := g.TaskService.GenerateFromPlan(po)
	result.Success(fmt.Sprintf("create %d tasks for plan %d.", count, po.ID))
	return result
}
