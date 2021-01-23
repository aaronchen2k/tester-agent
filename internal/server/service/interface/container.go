package serviceInterface

import (
	"github.com/aaronchen2k/tester/internal/server/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
)

type ContainerInterface interface {
	MachineInterface

	ListContainer(hostNode *domain.ResNode) ([]*model.Container, error)
}
