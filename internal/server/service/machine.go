package service

import (
	"github.com/aaronchen2k/tester/internal/pkg/const"
	serverConf "github.com/aaronchen2k/tester/internal/server/cfg"
	"github.com/aaronchen2k/tester/internal/server/domain"
	serviceImpl "github.com/aaronchen2k/tester/internal/server/service/impl"
	serviceInterface "github.com/aaronchen2k/tester/internal/server/service/interface"
	serverConst "github.com/aaronchen2k/tester/internal/server/utils/const"
	"strconv"
)

type MachineService struct {
	ClusterService *ClusterService `inject:""`

	VmService        serviceInterface.VmInterface
	ContainerService serviceInterface.ContainerInterface
}

func NewMachineService() *MachineService {
	inst := &MachineService{}

	if serverConf.Config.Adapter.VmPlatform == serverConst.Pve {
		inst.VmService = serviceImpl.NewPveService()
	}

	if serverConf.Config.Adapter.ContainerPlatform == serverConst.Portainer {
		inst.ContainerService = serviceImpl.NewPortainerService()
	}

	return inst
}

func (s *MachineService) ListVm() (rootNode *domain.ResNode) {
	rootNode = &domain.ResNode{Name: "虚拟机", Type: _const.ResRoot, Id: "0"}
	hosts := s.ClusterService.ListByType("pve")

	for _, host := range hosts {
		id := strconv.Itoa(int(host.ID))

		hostNode := &domain.ResNode{
			Name: host.Name + "(集群)", Type: _const.ResCluster,
			Id: id, Key: string(_const.ResCluster) + "-" + id,
			Ip: host.Ip, Port: host.Port,
			Username: host.Username, Password: host.Password}

		rootNode.Children = append(rootNode.Children, hostNode)

		s.VmService.GetNodeTree(hostNode)
	}

	return
}

func (s *MachineService) ListContainers() (rootNode *domain.ResNode) {
	rootNode = &domain.ResNode{Name: "容器", Type: _const.ResRoot, Id: "0"}
	hosts := s.ClusterService.ListByType("portainer")

	for _, host := range hosts {
		id := strconv.Itoa(int(host.ID))

		hostNode := &domain.ResNode{Name: host.Name + "(集群)", Type: _const.ResCluster,
			Id: id, Key: string(_const.ResCluster) + "-" + id,
			Ip: host.Ip, Port: host.Port,
			Username: host.Username, Password: host.Password}
		rootNode.Children = append(rootNode.Children, hostNode)

		s.ContainerService.GetNodeTree(hostNode)
	}

	return
}
