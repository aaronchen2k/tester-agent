package service

import (
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
)

type ProxmoxService struct {
}

func NewProxmoxService() *HostService {
	return &HostService{}
}

func (s *ProxmoxService) Register(host _domain.Host) (result _domain.RpcResult) {

	return
}
