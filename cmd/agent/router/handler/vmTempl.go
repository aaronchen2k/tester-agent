package handler

import (
	"github.com/aaronchen2k/tester/internal/pkg/domain"
	"golang.org/x/net/context"
)

type VmTemplAction struct{}

func (c *VmTemplAction) Test(ctx context.Context, req _domain.PveReq, reply *_domain.RpcResult) (err error) {
	reply.Success("")
	return err
}
