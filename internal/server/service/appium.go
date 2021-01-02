package service

import (
	_domain "github.com/aaronchen2k/openstc/internal/pkg/domain"
	_const "github.com/aaronchen2k/openstc/internal/pkg/libs/const"
	"github.com/aaronchen2k/openstc/internal/server/model"
	"github.com/aaronchen2k/openstc/internal/server/repo"
	"github.com/mitchellh/mapstructure"
)

type AppiumService struct {
	QueueService *QueueService `inject:""`
	RpcService   *RpcService   `inject:""`

	DeviceRepo *repo.DeviceRepo `inject:""`
	BuildRepo  *repo.BuildRepo  `inject:""`
}

func NewAppiumService() *AppiumService {
	return &AppiumService{}
}

func (s *AppiumService) Start(queue model.Queue) (result _domain.RpcResult) {
	serial := queue.Serial
	device := s.DeviceRepo.GetBySerial(serial)

	build := model.NewBuildDetail(queue.ID, uint(0), queue.BuildType,
		serial, queue.Priority, device.NodeIp, device.NodePort)
	s.BuildRepo.Save(&build)

	build = s.BuildRepo.GetBuild(build.ID)
	build.AppiumPort = device.AppiumPort

	rpcResult := s.RpcService.AppiumTest(build)
	if rpcResult.IsSuccess() {
		s.BuildRepo.Start(build)
	} else {
		s.BuildRepo.Delete(build)
	}

	result.Success("")
	return
}

func (s *AppiumService) SaveResult(buildResult _domain.RpcResult, resultPath string) {
	appiumTestTo := _domain.BuildTo{}
	mp := buildResult.Payload.(map[string]interface{})
	mapstructure.Decode(mp, &appiumTestTo)

	progress := _const.ProgressCompleted
	var status _const.BuildStatus
	if buildResult.IsSuccess() {
		status = _const.StatusPass
	} else {
		status = _const.StatusFail
	}

	s.BuildRepo.SaveResult(appiumTestTo, resultPath, progress, status, buildResult.Msg)
	s.QueueService.SetQueueResult(appiumTestTo.QueueId, progress, status)
	if progress == _const.ProgressTimeout {
		s.BuildRepo.SetTimeoutByQueueId(appiumTestTo.QueueId)
	}
}
