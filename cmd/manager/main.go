package main

import (
	"github.com/aaronchen2k/tester/internal/manager/program"
	configUtils "github.com/easysoft/zmanager/pkg/config"
	commonUtils "github.com/easysoft/zmanager/pkg/utils/common"
	constant "github.com/easysoft/zmanager/pkg/utils/const"
	i118Utils "github.com/easysoft/zmanager/pkg/utils/i118"
	"log"
	"os"

	"github.com/kardianos/service"
)

func main() {
	configUtils.Init()

	action := ""
	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	if action != "" && commonUtils.StrInArr(action, constant.Actions) {
		log.Println(i118Utils.I118Prt.Sprintf("valid_actions", service.ControlAction))
	}

	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"
	config := &service.Config{
		Name:        constant.AppName,
		DisplayName: constant.AppName,
		Description: constant.AppName + " service.",
		Dependencies: []string{
			"Requires=network.target",
			"After=network-online.target syslog.target"},
		Option: options,
	}

	prg := &program.Program{}
	s, err := service.New(prg, config)
	if err != nil {
		log.Fatal(err)
	}
	errs := make(chan error, 5)
	program.Logger, err = s.Logger(errs)
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

	if action != "" {
		err := service.Control(s, action)
		if err != nil {
			log.Println(i118Utils.I118Prt.Sprintf("valid_actions", service.ControlAction))
			log.Fatal(err)
		}
		return
	}

	err = s.Run()
	if err != nil {
		program.Logger.Error(err)
	}
}
