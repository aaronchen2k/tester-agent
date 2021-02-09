package handler

import (
	v1 "github.com/aaronchen2k/tester/cmd/server/router/v1"
	_utils "github.com/aaronchen2k/tester/internal/pkg/utils"
	"github.com/aaronchen2k/tester/internal/server/domain"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
)

type VmTemplCtrl struct {
	Ctx            iris.Context
	VmTemplService *service.VmTemplService `inject:""`
}

func NewVmTemplCtrl() *VmTemplCtrl {
	return &VmTemplCtrl{}
}

func (c *VmTemplCtrl) Load(ctx iris.Context) {
	item := domain.ResItem{}
	err := ctx.ReadJSON(&item)
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	templ := c.VmTemplService.GetByIdent(item.Ident)

	if templ.ID == 0 {
		templ = c.VmTemplService.CreateByNode(item)
	}

	_, _ = ctx.JSON(_utils.ApiRes(200, "", templ))

	return
}

func (c *VmTemplCtrl) Update(ctx iris.Context) {
	data := v1.VmData{}
	err := ctx.ReadJSON(&data)
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	c.VmTemplService.Update(data)

	return
}
