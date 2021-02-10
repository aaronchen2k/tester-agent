package service

import (
	"github.com/aaronchen2k/tester/internal/pkg/const"
	serverConf "github.com/aaronchen2k/tester/internal/server/cfg"
	"github.com/aaronchen2k/tester/internal/server/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/repo"
	serviceInterface "github.com/aaronchen2k/tester/internal/server/service/interface"
	serverConst "github.com/aaronchen2k/tester/internal/server/utils/const"
	"strconv"
)

type ResService struct {
	ClusterService *ClusterService `inject:""`

	VmPlatform        serviceInterface.VmPlatformInterface
	ContainerPlatform serviceInterface.ContainerPlatformInterface

	VmTemplRepo        repo.VmTemplRepo
	ContainerImageRepo repo.ContainerImageRepo
}

func NewResService() *ResService {
	inst := &ResService{}

	if serverConf.Config.Adapter.VmPlatform == serverConst.Pve {
		inst.VmPlatform = NewPveService()
	}

	if serverConf.Config.Adapter.ContainerPlatform == serverConst.Portainer {
		inst.ContainerPlatform = NewPortainerService()
	}

	return inst
}

func (s *ResService) GetValidForQueue(queue model.Queue) (hostId, templOrImageId uint) {
	buildType := queue.BuildType

	if buildType == _const.SeleniumTest {
		//templ := s.VmTemplRepo.Get(queue.VmTemplId)
		//node := templ.Node

	} else if buildType == _const.AppiumTest {

	} else if buildType == _const.UnitTest {

	}

	return
}

//func (s *ResService) getIdleHost() (ids []uint) {
//	// keys: hostId, vmCount
//	hostToVmCountList := s.HostRepo.QueryIdle(_const.MaxVmOnHost)
//
//	hostIds := make([]uint, 0)
//	for _, mp := range hostToVmCountList {
//		hostId := mp["hostId"]
//		hostIds = append(hostIds, hostId)
//	}
//
//	return hostIds
//}

func (s *ResService) ListVm() (rootNode *domain.ResItem) {
	rootNode = &domain.ResItem{Name: "虚拟机", Type: _const.ResRoot, Ident: "0"}
	clusters := s.ClusterService.ListByType("pve")

	for _, cluster := range clusters {
		ident := strconv.Itoa(int(cluster.ID))

		clusterItem := &domain.ResItem{
			Name: cluster.Name + "(集群)", Type: _const.ResCluster,
			Ident: ident, Key: string(_const.ResCluster) + "-" + ident,
			Ip: cluster.Ip, Port: cluster.Port,
			Username: cluster.Username, Password: cluster.Password}

		rootNode.Children = append(rootNode.Children, clusterItem)

		s.VmPlatform.GetNodeTree(clusterItem)
	}

	return
}

func (s *ResService) ListContainers() (rootNode *domain.ResItem) {
	rootNode = &domain.ResItem{Name: "容器", Type: _const.ResRoot, Ident: "0"}
	clusters := s.ClusterService.ListByType("portainer")

	for _, cluster := range clusters {
		id := strconv.Itoa(int(cluster.ID))

		hostNode := &domain.ResItem{Name: cluster.Name + "(集群)", Type: _const.ResCluster,
			Ident: id, Key: string(_const.ResCluster) + "-" + id,
			Ip: cluster.Ip, Port: cluster.Port,
			Username: cluster.Username, Password: cluster.Password}
		rootNode.Children = append(rootNode.Children, hostNode)

		s.ContainerPlatform.GetNodeTree(hostNode)
	}

	return
}

func (s *ResService) CreateVm(name string, templ model.VmTempl, node model.Node, cluster model.Cluster) (
	vm model.Vm, err error) {

	vm, err = s.VmPlatform.CreateVm(name, templ, node, cluster)

	return
}

func (s *ResService) CreateContainer(image model.ContainerImage, node model.Node, cluster model.Cluster) (
	container model.Container, err error) {

	container, err = s.ContainerPlatform.CreateContainer(image, node, cluster)

	return
}
