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

	QueueRepo       *repo.QueueRepo          `inject:""`
	DeviceRepo      *repo.DeviceRepo         `inject:""`
	VmTemplRepo     *repo.VmTemplRepo        `inject:""`
	DockerImageRepo *repo.ContainerImageRepo `inject:""`
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
	} else if task.BuildType == _const.UnitTest {
		s.GenerateUnitTestQueue(task)
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
		device = s.getDeviceByEnv(env)
	}

	if device.ID == 0 { // not found
		return
	}

	env.DeviceId = device.ID
	env.Serial = device.Serial

	if device.ID != 0 {
		queue := model.NewQueueDetail(
			_const.AppiumTest, task.Priority, task.GroupId, task.ID,
			task.TaskName, task.UserName,
			env, task.TestObject)

		s.QueueRepo.Save(&queue)
	}

	return
}

func (s *QueueService) GenerateSeleniumQueue(task model.Task) {
	env := task.TestEnv

	vmTempl := s.getVmTemplByEnv(env)
	if vmTempl.ID == 0 {
		return
	}

	env.VmTemplId = vmTempl.ID

	queue := model.NewQueueDetail(
		_const.SeleniumTest, task.Priority, task.GroupId, task.ID,
		task.TaskName, task.UserName,
		env, task.TestObject)

	s.QueueRepo.Save(&queue)

	return
}

func (s *QueueService) GenerateUnitTestQueue(task model.Task) {
	env := task.TestEnv

	vmTempl := s.getVmTemplByEnv(env)
	env.VmTemplId = vmTempl.ID

	queue := model.NewQueueDetail(
		_const.UnitTest, task.Priority, task.GroupId, task.ID,
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

func (s *QueueService) getDeviceByEnv(env base.TestEnv) (dev model.Device) {
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

func (s *QueueService) getVmTemplByEnv(env base.TestEnv) (templ model.VmTempl) {
	if env.VmTemplId != 0 {
		templ := s.VmTemplRepo.Get(env.VmTemplId)
		if templ.ID > 0 {
			return templ
		}
	}

	templ = s.VmTemplRepo.GetByEnv(env)

	return
}

func (s *QueueService) getDockerImageByEnv(env base.TestEnv) (image model.ContainerImage) {
	image = s.DockerImageRepo.GetByEnv(env)

	return
}
