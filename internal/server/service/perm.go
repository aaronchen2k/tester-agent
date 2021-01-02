package service

import (
	"github.com/aaronchen2k/tester/internal/server/repo"
)

type PermService struct {
	PermRepo *repo.PermRepo `inject:""`
}

func NewPermService() *PermService {
	return &PermService{}
}
