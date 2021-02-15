package appiumService

import (
	"fmt"
	commonService "github.com/aaronchen2k/tester/internal/agent/service/common"
	deviceService "github.com/aaronchen2k/tester/internal/agent/service/device"
	androidService "github.com/aaronchen2k/tester/internal/agent/service/device/android"
	execService "github.com/aaronchen2k/tester/internal/agent/service/exec"
	agentConst "github.com/aaronchen2k/tester/internal/agent/utils/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"strconv"
	"strings"
)

func ExecTest(build *_domain.BuildTo) {
	result := _domain.RpcResult{}

	device, ok := deviceService.GetDevice(build.Serial)
	if !ok {
		result.Fail(fmt.Sprintf("can't find device %s@%s。", device.Serial, device.ComputerIp))
		return
	}

	commonService.SetBuildWorkDir(build)

	// get script
	execService.GetTestScript(build)
	if build.ProjectDir == "" {
		result.Fail(fmt.Sprintf("failed to get test script, %#v。", build))
		return
	}

	// get app
	execService.GetTestApp(build)

	// exec test
	parseBuildCommand(build)
	result = execService.ExcCommand(build)
	if !result.IsSuccess() {
		result.Fail(fmt.Sprintf("failed to ext test,\n dir: %s\n app: %s\n cmd: \n%s",
			build.ProjectDir, build.AppPath, build.BuildCommands))
	}

	// submit result
	execService.UploadResult(*build, result)
}

func parseBuildCommand(build *_domain.BuildTo) {
	// mvn clean test -Dtestng.suite=target/test-classes/android-test.xml
	//                -DappPath=${appPath}
	//                -DappPackage=${appPackage}
	//                -DappActivity==${appActivity}
	//                -DappiumPort==${appiumPort}

	app := androidService.GetAppInfo(build.AppPath)
	appPackage := app.MainPackage
	appActivity := app.MainActivity

	command := strings.ReplaceAll(build.BuildCommands, agentConst.BuildParamAppPath, build.AppPath)
	command = strings.ReplaceAll(command, agentConst.BuildParamAppPackage, appPackage)
	command = strings.ReplaceAll(command, agentConst.BuildParamAppActivity, appActivity)
	command = strings.ReplaceAll(command, agentConst.BuildParamAppiumPort, strconv.Itoa(build.AppiumPort))

	build.BuildCommands = command
}
