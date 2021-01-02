package deviceService

import (
	"github.com/aaronchen2k/openstc/internal/agent/cfg"
	_domain "github.com/aaronchen2k/openstc/internal/pkg/domain"
	_const "github.com/aaronchen2k/openstc/internal/pkg/libs/const"
	_httpUtils "github.com/aaronchen2k/openstc/internal/pkg/libs/http"
	_logUtils "github.com/aaronchen2k/openstc/internal/pkg/libs/log"
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

	url := _httpUtils.GenUrl(agentConf.Inst.FarmServer, "device/register")

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
