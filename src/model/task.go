package model

import (
	"github.com/aaronchen2k/openstc-common/src/domain"
	"github.com/aaronchen2k/openstc-common/src/libs/const"
	"time"
)

type Task struct {
	BaseModel

	Priority     int
	Serials      string // for appium test
	Environments string // for selenium test
	BuildType    _const.BuildType

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

	StartTime   time.Time
	PendingTime time.Time
	ResultTime  time.Time

	TaskName string
	UserName string

	GroupId uint
}

func NewTask() Task {
	task := Task{
		Progress: _const.ProgressCreated,
		Status:   _const.StatusCreated,
	}
	return task
}

func (Task) TableName() string {
	return "biz_task"
}
