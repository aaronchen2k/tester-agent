package serviceInterface

import "github.com/aaronchen2k/tester/internal/server/domain"

type VirtualPlatformInterface interface {
	GetNodeTree(clusterItem *domain.ResItem) error
}
