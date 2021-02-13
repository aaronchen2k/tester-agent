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

func (s BuildService) SaveResult(result _domain.RpcResult, path string) (buildTo _domain.BuildTo) {
	mp := result.Payload.(map[string]interface{})
	mapstructure.Decode(mp, &buildTo)

	progress := _const.ProgressCompleted
	var status _const.BuildStatus
	if result.IsSuccess() {
		status = _const.StatusPass
	} else {
		status = _const.StatusFail
	}

	s.BuildRepo.SaveResult(buildTo, path, progress, status, result.Msg)
	s.QueueService.SetQueueResult(buildTo.QueueId, progress, status)

	return buildTo
}

func NewBuildService() *BuildService {
	return &BuildService{}
}
