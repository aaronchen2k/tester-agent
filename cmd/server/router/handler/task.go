package handler

import (
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
