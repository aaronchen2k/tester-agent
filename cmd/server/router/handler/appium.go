package handler

import (
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
)

type AppiumCtrl struct {
	Ctx     iris.Context
	Service *service.AppiumService `inject:""`
}

func NewAppiumCtrl() *AppiumCtrl {
	return &AppiumCtrl{}
}
