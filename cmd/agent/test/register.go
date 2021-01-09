package test

import (
	"fmt"
	"github.com/aaronchen2k/tester/internal/agent/cfg"
	deviceService "github.com/aaronchen2k/tester/internal/agent/service/device"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_httpUtils "github.com/aaronchen2k/tester/internal/pkg/libs/http"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"log"
	"testing"
)

func TestRegister(t *testing.T) {
	_logUtils.Init()
	url := _httpUtils.GenUrl("http://localhost:8848", "device/register")
	log.Println(url)

	agentConf.Inst.Platform = _const.Android
	devices := deviceService.RefreshDevices()
	resp, ok := _httpUtils.Post(url, devices)

	msg := ""
	str := "%s to register devices, response is %#v"
	if ok {
		msg = "success"
		log.Println(fmt.Sprintf(str, msg, resp))
	} else {
		msg = "fail"
		log.Println(fmt.Sprintf(str, msg, resp))
	}
}
