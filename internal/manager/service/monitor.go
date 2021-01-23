package manageService

import (
	commonUtils "github.com/easysoft/zmanager/pkg/utils/common"
	constant "github.com/easysoft/zmanager/pkg/utils/const"
	shellUtils "github.com/easysoft/zmanager/pkg/utils/shell"
	"github.com/easysoft/zmanager/pkg/utils/vari"
	"strings"
)

func CheckStatus(app string) {
	output, _ := shellUtils.GetProcess(app)
	output = strings.TrimSpace(output)

	if output != "" {
		return
	}

	version := ""
	if app == constant.ZTF {
		version = vari.Config.ZTFVersion
	} else if app == constant.ZenData {
		version = vari.Config.ZDVersion
	}

	startApp(app, version)
}

func startApp(app string, version string) (err error) {
	appDir := vari.WorkDir + app + constant.PthSep

	newExePath := appDir + version + constant.PthSep + app + constant.PthSep + app
	if commonUtils.IsWin() {
		newExePath += ".exe"
	}

	shellUtils.StartProcess(newExePath, app)

	return
}
