package service

import (
	"github.com/aaronchen2k/openstc/src/repo"
)

type IsoService struct {
	IsoRepo *repo.IsoRepo `inject:""`
}

func NewIsoService() *IsoService {
	return &IsoService{}
}
