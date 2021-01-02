package testService

import (
	appiumService "github.com/aaronchen2k/openstc/internal/agent/service/appium"
	seleniumService "github.com/aaronchen2k/openstc/internal/agent/service/selenium"
	_domain "github.com/aaronchen2k/openstc/internal/pkg/domain"
	_const "github.com/aaronchen2k/openstc/internal/pkg/libs/const"
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