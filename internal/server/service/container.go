package service

import (
	"fmt"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"github.com/aaronchen2k/tester/internal/server/repo"
)

type ContainerService struct {
	RpcService     *RpcService     `inject:""`
	MachineService *VirtualService `inject:""`

	ContainerRepo      *repo.ContainerRepo      `inject:""`
	ContainerImageRepo *repo.ContainerImageRepo `inject:""`
	ClusterRepo        *repo.ClusterRepo        `inject:""`
	NodeRepo           *repo.NodeRepo           `inject:""`
}

func NewContainerService() *ContainerService {
	return &ContainerService{}
}

func (s *ContainerService) Create(containerImageId uint) (dockerId string, err error) {
	imagePo := s.ContainerImageRepo.Get(containerImageId)
	node := s.NodeRepo.Get(imagePo.Node)
	cluster := s.ClusterRepo.Get(imagePo.Cluster)

	vm, err := s.MachineService.CreateContainer(imagePo, node, cluster)
	// TODO: save to db?
	_logUtils.Info(fmt.Sprintf("%#v, %s", vm, err.Error()))
	return
}

func (s *ContainerService) Register(container _domain.Container) (result _domain.RpcResult) {
	err := s.ContainerRepo.Register(container)
	if err != nil {
		result.Fail(fmt.Sprintf("fail to register host %s ", container.MacAddress))
	}
	return
}
