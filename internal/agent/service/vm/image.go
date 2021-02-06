package vmService

import (
	"fmt"
	"github.com/aaronchen2k/tester/internal/agent/cfg"
	agentConst "github.com/aaronchen2k/tester/internal/agent/utils/const"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	_shellUtils "github.com/aaronchen2k/tester/internal/pkg/libs/shell"
)

func RemoveDefImage(vmUniqueName string) (err error) {
	rawPath := fmt.Sprintf("%s%s.qcow2", agentConst.FolderImage, vmUniqueName)

	cmd := "rm -rf " + rawPath
	_, err = _shellUtils.ExeShellInDir(cmd, agentConf.Inst.WorkDir)
	if err != nil {
		_logUtils.Errorf("fail to remove image, cmd %s, err %s.", cmd, err.Error())
	}

	return
}
