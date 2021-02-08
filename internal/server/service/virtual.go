package service

import (
	"github.com/aaronchen2k/tester/internal/pkg/const"
	serverConf "github.com/aaronchen2k/tester/internal/server/cfg"
	"github.com/aaronchen2k/tester/internal/server/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
	serviceInterface "github.com/aaronchen2k/tester/internal/server/service/interface"
	serverConst "github.com/aaronchen2k/tester/internal/server/utils/const"
	"strconv"
)

type VirtualService struct {
	ClusterService *ClusterService `inject:""`

	VmService        serviceInterface.VmInterface
	ContainerService serviceInterface.ContainerInterface
}

func NewVirtualService() *VirtualService {
	inst := &VirtualService{}

	if serverConf.Config.Adapter.VmPlatform == serverConst.Pve {
		inst.VmService = NewPveService()
	}

	if serverConf.Config.Adapter.ContainerPlatform == serverConst.Portainer {
		inst.ContainerService = NewPortainerService()
	}

	return inst
}

func (s *VirtualService) ListVm() (rootNode *domain.ResItem) {
	rootNode = &domain.ResItem{Name: "虚拟机", Type: _const.ResRoot, Ident: "0"}
	hosts := s.ClusterService.ListByType("pve")

	for _, host := range hosts {
		ident := strconv.Itoa(int(host.ID))

		clusterItem := &domain.ResItem{
			Name: host.Name + "(集群)", Type: _const.ResCluster,
			Ident: ident, Key: string(_const.ResCluster) + "-" + ident,
			Ip: host.Ip, Port: host.Port,
			Username: host.Username, Password: host.Password}

		rootNode.Children = append(rootNode.Children, clusterItem)

		s.VmService.GetNodeTree(clusterItem)
	}

	return
}

func (s *VirtualService) ListContainers() (rootNode *domain.ResItem) {
	rootNode = &domain.ResItem{Name: "容器", Type: _const.ResRoot, Ident: "0"}
	hosts := s.ClusterService.ListByType("portainer")

	for _, host := range hosts {
		id := strconv.Itoa(int(host.ID))

		hostNode := &domain.ResItem{Name: host.Name + "(集群)", Type: _const.ResCluster,
			Ident: id, Key: string(_const.ResCluster) + "-" + id,
			Ip: host.Ip, Port: host.Port,
			Username: host.Username, Password: host.Password}
		rootNode.Children = append(rootNode.Children, hostNode)

		s.ContainerService.GetNodeTree(hostNode)
	}

	return
}

func (s *VirtualService) CreateVm(templ model.VmTempl, node model.Node, cluster model.Cluster) (
	vm model.Vm, err error) {

	vm, err = s.VmService.CreateVm(templ, node, cluster)

	return
}

func (s *VirtualService) CreateContainer(image model.ContainerImage, node model.Node, cluster model.Cluster) (
	container model.Container, err error) {

	container, err = s.ContainerService.CreateContainer(image, node, cluster)

	return
}
