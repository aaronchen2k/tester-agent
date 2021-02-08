package main

import (
	"flag"
	"github.com/aaronchen2k/tester/cmd/agent/router"
	"github.com/aaronchen2k/tester/internal/agent/cfg"
	"github.com/aaronchen2k/tester/internal/agent/cron"
	agentUntils "github.com/aaronchen2k/tester/internal/agent/utils/common"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"github.com/fatih/color"
	"os"
	"os/signal"
	"syscall"
)

var (
	help     bool
	flagSet  *flag.FlagSet
	platform string
)

func main() {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		cleanup()
		os.Exit(0)
	}()

	flagSet = flag.NewFlagSet("tester", flag.ContinueOnError)

	flagSet.StringVar(&agentConf.Inst.KvmDir, "k", "kvm", "")
	flagSet.StringVar(&agentConf.Inst.WorkDir, "w", "work", "")

	flagSet.StringVar(&agentConf.Inst.Server, "s", "", "")
	flagSet.StringVar(&agentConf.Inst.Ip, "i", "", "")
	flagSet.IntVar(&agentConf.Inst.Port, "p", 10, "")
	flagSet.StringVar(&platform, "t", string(_const.OsAndroid), "")

	flagSet.BoolVar(&help, "h", false, "")

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "-h")
	}

	switch os.Args[1] {
	case "start":
		start(os.Args)

	case "help", "-h":
		agentUntils.PrintUsage()

	default:
		if len(os.Args) > 1 {
			args := []string{os.Args[0], "start"}
			args = append(args, os.Args[1:]...)

			start(args)
		} else {
			agentUntils.PrintUsage()
		}
	}
}

func start(args []string) {
	if err := flagSet.Parse(args[2:]); err == nil {
		agentConf.Inst.Platform = _const.WorkPlatform(platform)

		_logUtils.Init()
		agentConf.Init()
		cron.Init()

		router.App()
	}
}

func init() {
	cleanup()

}

func cleanup() {
	color.Unset()
}
