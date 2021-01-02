package handler

import (
	"fmt"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
)

type TaskController struct {
	Ctx          iris.Context
	TaskService  *service.TaskService  `inject:""`
	QueueService *service.QueueService `inject:""`
}

func NewTaskController() *TaskController {
	return &TaskController{}
}
func (g *TaskController) PostCreate() (result _domain.RpcResult) {
	var task model.Task
	if err := g.Ctx.ReadJSON(&task); err != nil {
		_logUtils.Error(err.Error())
		result.Fail("wrong request data")
		return
	}

	po := g.TaskService.Save(task)
	count := g.QueueService.GenerateFromTask(po)
	result.Success(fmt.Sprintf("create %d queues for task %d.", count, po.ID))
	return result
}
