package handler

import (
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"github.com/aaronchen2k/tester/internal/pkg/utils"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
	"github.com/mitchellh/mapstructure"
)

type RpcCtrl struct {
	Ctx        iris.Context
	RpcService *service.RpcService `inject:""`
}

func NewRpcCtrl() *RpcCtrl {
	return &RpcCtrl{}
}

func (c *RpcCtrl) Request(ctx iris.Context) {
	rpcReq := _domain.RpcReq{}
	err := ctx.ReadJSON(&rpcReq)
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	var obj interface{}
	if rpcReq.ApiPath == "vm" {
		var vm _domain.Vm
		err = mapstructure.Decode(rpcReq.Data, &vm)
		obj = interface{}(vm)
	}

	rpcResult := c.RpcService.Request(rpcReq.NodeIp, rpcReq.NodePort, rpcReq.ApiPath, rpcReq.ApiMethod, &obj)

	_, _ = ctx.JSON(_utils.ApiRes(200, "请求成功", rpcResult))
}
