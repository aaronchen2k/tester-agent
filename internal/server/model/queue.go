package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	"github.com/aaronchen2k/tester/internal/server/model/base"
	"time"
)

type Queue struct {
	BaseModel
	base.TestObject
	base.TestEnv

	// info
	Name string `json:"name,omitempty"`

	// job
	BuildType _const.BuildType `json:"buildType,omitempty"`
	Priority  int              `json:"priority,omitempty"`
	GroupId   uint             `json:"groupId,omitempty"`
	PlanId    uint             `json:"planId,omitempty"`
	TaskId    uint             `json:"taskId,omitempty"`

	// status
	Progress _const.BuildProgress `json:"progress,omitempty"`
	Status   _const.BuildStatus   `json:"status,omitempty"`

	StartTime   time.Time `json:"startTime,omitempty"`
	PendingTime time.Time `json:"pendingTime,omitempty"`
	ResultTime  time.Time `json:"resultTime,omitempty"`
	TimeoutTime time.Time `json:"timeoutTime,omitempty"`

	Retry int `json:"retry,omitempty"`
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
	name string, testObj base.TestObject,
	testEnv base.TestEnv) Queue {

	queue := Queue{
		BuildType: buildType,
		Priority:  priority,
		GroupId:   groupId,
		TaskId:    taskId,

		Name: name,

		TestEnv:    testEnv,
		TestObject: testObj,

		Progress: _const.ProgressCreated,
		Status:   _const.StatusCreated,
	}
	return queue
}

func (Queue) TableName() string {
	return "biz_queue"
}
