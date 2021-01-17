package program

import (
	manageService "github.com/easysoft/zmanager/pkg/service"
	constant "github.com/easysoft/zmanager/pkg/utils/const"
	logUtils "github.com/easysoft/zmanager/pkg/utils/log"
	"github.com/easysoft/zmanager/pkg/utils/vari"
	"github.com/kardianos/service"
	"log"
	"os"
	"time"
)

type Program struct {
	exit chan struct{}
}

var Logger service.Logger

func (p *Program) Start(s service.Service) error {
	if service.Interactive() {
		Logger.Info("zmanager run in terminal.")
	} else {
		Logger.Info("zmanager run as service.")
	}
	p.exit = make(chan struct{})

	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *Program) run() error {
	file, _ := os.OpenFile(vari.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	logUtils.Init(file)

	Logger.Warningf("I'm running %v.", service.Platform())
	ticker := time.NewTicker(time.Duration(vari.Config.Interval) * time.Second)
	for {
		select {
		case tm := <-ticker.C:
			_ = tm
			Logger.Warningf("start to run.")

			for _, app := range constant.Apps {
				log.Printf("start to check %s.", app)

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
	Logger.Info("I'm Stopping!")
	close(p.exit)
	return nil
}
