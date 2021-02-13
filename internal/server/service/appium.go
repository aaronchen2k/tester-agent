package service

import (
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/repo"
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

func (s *AppiumService) Run(queue model.Queue) (result _domain.RpcResult) {
	serial := queue.Serial
	device := s.DeviceRepo.GetBySerial(serial)

	build := model.NewBuildDetail(queue.ID, 0, device.NodeIp, device.NodePort)
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
