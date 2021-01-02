package service

import (
	"github.com/aaronchen2k/openstc/internal/server/repo"
)

type IsoService struct {
	IsoRepo *repo.IsoRepo `inject:""`
}

func NewIsoService() *IsoService {
	return &IsoService{}
}
