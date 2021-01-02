package model

import (
	_domain "github.com/aaronchen2k/openstc/internal/pkg/domain"
	_const "github.com/aaronchen2k/openstc/internal/pkg/libs/const"
	"time"
)

type Queue struct {
	BaseModel

	Priority int
	Serial   string
	VmId     uint

	BuildType      _const.BuildType
	OsPlatform     _const.OsPlatform
	OsType         _const.OsType
	OsLang         _const.OsLang
	BrowserType    _const.BrowserType
	BrowserVersion string

	ScriptUrl   string
	ScmAddress  string
	ScmAccount  string
	ScmPassword string

	AppUrl          string
	BuildCommands   string
	ResultFiles     string
	KeepResultFiles _domain.MyBool

	Progress _const.BuildProgress
	Status   _const.BuildStatus

	Retry int

	TaskName string
	UserName string

	StartTime   time.Time
	PendingTime time.Time
	ResultTime  time.Time
	TimeoutTime time.Time

	TaskId  uint
	GroupId uint
}

func NewQueue() Queue {
	queue := Queue{
		Progress: _const.ProgressCreated,
		Status:   _const.StatusCreated,
	}
	return queue
}
func NewQueueDetail(serial string, buildType _const.BuildType, groupId uint, taskId uint, taskPriority int,
	osPlatform _const.OsPlatform, osType _const.OsType,
	osLang _const.OsLang, browserType _const.BrowserType, browserVersion string,
	scriptUrl string, scmAddress string, scmAccount string, scmPassword string,
	resultFiles string, keepResultFiles _domain.MyBool, taskName string, userName string,
	appUrl string, buildCommands string) Queue {
	queue := Queue{
		Progress: _const.ProgressCreated,
		Status:   _const.StatusCreated,

		Serial:      serial,
		BuildType:   buildType,
		OsLang:      osLang,
		OsPlatform:  osPlatform,
		OsType:      osType,
		BrowserType: browserType,

		GroupId:         groupId,
		TaskId:          taskId,
		Priority:        taskPriority,
		ScriptUrl:       scriptUrl,
		ScmAddress:      scmAddress,
		ScmAccount:      scmAccount,
		ScmPassword:     scmPassword,
		ResultFiles:     resultFiles,
		KeepResultFiles: keepResultFiles,
		TaskName:        taskName,
		UserName:        userName,

		AppUrl:        appUrl,
		BuildCommands: buildCommands,
	}
	return queue
}

func (Queue) TableName() string {
	return "biz_queue"
}
