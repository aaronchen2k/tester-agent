package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"time"
)

type Queue struct {
	BaseModel
	TestObject
	TestEnv

	// job
	BuildType _const.BuildType
	Priority  int
	GroupId   uint
	TaskId    uint

	// env
	Serial      string  // for appium test, specific a SN
	Environment TestEnv // for appium, selenium test
	VmId        uint

	// status
	Progress _const.BuildProgress
	Status   _const.BuildStatus

	StartTime   time.Time
	PendingTime time.Time
	ResultTime  time.Time
	TimeoutTime time.Time

	Retry int

	// desc
	TaskName string
	UserName string
}

func NewQueue() Queue {
	queue := Queue{
		Progress: _const.ProgressCreated,
		Status:   _const.StatusCreated,
	}
	return queue
}
func NewQueueDetail(
	buildType _const.BuildType, priority int, groupId uint, taskId uint,
	serial string, environment TestEnv, vmId uint,
	osPlatform _const.OsPlatform, osType _const.OsName, osLang _const.SysLang,
	browserType _const.BrowserType, browserVersion string,

	scriptUrl string, scmAddress string, scmAccount string, scmPassword string,
	resultFiles string, keepResultFiles _domain.MyBool,
	appUrl string, buildCommands string,

	taskName string, userName string) Queue {

	queue := Queue{
		Serial:      serial,
		Environment: environment,
		VmId:        vmId,

		BuildType: buildType,
		Priority:  priority,
		GroupId:   groupId,
		TaskId:    taskId,

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

		Progress: _const.ProgressCreated,
		Status:   _const.StatusCreated,
	}
	return queue
}

func (Queue) TableName() string {
	return "biz_queue"
}
