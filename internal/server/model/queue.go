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

	// job
	BuildType _const.BuildType
	Priority  int
	GroupId   uint
	TaskId    uint

	// status
	Progress _const.BuildProgress
	Status   _const.BuildStatus

	StartTime   time.Time
	PendingTime time.Time
	ResultTime  time.Time
	TimeoutTime time.Time

	Retry int

	// desc
	QueueName string
	UserName  string
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
