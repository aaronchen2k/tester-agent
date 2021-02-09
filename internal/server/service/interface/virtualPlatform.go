package serviceInterface

import "github.com/aaronchen2k/tester/internal/server/domain"

type VirtualPlatformInterface interface {
	GetNodeTree(clusterNode *domain.ResItem) error
}
