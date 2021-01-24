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
	BuildType _const.BuildType
	Priority  int
	GroupId   uint
	PlanId    uint

	// env
	Environment base.TestEnv // for appium, selenium test

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
