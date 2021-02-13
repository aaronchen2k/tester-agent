package seleniumService

import (
	"fmt"
	commonService "github.com/aaronchen2k/tester/internal/agent/service/common"
	execService "github.com/aaronchen2k/tester/internal/agent/service/exec"
	agentConst "github.com/aaronchen2k/tester/internal/agent/utils/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
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
	//		 		  -DdriverPath=${driverPath}  // computer in web page by browser type and version

	command := strings.ReplaceAll(build.BuildCommands, agentConst.BuildParamSeleniumDriverPath, build.SeleniumDriverPath)
	build.BuildCommands = command
}
