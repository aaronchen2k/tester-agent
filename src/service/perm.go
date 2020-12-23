package service

import (
	"github.com/aaronchen2k/openstc/src/repo"
)

type PermService struct {
	PermRepo *repo.PermRepo `inject:""`
}

func NewPermService() *PermService {
	return &PermService{}
}
