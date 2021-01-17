package _const

import (
	"fmt"
	"os"
)

const (
	AppServer  = "tester-server"
	AppAgent   = "tester-agent"
	AppManager = "tester-manager"

	UserTokenExpireTime = 365 * 24 * 60 * 60 * 1000

	RegisterExpireTime = 5  // min
	WaitForExecTime    = 60 // min
	WaitForResultTime  = 30 // min
	VmTimeout          = 20 // min

	RetryTime    = 3
	AgentRunTime = 20 // sec

	WebCheckQueueTime = 10 // sec
	AgentCheckDevice  = 10 // sec

	MaxVmOnHost = 3
	RpcPort     = 8972

	UploadDir          = ""
	Sep_Of_Mac_Address = ":"

	Build_Command_Param_AppPath        = "${appPath}"
	Build_Command_Param_AppPackage     = "${appPackage}"
	Build_Command_Param_AppActivity    = "${appActivity}"
	Build_Command_Param_AppiumPort     = "${appiumPort}"
	Build_Command_Param_SeleniumDriver = "${driverPath}"

	PageSize = 15
)

var (
	PthSep = string(os.PathSeparator)

	LanguageDefault = "en"
	LanguageEN      = "en"
	LanguageZH      = "zh"
	EnRes           = fmt.Sprintf("res%sen%smessages.json", string(os.PathSeparator), string(os.PathSeparator))
	ZhRes           = fmt.Sprintf("res%szh%smessages.json", string(os.PathSeparator), string(os.PathSeparator))
)
