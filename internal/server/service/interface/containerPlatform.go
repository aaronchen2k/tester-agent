package serviceInterface

import (
	"github.com/aaronchen2k/tester/internal/server/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
)

type ContainerPlatformInterface interface {
	VirtualPlatformInterface

	ListContainer(clusterItem *domain.ResItem) ([]*model.Container, error)
	CreateContainer(queueId uint, image model.ContainerImage, computer model.Computer, cluster model.Cluster) (model.Container, error)
	DestroyContainer(ident string, computer model.Computer, cluster model.Cluster) error
}
