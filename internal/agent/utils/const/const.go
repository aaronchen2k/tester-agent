package agentConst

import (
	"fmt"
	"os"
)

var (
	AppName    = "agent"
	ConfigVer  = 1
	ConfigFile = fmt.Sprintf("conf%stester.yaml", string(os.PathSeparator))

	EnRes = fmt.Sprintf("res%smessages_en.json", string(os.PathSeparator))
	ZhRes = fmt.Sprintf("res%smessages_zh.json", string(os.PathSeparator))

	BrowserDriverDir = "browser_driver"
	LogDir           = fmt.Sprintf("log%s", string(os.PathSeparator))

	BuildParamAppPath     = "${appPath}"
	BuildParamAppPackage  = "${appPackage}"
	BuildParamAppActivity = "${appActivity}"
	BuildParamAppiumPort  = "${appiumPort}"

	BuildParamSeleniumDriverPath = "${driverPath}"

	FolderIso   = "iso/"
	FolderImage = "image/"
	FolderDef   = "def/"
	FolderTempl = "templ/"

	ControlActions = []string{"start", "stop", "restart", "install", "uninstall"}
)
