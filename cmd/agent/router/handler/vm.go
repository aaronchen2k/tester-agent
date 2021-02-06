package handler

import (
	vmService "github.com/aaronchen2k/tester/internal/agent/service/vm"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"golang.org/x/net/context"
)

type VmAction struct{}

func (c *VmAction) Destroy(ctx context.Context, req _domain.PveReq, reply *_domain.RpcResult) (err error) {
	vmUniqueName := req.VmUniqueName
	err = vmService.Define(vmUniqueName)
	reply.Success("")
	return err
}
