package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"time"
)

type Plan struct {
	BaseModel
	TestObject

	// job
	BuildType _const.BuildType
	Priority  int
	GroupId   uint

	// env
	Serials      []string  // for appium test, , specific several SN
	Environments []TestEnv // for selenium, appium test

	// test object
	ScriptUrl   string
	ScmAddress  string
	ScmAccount  string
	ScmPassword string

	AppUrl          string
	BuildCommands   string
	ResultFiles     string
	KeepResultFiles _domain.MyBool

	// status
	Progress _const.BuildProgress
	Status   _const.BuildStatus

	StartTime   time.Time
	PendingTime time.Time
	ResultTime  time.Time

	// desc
	PlanName string
	UserName string
}

func NewPlan() Plan {
	plan := Plan{
		Progress: _const.ProgressCreated,
		Status:   _const.StatusCreated,
	}
	return plan
}

func (Plan) TableName() string {
	return "biz_plan"
}
