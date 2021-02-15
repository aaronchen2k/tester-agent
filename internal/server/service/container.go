package service

import (
	"fmt"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/repo"
)

type ContainerService struct {
	RpcService     *RpcService `inject:""`
	MachineService *ResService `inject:""`

	ContainerRepo      *repo.ContainerRepo      `inject:""`
	ContainerImageRepo *repo.ContainerImageRepo `inject:""`
	ClusterRepo        *repo.ClusterRepo        `inject:""`
	ComputerRepo       *repo.ComputerRepo       `inject:""`
}

func NewContainerService() *ContainerService {
	return &ContainerService{}
}

func (s *ContainerService) CreateByQueue(queue model.Queue) (dockerId string, err error) {
	//imagePo := s.ContainerImageRepo.GetByIdent(queue.ContainerImageId)
	//computer := s.ComputerRepo.GetByIndent(imagePo.Computer)
	//cluster := s.ClusterRepo.GetByIdent(imagePo.Cluster)

	return
}

func (s *ContainerService) Register(container _domain.Container) (result _domain.RpcResult) {
	err := s.ContainerRepo.Register(container)
	if err != nil {
		result.Fail(fmt.Sprintf("fail to register host %s ", container.MacAddress))
	}
	return
}
