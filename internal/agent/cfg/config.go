package agentConf

import (
	"github.com/aaronchen2k/tester/internal/agent/model"
	"github.com/aaronchen2k/tester/internal/agent/utils/common"
	_commonUtils "github.com/aaronchen2k/tester/internal/pkg/libs/common"
	_const "github.com/aaronchen2k/tester/internal/pkg/libs/const"
	_fileUtils "github.com/aaronchen2k/tester/internal/pkg/libs/file"
	_httpUtils "github.com/aaronchen2k/tester/internal/pkg/libs/http"
)

var (
	Inst = model.Config{}
)

func Init() {
	if agentUntils.IsVmAgent() {
		ip, mac := _commonUtils.GetIp()
		Inst.MacAddress = mac.String()
		Inst.Ip = ip.String()
		Inst.Port = _const.RpcPort
	}

	Inst.KvmDir = _fileUtils.UpdateDir(Inst.KvmDir)
	Inst.WorkDir = _fileUtils.UpdateDir(Inst.WorkDir)
	Inst.FarmServer = _httpUtils.UpdateUrl(Inst.FarmServer)
}
