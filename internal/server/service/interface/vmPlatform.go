package serviceInterface

import (
	"github.com/aaronchen2k/tester/internal/server/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
)

type VmPlatformInterface interface {
	VirtualPlatformInterface

	ListVm(clusterNode *domain.ResItem) ([]*model.Vm, error)
	CreateVm(name string, templ model.VmTempl, node model.Node, cluster model.Cluster) (model.Vm, error)
}
