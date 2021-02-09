package serverCron

import (
	"fmt"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_cronUtils "github.com/aaronchen2k/tester/internal/pkg/libs/cron"
	"github.com/aaronchen2k/tester/internal/server/service"
)

type ServerCron struct {
	ExeService *service.ExecService `inject:""`
}

func NewServerCron() *ServerCron {
	inst := &ServerCron{}
	inst.Init()
	return inst
}

func (s *ServerCron) Init() {
	_cronUtils.AddTaskFuc(
		"check",
		fmt.Sprintf("@every %ds", _const.WebCheckQueueInterval),
		func() {
			s.ExeService.CheckAll()
		},
	)
}
