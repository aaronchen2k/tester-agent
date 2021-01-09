package agentUntils

import (
	"github.com/aaronchen2k/tester/internal/agent/cfg"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
)

func IsDeviceAgent() bool {
	return IsIosAgent() || IsAndroidAgent()
}

func IsAndroidAgent() bool {
	return agentConf.Inst.Platform == _const.Android
}

func IsIosAgent() bool {
	return agentConf.Inst.Platform == _const.Ios
}

func IsHostAgent() bool {
	return agentConf.Inst.Platform == _const.Host
}
func IsVmAgent() bool {
	return agentConf.Inst.Platform == _const.Vm
}
