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
	BuildType _const.BuildType `gorm:"buildType" json:"buildType,omitempty"`
	Priority  int              `gorm:"priority" json:"priority,omitempty"`
	GroupId   uint             `gorm:"groupId" json:"groupId,omitempty"`
	PlanId    uint             `gorm:"planId" json:"planId,omitempty"`

	// env
	Environment base.TestEnv `gorm:"-"` // for appium, selenium test

	// status
	Progress _const.BuildProgress `gorm:"progress" json:"progress,omitempty"`
	Status   _const.BuildStatus   `gorm:"status" json:"status,omitempty"`

	StartTime   time.Time `gorm:"startTime" json:"startTime,omitempty"`
	PendingTime time.Time `gorm:"pendingTime" json:"pendingTime,omitempty"`
	ResultTime  time.Time `gorm:"resultTime" json:"resultTime,omitempty"`

	// desc
	TaskName string `gorm:"taskName" json:"taskName,omitempty"`
	UserName string `gorm:"userName" json:"userName,omitempty"`
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
