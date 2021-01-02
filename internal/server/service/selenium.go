package service

import (
	_domain "github.com/aaronchen2k/openstc/internal/pkg/domain"
	_const "github.com/aaronchen2k/openstc/internal/pkg/libs/const"
	"github.com/aaronchen2k/openstc/internal/server/model"
	"github.com/aaronchen2k/openstc/internal/server/repo"
	"github.com/mitchellh/mapstructure"
)

type SeleniumService struct {
	RpcService   *RpcService   `inject:""`
	QueueService *QueueService `inject:""`

	VmRepo    *repo.VmRepo    `inject:""`
	BuildRepo *repo.BuildRepo `inject:""`
}

func NewSeleniumService() *SeleniumService {
	return &SeleniumService{}
}

func (s *SeleniumService) Start(queue model.Queue) (result _domain.RpcResult) {
	vmId := queue.VmId
	vm := s.VmRepo.GetById(vmId)

	build := model.NewBuildDetail(queue.ID, vmId, queue.BuildType,
		"", queue.Priority, vm.Ip, vm.Port)
	s.BuildRepo.Save(&build)

	build = s.BuildRepo.GetBuild(build.ID)

	rpcResult := s.RpcService.SeleniumTest(build)
	if rpcResult.IsSuccess() {
		s.BuildRepo.Start(build)
	} else {
		s.BuildRepo.Delete(build)
	}

	result.Success("")
	return
}

func (s *SeleniumService) SaveResult(buildResult _domain.RpcResult, resultPath string) {
	seleniumTestTo := _domain.BuildTo{}
	mp := buildResult.Payload.(map[string]interface{})
	mapstructure.Decode(mp, &seleniumTestTo)

	progress := _const.ProgressCompleted
	var result _const.BuildStatus
	if buildResult.IsSuccess() {
		result = _const.StatusPass
	} else {
		result = _const.StatusFail
	}

	s.BuildRepo.SaveResult(seleniumTestTo, resultPath, progress, result, buildResult.Msg)
	s.QueueService.SetQueueResult(seleniumTestTo.QueueId, progress, result)
	if progress == _const.ProgressTimeout {
		s.BuildRepo.SetTimeoutByQueueId(seleniumTestTo.QueueId)
	}
}
