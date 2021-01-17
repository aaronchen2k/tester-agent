package manageService

import (
	_managerVari "github.com/aaronchen2k/tester/internal/manager/utils/var"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_commonUtils "github.com/aaronchen2k/tester/internal/pkg/libs/common"
	"strings"
)

func CheckStatus(app string) {
	output, _ := GetAgentProcess(app)
	output = strings.TrimSpace(output)

	if output != "" {
		return
	}

	version := _managerVari.Config.AgentVersion
	startApp(app, version)
}

func startApp(app string, version string) (err error) {
	appDir := _managerVari.WorkDir + app + _const.PthSep

	newExePath := appDir + version + _const.PthSep + app + _const.PthSep + app
	if _commonUtils.IsWin() {
		newExePath += ".exe"
	}

	StartAgentProcess(newExePath, app)

	return
}
