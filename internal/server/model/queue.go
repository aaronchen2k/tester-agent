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
	VmId uint `gorm:"vmId" json:"vmId,omitempty"`

	// job
	BuildType _const.BuildType `gorm:"buildType" json:"buildType,omitempty"`
	Priority  int              `gorm:"priority" json:"priority,omitempty"`
	GroupId   uint             `gorm:"groupId" json:"groupId,omitempty"`
	PlanId    uint             `gorm:"planId" json:"planId,omitempty"`
	TaskId    uint             `gorm:"taskId" json:"taskId,omitempty"`

	// status
	Progress _const.BuildProgress `gorm:"progress" json:"progress,omitempty"`
	Status   _const.BuildStatus   `gorm:"status" json:"status,omitempty"`

	StartTime   time.Time `gorm:"startTime" json:"startTime,omitempty"`
	PendingTime time.Time `gorm:"pendingTime" json:"pendingTime,omitempty"`
	ResultTime  time.Time `gorm:"resultTime" json:"resultTime,omitempty"`
	TimeoutTime time.Time `gorm:"timeoutTime" json:"timeoutTime,omitempty"`

	Retry int `gorm:"retry"`

	// desc
	QueueName string `gorm:"queueName"`
	UserName  string `gorm:"userName"`
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
	queueName string, userName string,
	testEnv base.TestEnv, testObj base.TestObject) Queue {

	queue := Queue{
		BuildType: buildType,
		Priority:  priority,
		GroupId:   groupId,
		TaskId:    taskId,

		QueueName: queueName,
		UserName:  userName,

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
