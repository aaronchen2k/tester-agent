package serviceInterface

import "github.com/aaronchen2k/tester/internal/server/domain"

type MachineInterface interface {
	GetNodeTree(clusterNode *domain.ResNode) (domain.ResNode, error)
}
