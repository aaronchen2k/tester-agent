package handler

import (
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"golang.org/x/net/context"
)

type ImageAction struct{}

func (t *ImageAction) Create(ctx context.Context, vm _domain.Vm, reply *_domain.RpcResult) (err error) {
	return
}

func (t *ImageAction) Remove(ctx context.Context, req _domain.PveReq, reply *_domain.RpcResult) (err error) {
	return
}
