package serviceInterface

import (
	"github.com/aaronchen2k/tester/internal/server/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
)

type ContainerInterface interface {
	MachineInterface

	ListContainer(clusterNode *domain.ResItem) ([]*model.Container, error)
	CreateContainer(image model.ContainerImage, node model.Node, cluster model.Cluster) (model.Container, error)
}
