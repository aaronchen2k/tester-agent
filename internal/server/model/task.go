package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"time"
)

type Task struct {
	BaseModel
	TestObject
	TestEnv

	// job
	BuildType _const.BuildType
	Priority  int
	GroupId   uint
	PlanId    uint

	// env
	Serial      string  // for appium test, specific a SN
	Environment TestEnv // for appium, selenium test

	// status
	Progress _const.BuildProgress
	Status   _const.BuildStatus

	StartTime   time.Time
	PendingTime time.Time
	ResultTime  time.Time

	// desc
	TaskName string
	UserName string
}

func NewTask(
	buildType _const.BuildType, priority int, groupId uint, planId uint,
	serial string, environment TestEnv,

	osPlatform _const.OsPlatform, osType _const.OsName, osLang _const.SysLang,
	browserType _const.BrowserType, browserVersion string,

	scriptUrl string, scmAddress string, scmAccount string, scmPassword string,
	resultFiles string, keepResultFiles _domain.MyBool,
	appUrl string, buildCommands string,

	taskName string, userName string) Task {

	queue := Task{
		Progress: _const.ProgressCreated,
		Status:   _const.StatusCreated,

		BuildType: buildType,
		Priority:  priority,
		GroupId:   groupId,
		PlanId:    planId,

		Serial:      serial,
		Environment: environment,

		TestEnv: TestEnv{
			OsPlatform:  osPlatform,
			OsName:      osType,
			OsLang:      osLang,
			BrowserType: browserType,
			BrowserVer:  browserVersion,
		},

		TestObject: TestObject{
			ScriptUrl:       scriptUrl,
			ScmAddress:      scmAddress,
			ScmAccount:      scmAccount,
			ScmPassword:     scmPassword,
			ResultFiles:     resultFiles,
			KeepResultFiles: keepResultFiles,
			AppUrl:          appUrl,
			BuildCommands:   buildCommands,
		},

		TaskName: taskName,
		UserName: userName,
	}
	return queue
}

func (Task) TableName() string {
	return "biz_task"
}
