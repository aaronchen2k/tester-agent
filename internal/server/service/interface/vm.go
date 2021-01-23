package serviceInterface

import "github.com/aaronchen2k/tester/internal/server/domain"

type VmInterface interface {
	List(hostNode *domain.ResNode) (domain.ResNode, error)
}
