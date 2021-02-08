package handler

import (
	"github.com/aaronchen2k/tester/internal/pkg/utils"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
)

type EnvCtrl struct {
	Ctx        iris.Context
	EnvService *service.EnvService `inject:""`
}

func NewEnvCtrl() *EnvCtrl {
	return &EnvCtrl{}
}

func (c *EnvCtrl) List(ctx iris.Context) {
	mp := map[string]interface{}{}
	osPlatforms, osTypes, osLangs, browserTypes := c.EnvService.List()

	mp["osPlatforms"] = osPlatforms
	mp["osTypes"] = osTypes
	mp["osLangs"] = osLangs
	mp["browserTypes"] = browserTypes

	_, _ = ctx.JSON(_utils.ApiRes(200, "", mp))
}
