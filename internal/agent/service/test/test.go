package testService

import (
	appiumService "github.com/aaronchen2k/tester/internal/agent/service/appium"
	seleniumService "github.com/aaronchen2k/tester/internal/agent/service/selenium"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
)

func Exec(build _domain.BuildTo) {
	StartTask()

	if build.BuildType == _const.AppiumTest {
		appiumService.ExecTest(&build)
	} else if build.BuildType == _const.SeleniumTest {
		seleniumService.ExecTest(&build)
	}

	RemoveTask()
	EndTask()
}
