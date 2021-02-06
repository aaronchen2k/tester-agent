package handler

import (
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"golang.org/x/net/context"
)

type ContainerImageAction struct{}

func (c *ContainerImageAction) Test(ctx context.Context, req _domain.PveReq, reply *_domain.RpcResult) (err error) {
	reply.Success("")
	return err
}
