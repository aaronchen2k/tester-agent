package vmService

import (
	"github.com/aaronchen2k/openstc/internal/agent/cfg"
	_domain "github.com/aaronchen2k/openstc/internal/pkg/domain"
	_const "github.com/aaronchen2k/openstc/internal/pkg/libs/const"
	_httpUtils "github.com/aaronchen2k/openstc/internal/pkg/libs/http"
	_logUtils "github.com/aaronchen2k/openstc/internal/pkg/libs/log"
	_shellUtils "github.com/aaronchen2k/openstc/internal/pkg/libs/shell"
	"strings"
)

func RegisterHost() {
	host := _domain.Host{Ip: agentConf.Inst.Ip, Port: agentConf.Inst.Port, Status: _const.HostActive}
	host.Vms = getVms()
	updateVms(host.Vms)

	url := _httpUtils.GenUrl(agentConf.Inst.FarmServer, "host/register")
	resp, ok := _httpUtils.Post(url, host)

	msg := ""
	str := "%s to register host, response is %#v"
	if ok {
		msg = "success"
		_logUtils.Infof(str, msg, resp)
	} else {
		msg = "fail"
		_logUtils.Errorf(str, msg, resp)
	}
}

func RegisterVm(isBusy bool) {
	vm := _domain.Vm{MacAddress: agentConf.Inst.MacAddress, WorkDir: agentConf.Inst.WorkDir,
		PublicIp: agentConf.Inst.Ip, PublicPort: agentConf.Inst.Port}

	if isBusy {
		vm.Status = _const.VmBusy
	} else {
		vm.Status = _const.VmActive
	}

	url := _httpUtils.GenUrl(agentConf.Inst.FarmServer, "vm/register")
	resp, ok := _httpUtils.Post(url, vm)

	msg := ""
	str := "%s to register vm, response is %#v"
	if ok {
		msg = "success"
		_logUtils.Infof(str, msg, resp)
	} else {
		msg = "fail"
		_logUtils.Errorf(str, msg, resp)
	}
}

func getVms() (vms []_domain.Vm) {
	cmd := "virsh list --all"
	out, _ := _shellUtils.ExeShell(cmd)

	lines := strings.Split(out, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.Index(line, "Id") == 0 || strings.Index(line, "---") == 0 {
			continue
		}

		cols := strings.Split(line, " ")
		name := strings.TrimSpace(cols[1])
		status := strings.TrimSpace(cols[2])

		if len(name) < 32 { // not created by farm
			continue
		}

		vm := _domain.Vm{}
		vm.Name = name

		vm.Status = _const.VmUnknown
		if status == "running" {
			vm.Status = _const.VmRunning
		} else if status == "shut off" {
			vm.Status = _const.VmDestroy
		}

		vms = append(vms, vm)
	}

	return vms
}
