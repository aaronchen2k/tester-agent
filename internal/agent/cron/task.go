package cron

import (
	"fmt"
	checkService "github.com/aaronchen2k/tester/internal/agent/service/check"
	_const "github.com/aaronchen2k/tester/internal/pkg/libs/const"
	_cronUtils "github.com/aaronchen2k/tester/internal/pkg/libs/cron"
)

func Init() {
	_cronUtils.AddTaskFuc(
		"check",
		fmt.Sprintf("@every %ds", _const.AgentCheckDevice),
		func() {
			checkService.Check()
		},
	)
}
