package program

import (
	manageService "github.com/aaronchen2k/tester/cmd/manager/service"
	managerConf "github.com/aaronchen2k/tester/cmd/manager/utils/conf"
	_i118Utils "github.com/aaronchen2k/tester/internal/pkg/libs/i118"
	"github.com/kardianos/service"
	"time"
)

type Program struct {
	exit chan struct{}
}

var Logger service.Logger

func (p *Program) Start(s service.Service) error {
	if service.Interactive() {
		Logger.Info(_i118Utils.I118Prt.Sprintf("launch_in_terminal"))
	} else {
		Logger.Info(_i118Utils.I118Prt.Sprintf("launch_in_service"))
	}
	p.exit = make(chan struct{})

	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *Program) run() error {
	Logger.Warningf(_i118Utils.I118Prt.Sprintf("running", service.Platform()))

	ticker := time.NewTicker(time.Duration(managerConf.Inst.Interval) * time.Second)
	for {
		select {
		case tm := <-ticker.C:
			_ = tm
			Logger.Warningf(_i118Utils.I118Prt.Sprintf("start_to_run"))

			for _, app := range managerConf.Inst.Clients {
				Logger.Infof(_i118Utils.I118Prt.Sprintf("start_to_check", app.Name))

				manageService.CheckUpgrade(app)
				manageService.CheckStatus(app)
			}

		case <-p.exit:
			ticker.Stop()
			return nil
		}
	}
}

func (p *Program) Stop(s service.Service) error {
	// Any work in Stop should be quick, usually a few seconds at most.
	Logger.Info(_i118Utils.I118Prt.Sprintf("stopping"))
	close(p.exit)
	return nil
}
