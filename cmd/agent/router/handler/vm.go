package handler

import (
	"encoding/json"
	vmService "github.com/aaronchen2k/tester/internal/agent/service/vm"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"golang.org/x/net/context"
)

type VmAction struct{}

func (t *VmAction) Create(ctx context.Context, vm _domain.Vm, reply *_domain.RpcResult) (err error) {
	err = vmService.Create(&vm)

	if err != nil {
		reply.Fail(err.Error())
		return
	}

	jsonStr, _ := json.Marshal(vm)
	reply.Payload = string(jsonStr)
	reply.Success("")
	return err
}

func (t *VmAction) Remove(ctx context.Context, req _domain.PveReq, reply *_domain.RpcResult) (err error) {
	vmUniqueName := req.VmUniqueName
	err = vmService.Define(vmUniqueName)
	reply.Success("")
	return err
}
