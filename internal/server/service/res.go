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
	ClusterService    *ClusterService `inject:""`
	VmPlatform        serviceInterface.VmPlatformInterface
	ContainerPlatform serviceInterface.ContainerPlatformInterface

	ClusterRepo *repo.ClusterRepo `inject:""`
	NodeRepo    *repo.NodeRepo    `inject:""`

	BuildRepo          *repo.BuildRepo          `inject:""`
	VmRepo             *repo.VmRepo             `inject:""`
	ContainerRepo      *repo.ContainerRepo      `inject:""`
	VmTemplRepo        *repo.VmTemplRepo        `inject:""`
	ContainerImageRepo *repo.ContainerImageRepo `inject:""`
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
	vmIdent string, err error) {

	vmIdent, err = s.VmPlatform.CreateVm(name, templ, node, cluster)

	return
}
func (s *ResService) CreateContainer(queueId uint, image model.ContainerImage, node model.Node, cluster model.Cluster) (
	container model.Container, err error) {

	container, err = s.ContainerPlatform.CreateContainer(queueId, image, node, cluster)

	return
}

func (s *ResService) DestroyByBuild(buildId uint) {
	build := s.BuildRepo.GetBuild(buildId)
	if build.BuildType == _const.SeleniumTest {
		vm := s.VmRepo.GetById(build.VmId)
		cluster := s.ClusterRepo.Get(vm.ClusterId)

		s.DestroyVm(vm.Ident, cluster)
	} else if build.BuildType == _const.AppiumTest {
		container := s.VmRepo.GetById(build.ContainerId)
		node := s.NodeRepo.Get(container.NodeId)
		cluster := s.ClusterRepo.Get(container.ClusterId)

		s.DestroyContainer(container.Ident, node, cluster)
	}
}

func (s *ResService) DestroyTimeout() {
	s.DestroyTimeoutVm()
	s.DestroyTimeoutContainer()
}
func (s *ResService) DestroyTimeoutVm() {
	vms := s.VmRepo.QueryForDestroy()
	for _, vm := range vms {
		cluster := s.ClusterRepo.Get(vm.ClusterId)
		s.DestroyVm(vm.Ident, cluster)
	}
}
func (s *ResService) DestroyTimeoutContainer() {
	containers := s.ContainerRepo.QueryForDestroy()
	for _, container := range containers {
		cluster := s.ClusterRepo.Get(container.ClusterId)
		s.DestroyVm(container.Ident, cluster)
	}
}

func (s *ResService) DestroyVm(vmIdent string, cluster model.Cluster) (err error) {
	err = s.VmPlatform.DestroyVm(vmIdent, cluster)

	return
}
func (s *ResService) DestroyContainer(containerIdent string, node model.Node, cluster model.Cluster) (err error) {
	err = s.ContainerPlatform.DestroyContainer(containerIdent, node, cluster)

	return
}
