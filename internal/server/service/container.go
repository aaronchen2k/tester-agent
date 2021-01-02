package service

import (
	"fmt"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"github.com/aaronchen2k/tester/internal/server/repo"
)

type ContainerService struct {
	RpcService *RpcService `inject:""`

	ContainerRepo *repo.ContainerRepo `inject:""`
	HostRepo      *repo.HostRepo      `inject:""`
	ImageRepo     *repo.ImageRepo     `inject:""`
	IsoRepo       *repo.IsoRepo       `inject:""`
	QueueRepo     *repo.QueueRepo     `inject:""`
}

func NewContainerService() *ContainerService {
	return &ContainerService{}
}

func (s *ContainerService) Register(container _domain.Container) (result _domain.RpcResult) {
	err := s.ContainerRepo.Register(container)
	if err != nil {
		result.Fail(fmt.Sprintf("fail to register host %s ", container.MacAddress))
	}
	return
}
