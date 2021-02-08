package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	"github.com/aaronchen2k/tester/internal/server/model/base"
	"time"
)

type Task struct {
	BaseModel
	base.TestObject
	base.TestEnv

	// job
	BuildType _const.BuildType `json:"buildType,omitempty"`
	Priority  int              `json:"priority,omitempty"`
	GroupId   uint             `json:"groupId,omitempty"`
	PlanId    uint             `json:"planId,omitempty"`

	// env
	Environment base.TestEnv `gorm:"-"`

	// status
	Progress _const.BuildProgress `json:"progress,omitempty"`
	Status   _const.BuildStatus   `json:"status,omitempty"`

	StartTime   time.Time `json:"startTime,omitempty"`
	PendingTime time.Time `json:"pendingTime,omitempty"`
	ResultTime  time.Time `json:"resultTime,omitempty"`

	// desc
	TaskName string `json:"taskName,omitempty"`
	UserName string `json:"userName,omitempty"`
}

func NewTask(
	buildType _const.BuildType, priority int, groupId uint, planId uint,
	taskName string, userName string,
	testEnv base.TestEnv, testObj base.TestObject) Task {

	task := Task{
		BuildType: buildType,
		Priority:  priority,
		GroupId:   groupId,
		PlanId:    planId,

		TaskName: taskName,
		UserName: userName,

		TestEnv:    testEnv,
		TestObject: testObj,

		Progress: _const.ProgressCreated,
		Status:   _const.StatusCreated,
	}
	return task
}

func (Task) TableName() string {
	return "biz_task"
}
