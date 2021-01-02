package cron

import (
	"fmt"
	checkService "github.com/aaronchen2k/openstc/internal/agent/service/check"
	_const "github.com/aaronchen2k/openstc/internal/pkg/libs/const"
	_cronUtils "github.com/aaronchen2k/openstc/internal/pkg/libs/cron"
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
