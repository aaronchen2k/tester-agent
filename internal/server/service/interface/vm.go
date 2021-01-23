package serviceInterface

import (
	"github.com/aaronchen2k/tester/internal/server/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
)

type VmInterface interface {
	MachineInterface

	ListVm(clusterNode *domain.ResNode) ([]*model.Vm, error)
}
