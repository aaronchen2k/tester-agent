package deviceService

import (
	agentConf "github.com/aaronchen2k/tester/internal/agent/conf"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_httpUtils "github.com/aaronchen2k/tester/internal/pkg/libs/http"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
)

func Register(devices []_domain.DeviceInst, isBusy bool) {
	var deviceStatus _const.DeviceStatus
	var serviceStatus _const.ServiceStatus
	if isBusy {
		deviceStatus = _const.DeviceBusy
		serviceStatus = _const.ServiceBusy
	} else {
		deviceStatus = _const.DeviceActive
		serviceStatus = _const.ServiceActive
	}
	for i, _ := range devices {
		devices[i].DeviceStatus = deviceStatus
		devices[i].AppiumStatus = serviceStatus
	}

	url := _httpUtils.GenUrl(agentConf.Inst.Server, "device/register")

	resp, ok := _httpUtils.Post(url, devices)

	msg := ""
	str := "%s to register devices, response is %#v"
	if ok {
		msg = "success"
		_logUtils.Infof(str, msg, resp)
	} else {
		msg = "fail"
		_logUtils.Errorf(str, msg, resp)
	}
}
