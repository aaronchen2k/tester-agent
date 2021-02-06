package handler

import (
	"github.com/aaronchen2k/tester/internal/pkg/utils"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
)

type TaskCtrl struct {
	Ctx          iris.Context
	TaskService  *service.TaskService  `inject:""`
	QueueService *service.QueueService `inject:""`
}

func NewTaskCtrl() *TaskCtrl {
	return &TaskCtrl{}
}

func (c *TaskCtrl) Get(ctx iris.Context) {
	_, _ = ctx.JSON(_utils.ApiRes(200, "请求成功", ""))
}
