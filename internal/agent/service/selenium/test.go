package seleniumService

import (
	"fmt"
	constanct "github.com/aaronchen2k/openstc/internal/agent/libs/const"
	commonService "github.com/aaronchen2k/openstc/internal/agent/service/common"
	execService "github.com/aaronchen2k/openstc/internal/agent/service/exec"
	_domain "github.com/aaronchen2k/openstc/internal/pkg/domain"
	"strings"
)

func ExecTest(build *_domain.BuildTo) {
	result := _domain.RpcResult{}

	commonService.SetBuildWorkDir(build)

	// get script
	execService.GetTestScript(build)
	if build.ProjectDir == "" {
		result.Fail(fmt.Sprintf("failed to get test script, %#v。", build))
		return
	}

	// exec test
	parseBuildCommand(build)
	result = execService.ExcCommand(build)
	if !result.IsSuccess() {
		result.Fail(fmt.Sprintf("failed to ext test,\n dir: %s\n  cmd: \n%s",
			build.ProjectDir, build.BuildCommands))
	}

	// submit result
	execService.UploadResult(*build, result)
}

func parseBuildCommand(build *_domain.BuildTo) {
	// mvn clean test -Dtestng.suite=target/test-classes/baidu-test.xml
	//				  -DdriverType=${driverType}
	//		 		  -DdriverPath=${driverPath}

	command := strings.ReplaceAll(build.BuildCommands, constanct.BuildParamSeleniumDriverType, build.SeleniumDriverVersion)
	command = strings.ReplaceAll(command, constanct.BuildParamSeleniumDriverPath, build.SeleniumDriverPath)

	build.BuildCommands = command
}
