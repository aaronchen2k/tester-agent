package service

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/repo"
)

type ExecService struct {
	ResService      *ResService      `inject:""`
	DeviceService   *DeviceService   `inject:""`
	VmService       *VmService       `inject:""`
	AppiumService   *AppiumService   `inject:""`
	SeleniumService *SeleniumService `inject:""`
	TaskService     *TaskService     `inject:""`

	ExecRepo   *repo.ExecRepo   `inject:""`
	QueueRepo  *repo.QueueRepo  `inject:""`
	DeviceRepo *repo.DeviceRepo `inject:""`
	VmRepo     *repo.VmRepo     `inject:""`
	TaskRepo   *repo.TaskRepo   `inject:""`
}

func NewExecService() *ExecService {
	return &ExecService{}
}

func (s *ExecService) CheckAll() {
	s.CheckExec()
	s.SetTimeout()
	s.RetryTimeoutOrFailed()
}

func (s *ExecService) CheckExec() {
	queuesToBuild := s.QueueRepo.QueryForExec()
	for _, queue := range queuesToBuild {
		s.CheckAndCall(queue)
	}
}

func (s *ExecService) CheckAndCall(queue model.Queue) {
	if queue.BuildType == _const.SeleniumTest {
		s.CheckAndCallSeleniumTest(queue)
	} else if queue.BuildType == _const.AppiumTest {
		s.CheckAndCallAppiumTest(queue)
	}
}

func (s *ExecService) CheckAndCallSeleniumTest(queue model.Queue) {
	originalProgress := queue.Progress
	var newProgress _const.BuildProgress

	if queue.Progress == _const.ProgressCreated {
		// create kvm
		err := s.VmService.CreateByQueue(queue)
		if err == nil { // success to create
			newProgress = _const.ProgressInProgress
		} else {
			newProgress = _const.ProgressPending
		}

	} else if queue.Progress == _const.ProgressLaunchVm {
		vmId := queue.VmId
		vm := s.VmRepo.GetById(vmId)

		if vm.Status == _const.VmActive { // find ready vm, begin to run test
			result := s.SeleniumService.Start(queue)
			if result.IsSuccess() {
				s.QueueRepo.Start(queue)
				newProgress = _const.ProgressInProgress
			} else { // busy, pending
				s.QueueRepo.Pending(queue.ID)
				newProgress = _const.ProgressPending
			}
		}
	}

	if originalProgress != newProgress { // queue's progress changed
		s.TaskRepo.SetProgress(queue.TaskId, newProgress)
	}
}

func (s *ExecService) CheckAndCallAppiumTest(queue model.Queue) {
	serial := queue.Serial
	device := s.DeviceRepo.GetBySerial(serial)

	originalProgress := queue.Progress
	var newProgress _const.BuildProgress

	if s.DeviceService.IsDeviceReady(device) {
		rpcResult := s.AppiumService.Start(queue)

		if rpcResult.IsSuccess() {
			s.QueueRepo.Start(queue) // start
			newProgress = _const.ProgressInProgress
		} else {
			s.QueueRepo.Pending(queue.ID) // pending
			newProgress = _const.ProgressPending
		}
	} else {
		s.QueueRepo.Pending(queue.ID) // pending
		newProgress = _const.ProgressPending
	}

	if originalProgress != newProgress { // progress changed
		s.TaskService.SetProgress(queue.TaskId, newProgress)
	}
}

func (s *ExecService) SetTimeout() {
	queues := s.QueueRepo.QueryTimeout()

	for _, queue := range queues {
		s.QueueRepo.SetTimeout(queue.ID)
	}
}

func (s *ExecService) RetryTimeoutOrFailed() {
	queues := s.QueueRepo.QueryTimeoutOrFailedForRetry()

	for _, queue := range queues {
		s.CheckAndCall(queue)
	}
}
