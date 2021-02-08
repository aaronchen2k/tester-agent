package agentConf

import (
	"github.com/aaronchen2k/tester/internal/agent/agentModel"
	agentConst "github.com/aaronchen2k/tester/internal/agent/utils/const"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_commonUtils "github.com/aaronchen2k/tester/internal/pkg/libs/common"
	_fileUtils "github.com/aaronchen2k/tester/internal/pkg/libs/file"
	_httpUtils "github.com/aaronchen2k/tester/internal/pkg/libs/http"
	_i118Utils "github.com/aaronchen2k/tester/internal/pkg/libs/i118"
)

var (
	Inst = agentModel.Config{}
)

func Init() {
	_i118Utils.InitI118(Inst.Language, agentConst.AppName)

	ip, _, hostName := _commonUtils.GetIp()
	Inst.HostName = hostName
	Inst.Ip = ip.String()
	Inst.Port = _const.RpcPort

	Inst.KvmDir = _fileUtils.UpdateDir(Inst.KvmDir)
	Inst.WorkDir = _fileUtils.UpdateDir(Inst.WorkDir)
	Inst.Server = _httpUtils.UpdateUrl(Inst.Server)
}

func IsVmAgent() bool {
	return Inst.Platform == _const.Vm
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
