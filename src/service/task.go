package service

import (
	_const "github.com/aaronchen2k/openstc-common/src/libs/const"
	"github.com/aaronchen2k/openstc/src/model"
	"github.com/aaronchen2k/openstc/src/repo"
)

type TaskService struct {
	TaskRepo  *repo.TaskRepo  `inject:""`
	QueueRepo *repo.QueueRepo `inject:""`
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (s *TaskService) Save(task model.Task) model.Task {
	s.TaskRepo.Save(&task)
	return task
}

func (s *TaskService) SetProgress(id uint, progress _const.BuildProgress) {
	s.TaskRepo.SetProgress(id, progress)
}

func (s *TaskService) CheckCompleted(taskId uint) {
	queues := s.QueueRepo.QueryByTask(taskId)

	progress := _const.ProgressCompleted
	status := _const.StatusPass
	isAllQueuesCompleted := true

	for _, queue := range queues {
		if queue.Progress != _const.ProgressCompleted && queue.Progress != _const.ProgressTimeout { // 有queue在进行中
			isAllQueuesCompleted = false
			break
		}

		if queue.Progress == _const.ProgressTimeout { // 有一个超时，就超时
			progress = _const.ProgressTimeout
		}

		if queue.Status == _const.StatusFail { // 有一个失败，就失败
			status = _const.StatusFail
		}
	}

	if isAllQueuesCompleted {
		s.TaskRepo.SetResult(taskId, progress, status)
	}
}