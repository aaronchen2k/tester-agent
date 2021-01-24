package service

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/model/base"
	"github.com/aaronchen2k/tester/internal/server/repo"
	"strings"
)

type QueueService struct {
	TaskService *TaskService `inject:""`

	QueueRepo  *repo.QueueRepo  `inject:""`
	DeviceRepo *repo.DeviceRepo `inject:""`
}

func NewQueueService() *QueueService {
	return &QueueService{}
}

func (s *QueueService) GenerateFromTask(task model.Task) {
	if task.GroupId == 0 {
		task.GroupId = task.ID
	}

	if task.BuildType == _const.AppiumTest {
		s.GenerateAppiumQueue(task)
	} else if task.BuildType == _const.SeleniumTest {
		s.GenerateSeleniumQueue(task)
	}

	return
}

func (s *QueueService) GenerateAppiumQueue(task model.Task) {
	env := task.TestEnv
	serial := strings.TrimSpace(env.Serial)

	device := model.Device{}
	if serial == "" {
		device = s.DeviceRepo.GetBySerial(serial)
	} else {
		device = s.getDeviceFromEnv(env)
	}

	if device.ID == 0 { // not found
		return
	}

	env.DeviceId = device.ID
	env.Serial = device.Serial

	if device.ID != 0 {
		queue := model.NewQueueDetail(
			task.BuildType, task.Priority, task.GroupId, task.ID,
			task.TaskName, task.UserName,
			env, task.TestObject)

		s.QueueRepo.Save(&queue)
	}

	return
}

func (s *QueueService) GenerateSeleniumQueue(task model.Task) {
	env := task.TestEnv

	vmTempl := s.getVmTemplFromEnv(env)
	env.VmTemplId = vmTempl.ID

	queue := model.NewQueueDetail(
		task.BuildType, task.Priority, task.GroupId, task.ID,
		task.TaskName, task.UserName,
		env, task.TestObject)

	s.QueueRepo.Save(&queue)

	return
}

func (s *QueueService) SetQueueResult(queueId uint, progress _const.BuildProgress, status _const.BuildStatus) {
	queue := s.QueueRepo.GetQueue(queueId)

	s.QueueRepo.SetQueueStatus(queueId, progress, status)
	s.TaskService.CheckCompleted(queue.TaskId)
}

func (s *QueueService) getDeviceFromEnv(env base.TestEnv) (dev model.Device) {
	if env.DeviceId != 0 {
		dev := s.DeviceRepo.Get(env.DeviceId)
		if dev.ID > 0 {
			return dev
		}
	}

	if env.Serial != "" {
		dev := s.DeviceRepo.GetBySerial(env.Serial)
		if dev.ID > 0 {
			return dev
		}
	}

	dev = s.DeviceRepo.GetByEnv(env)

	return
}

func (s *QueueService) getVmTemplFromEnv(env base.TestEnv) (dev model.VmTempl) {

	return
}
