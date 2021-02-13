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

func (s *PortainerService) ListContainer(clusterNode *domain.ResItem) (containers []*model.Container, err error) {
	s.GetNodeTree(clusterNode)

	return
}

func (s *PortainerService) CreateContainer(queueId uint, image model.ContainerImage, node model.Node, cluster model.Cluster) (
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

	endpoint, _ := strconv.Atoi(node.Ident)
	vmHostName := serverUtils.GenVmHostName(queueId, image.OsPlatform, image.OsType, image.OsLang)

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

func (s *PortainerService) DestroyContainer(ident string, node model.Node, cluster model.Cluster) (err error) {
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

	endpoint, _ := strconv.Atoi(node.Ident)

	dockerId, err := portainer.RemoveContainer(uint(endpoint), ident)
	if err != nil {
		_logUtils.Printf("fail to create container, error: %d-%s", dockerId, err.Error())
		return
	}

	return
}

func (s *PortainerService) GetNodeTree(clusterItem *domain.ResItem) (err error) {
	config := go_portainer.Config{
		Host:     clusterItem.Ip,
		Port:     clusterItem.Port,
		User:     clusterItem.Username,
		Password: clusterItem.Password,
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
		ident := strconv.Itoa(int(endpoint.Id))

		nodeItem := &domain.ResItem{Name: endpoint.Name + "(节点)", Type: _const.ResNode,
			Ident: ident, Cluster: clusterItem.Ident, Key: string(_const.ResNode) + "-" + ident}
		clusterItem.Children = append(clusterItem.Children, nodeItem)

		containerFolderItem := &domain.ResItem{Name: "实例", Type: _const.ResFolder,
			Ident: ident + "-folder-vms", Key: ident + "-folder-container"}
		nodeItem.Children = append(nodeItem.Children, containerFolderItem)

		imageFolderItem := &domain.ResItem{Name: "镜像", Type: _const.ResFolder,
			Ident: ident + "-folder-templs", Key: ident + "-folder-image"}
		nodeItem.Children = append(nodeItem.Children, imageFolderItem)

		containers, _ := portainer.ListContainers(endpoint.Id)
		for _, container := range containers {
			containerId := container.ID
			name := s.getContainerName(strings.Join(container.Names, "/"))

			containerItem := &domain.ResItem{Name: name, Type: _const.ResContainer, IsTemplate: false,
				Ident: container.ID, Node: nodeItem.Ident, Cluster: clusterItem.Ident,
				Key: string(_const.ResContainer) + "-" + containerId}
			containerFolderItem.Children = append(containerFolderItem.Children, containerItem)
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

			imageItem := &domain.ResItem{Name: name, Path: path, Type: _const.ResImage, IsTemplate: false,
				Ident: image.ID, Node: nodeItem.Ident, Cluster: clusterItem.Ident,
				Key: string(_const.ResContainer) + "-" + containerId}
			imageFolderItem.Children = append(imageFolderItem.Children, imageItem)
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
