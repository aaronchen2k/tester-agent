package handler

import (
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_utils "github.com/aaronchen2k/tester/internal/pkg/utils"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
)

type ContainerCtrl struct {
	Ctx              iris.Context
	ContainerService *service.ContainerService `inject:""`
}

func NewContainerCtrl() *ContainerCtrl {
	return &ContainerCtrl{}
}

func (c *ContainerCtrl) Register(ctx iris.Context) {
	var container _domain.Container
	if err := c.Ctx.ReadJSON(&container); err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), ""))
		return
	}

	c.ContainerService.Register(container)

	_, _ = ctx.JSON(_utils.ApiRes(200, "", ""))
	return
}
