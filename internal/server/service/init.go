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
		&model.User{},
		&model.Role{},
		&model.Permission{},
		&middlewareUtils.CasbinRule{},

		&model.Build{},
		&model.Container{},
		&model.Device{},
		&model.Cluster{},
		&model.Image{},
		&model.Iso{},
		&model.Queue{},
		&model.Task{},
		&model.Vm{},
	)

	if err != nil {
		color.Yellow(fmt.Sprintf("初始化数据表错误 ：%+v", err))
	}
}
