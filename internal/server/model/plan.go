package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	"github.com/aaronchen2k/tester/internal/server/model/base"
	"time"
)

type Plan struct {
	BaseModel
	base.TestObject

	// job
	BuildType _const.BuildType `json:"buildType,omitempty"`
	Priority  int              `json:"priority,omitempty"`
	GroupId   uint             `json:"groupId,omitempty"`

	// env
	Environments []base.TestEnv `gorm:"-"`

	// status
	Progress _const.BuildProgress `json:"progress,omitempty"`
	Status   _const.BuildStatus   `json:"status,omitempty"`

	StartTime   time.Time `json:"startTime,omitempty"`
	PendingTime time.Time `json:"pendingTime,omitempty"`
	ResultTime  time.Time `json:"resultTime,omitempty"`

	// desc
	PlanName string `json:"planName,omitempty"`
	UserName string `json:"userName,omitempty"`
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
