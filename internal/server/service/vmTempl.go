package service

import (
	"github.com/aaronchen2k/tester/internal/server/repo"
)

type VmTemplService struct {
	VmTemplRepo *repo.VmTemplRepo `inject:""`
}

func NewVmTemplService() *VmTemplService {
	return &VmTemplService{}
}
