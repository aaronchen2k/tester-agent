package handler

import (
	"fmt"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"github.com/aaronchen2k/tester/internal/pkg/utils"
	"github.com/aaronchen2k/tester/internal/server/domain"
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

func (c *ContainerCtrl) Create(ctx iris.Context) {
	req := domain.VmReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	c.ContainerService.Create(req.TemplId)

	return
}

func (c *ContainerCtrl) Register() (result _domain.RpcResult) {
	var container _domain.Container
	if err := c.Ctx.ReadJSON(&container); err != nil {
		_logUtils.Error(err.Error())
		result.Fail("wrong request data")
		return
	}

	po := c.ContainerService.Register(container)

	result.Success(fmt.Sprintf("succes to register host %d.", po))
	return result
}
