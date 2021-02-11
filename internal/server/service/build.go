package service

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"github.com/aaronchen2k/tester/internal/server/repo"
	"github.com/mitchellh/mapstructure"
)

type BuildService struct {
	QueueService *QueueService   `inject:""`
	BuildRepo    *repo.BuildRepo `inject:""`
}

func (s BuildService) SaveResult(result _domain.RpcResult, path string) {
	appiumTestTo := _domain.BuildTo{}
	mp := result.Payload.(map[string]interface{})
	mapstructure.Decode(mp, &appiumTestTo)

	progress := _const.ProgressCompleted
	var status _const.BuildStatus
	if result.IsSuccess() {
		status = _const.StatusPass
	} else {
		status = _const.StatusFail
	}

	s.BuildRepo.SaveResult(appiumTestTo, path, progress, status, result.Msg)
	s.QueueService.SetQueueResult(appiumTestTo.QueueId, progress, status)
	if progress == _const.ProgressTimeout {
		s.BuildRepo.SetTimeoutByQueueId(appiumTestTo.QueueId)
	}
}

func NewBuildService() *BuildService {
	return &BuildService{}
}
