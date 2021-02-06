package service

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"github.com/aaronchen2k/tester/internal/server/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
	serverUtils "github.com/aaronchen2k/tester/internal/server/utils/common"
	go_portainer "github.com/aaronchen2k/tester/vendors/github.com/leidruid/go-portainer"
	"strconv"
	"strings"
)

type PortainerService struct {
}

func NewPortainerService() *PortainerService {
	return &PortainerService{}
}

func (s *PortainerService) ListContainer(clusterNode *domain.ResNode) (containers []*model.Container, err error) {
	s.GetNodeTree(clusterNode)

	return
}

func (s *PortainerService) CreateContainer(image model.ContainerImage, node model.Node, cluster model.Cluster) (
	container model.Container, err error) {
	config := go_portainer.Config{
		Host:     cluster.Ip,
		Port:     cluster.Port,
		User:     cluster.Username,
		Password: cluster.Password,
		Schema:   "http",
		URL:      "/api",
	}
	portainer := go_portainer.NewPortainer(&config)
	err = portainer.Auth()
	if err != nil {
		_logUtils.Print("fail to connect portainer, error: " + err.Error())
		return
	}

	endpoint, _ := strconv.Atoi(node.Tag)
	vmHostName := serverUtils.GenVmHostName(image.OsPlatform, image.OsName, image.OsLang)

	body := map[string]interface{}{}
	body["Hostname"] = vmHostName
	body["Image"] = image.Name

	dockerId, err := portainer.CreateContainer(uint(endpoint), body)
	if err != nil {
		_logUtils.Printf("fail to create container, error: %d-%s", dockerId, err.Error())
		return
	}

	_, err = portainer.StartContainer(uint(endpoint), dockerId)
	if err != nil {
		_logUtils.Printf("fail to start container, error: %d-%s", dockerId, err.Error())
		return
	}

	container.Ident = dockerId
	container.NodeId = node.ID
	container.ClusterId = cluster.ID

	return
}

func (s *PortainerService) GetNodeTree(clusterNode *domain.ResNode) (root domain.ResNode, err error) {
	config := go_portainer.Config{
		Host:     clusterNode.Ip,
		Port:     clusterNode.Port,
		User:     clusterNode.Username,
		Password: clusterNode.Password,
		Schema:   "http",
		URL:      "/api",
	}
	portainer := go_portainer.NewPortainer(&config)
	err = portainer.Auth()
	if err != nil {
		_logUtils.Print("fail to connect portainer, error: " + err.Error())
		return
	}

	endpoints, _ := portainer.ListEndpoints()
	for _, endpoint := range endpoints {
		id := strconv.Itoa(int(endpoint.Id))

		nodeNode := &domain.ResNode{Name: endpoint.Name + "(节点)", Type: _const.ResNode,
			Id: id, HostId: clusterNode.Id, Key: string(_const.ResNode) + "-" + id}
		clusterNode.Children = append(clusterNode.Children, nodeNode)

		containerFolderNode := &domain.ResNode{Name: "实例", Type: _const.ResFolder,
			Id: id + "-folder-vms", Key: id + "-folder-container"}
		nodeNode.Children = append(nodeNode.Children, containerFolderNode)

		imageFolderNode := &domain.ResNode{Name: "镜像", Type: _const.ResFolder,
			Id: id + "-folder-templs", Key: id + "-folder-image"}
		nodeNode.Children = append(nodeNode.Children, imageFolderNode)

		containers, _ := portainer.ListContainers(endpoint.Id)
		for _, container := range containers {
			containerId := container.ID
			name := s.getContainerName(strings.Join(container.Names, "/"))

			vmNode := &domain.ResNode{Name: name, Type: _const.ResContainer, IsTemplate: false,
				Id: container.ID, HostId: clusterNode.Id, NodeId: nodeNode.Id,
				Key: string(_const.ResContainer) + "-" + containerId}
			containerFolderNode.Children = append(containerFolderNode.Children, vmNode)
		}

		images, _ := portainer.ListImages(endpoint.Id)
		for _, image := range images {
			containerId := image.ID

			path := ""
			if len(image.RepoTags) > 0 {
				path = strings.Join(image.RepoTags, "/")
			} else if len(image.RepoDigests) > 0 {
				path = strings.Join(image.RepoDigests, "/")
			}
			name := s.getImageName(path)

			vmNode := &domain.ResNode{Name: name, Path: path, Type: _const.ResImage, IsTemplate: false,
				Id: image.ID, HostId: clusterNode.Id, NodeId: nodeNode.Id,
				Key: string(_const.ResContainer) + "-" + containerId}
			imageFolderNode.Children = append(imageFolderNode.Children, vmNode)
		}
	}

	return
}

func (s *PortainerService) getContainerName(path string) string {
	if string(path[0]) == "/" {
		return path[1:]
	}
	return path
}

func (s *PortainerService) getImageName(path string) string {
	arr := strings.Split(path, "/")
	if len(arr) <= 2 {
		return path
	}

	name := strings.Join(arr[len(arr)-2:], "/")
	return name
}
