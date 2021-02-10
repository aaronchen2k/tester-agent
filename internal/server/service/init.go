package service

import (
	"fmt"
	middlewareUtils "github.com/aaronchen2k/tester/internal/server/biz/middleware/misc"
	"github.com/aaronchen2k/tester/internal/server/db"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/fatih/color"
)

type InitService struct {
}

func NewInitService() *VmService {
	return &VmService{}
}

func (s *InitService) Init() {
	err := db.GetInst().DB().AutoMigrate(
		model.Models...,
	)
	if err != nil {
		color.Yellow(fmt.Sprintf("初始化数据表错误 ：%+v", err))
	}

	err = db.GetInst().DB().AutoMigrate(
		&middlewareUtils.CasbinRule{},
	)
}
