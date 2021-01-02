package handler

import (
	"fmt"
	testService "github.com/aaronchen2k/openstc/internal/agent/service/test"
	_domain "github.com/aaronchen2k/openstc/internal/pkg/domain"
	"golang.org/x/net/context"
)

type AppiumAction struct{}

func (t *AppiumAction) AppiumTest(ctx context.Context, build _domain.BuildTo, reply *_domain.RpcResult) error {
	size := testService.GetTaskSize()
	if size == 0 {
		testService.AddTask(build)
		reply.Success("Success to add job.")
	} else {
		reply.Fail(fmt.Sprintf("already has %d jobs to be done.", size))
	}

	return nil
}
