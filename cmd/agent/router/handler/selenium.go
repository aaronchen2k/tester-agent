package handler

import (
	"fmt"
	testService "github.com/aaronchen2k/tester/internal/agent/service/test"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"golang.org/x/net/context"
)

type SeleniumAction struct{}

func (t *SeleniumAction) SeleniumTest(ctx context.Context, task _domain.BuildTo, reply *_domain.RpcResult) error {
	size := testService.GetTaskSize()
	if size == 0 {
		testService.AddTask(task)
		reply.Success("Success to add job.")
	} else {
		reply.Fail(fmt.Sprintf("already has %d jobs to be done.", size))
	}

	return nil
}
