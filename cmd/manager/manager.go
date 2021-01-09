package main

import (
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"log"
	"os"
	"time"

	"github.com/kardianos/service"
)

var (
	logger service.Logger
	action string
	user   string
)

type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	if service.Interactive() {
		logger.Info("run in terminal.")
	} else {
		logger.Info("run as service.")
	}
	p.exit = make(chan struct{})

	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() error {
	_logUtils.Init()

	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case tm := <-ticker.C:
			logger.Warningf("logger aaaaaa at %v...", tm)
			_logUtils.Printf("_logUtils aaaaaa at %v...", tm)
			_logUtils.Errorf("_logUtils aaaaaa at %v...", tm)
		case <-p.exit:
			ticker.Stop()
			return nil
		}
	}
}
func (p *program) Stop(s service.Service) error {
	// Any work in Stop should be quick, usually a few seconds at most.
	logger.Warningf("I'm Stopping!")
	close(p.exit)
	return nil
}

func main() {
	if len(os.Args) >= 2 {
		action = os.Args[1]
	}
	if action == "install" && len(os.Args) == 3 {
		user = os.Args[2]
	}

	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"
	svcConfig := &service.Config{
		Name:        "tester-manager",
		DisplayName: "Go Service Example for Logging",
		Description: "This is an example Go service that outputs log messages.",
		Dependencies: []string{
			"Requires=network.target",
			"After=network-online.target syslog.target"},
		Option: options,
	}
	if user != "" {
		svcConfig.UserName = user
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	if len(action) != 0 {
		err := service.Control(s, action)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
