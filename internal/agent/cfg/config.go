package agentConf

import (
	"github.com/aaronchen2k/tester/internal/agent/model"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_commonUtils "github.com/aaronchen2k/tester/internal/pkg/libs/common"
	_fileUtils "github.com/aaronchen2k/tester/internal/pkg/libs/file"
	_httpUtils "github.com/aaronchen2k/tester/internal/pkg/libs/http"
)

var (
	Inst = model.Config{}
)

func Init() {
	if IsVmAgent() {
		ip, mac := _commonUtils.GetIp()
		Inst.MacAddress = mac.String()
		Inst.Ip = ip.String()
		Inst.Port = _const.RpcPort
	}

	Inst.KvmDir = _fileUtils.UpdateDir(Inst.KvmDir)
	Inst.WorkDir = _fileUtils.UpdateDir(Inst.WorkDir)
	Inst.FarmServer = _httpUtils.UpdateUrl(Inst.FarmServer)
}

func IsDeviceAgent() bool {
	return IsIosAgent() || IsAndroidAgent()
}

func IsAndroidAgent() bool {
	return Inst.Platform == _const.Android
}

func IsIosAgent() bool {
	return Inst.Platform == _const.Ios
}

func IsHostAgent() bool {
	return Inst.Platform == _const.Host
}
func IsVmAgent() bool {
	return Inst.Platform == _const.Vm
}
