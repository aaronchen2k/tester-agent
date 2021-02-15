package serviceInterface

import (
	"github.com/aaronchen2k/tester/internal/server/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
)

type VmPlatformInterface interface {
	VirtualPlatformInterface

	ListVm(clusterItem *domain.ResItem) ([]*model.Vm, error)
	CreateVm(name string, templ model.VmTempl, computer model.Computer, cluster model.Cluster) (string, error)
	DestroyVm(ident string, cluster model.Cluster) error
}
