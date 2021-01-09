package testService

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"time"
)

var (
	tasks     = make([]_domain.BuildTo, 0)
	tm        = time.Now()
	isRunning = false
)

func AddTask(task _domain.BuildTo) {
	tasks = append(tasks, task)
}

func PeekTask() _domain.BuildTo {
	return tasks[0]
}

func RemoveTask() (task _domain.BuildTo) {
	if len(tasks) == 0 {
		return task
	}

	task = tasks[0]
	tasks = tasks[1:]

	return task
}

func StartTask() {
	tm = time.Now()
	isRunning = true
}
func EndTask() {
	isRunning = false
}

func GetTaskSize() int {
	return len(tasks)
}

func CheckTaskRunning() bool {
	if time.Now().Unix()-tm.Unix() > _const.AgentRunTime*60*1000 {
		isRunning = false
	}
	return isRunning
}
