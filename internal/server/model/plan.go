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
	BuildType _const.BuildType `gorm:"buildType" json:"buildType,omitempty"`
	Priority  int              `gorm:"priority" json:"priority,omitempty"`
	GroupId   uint             `gorm:"groupId" json:"groupId,omitempty"`

	// env
	Environments []base.TestEnv `gorm:"-"` // for selenium, appium test

	// status
	Progress _const.BuildProgress `gorm:"progress" json:"progress,omitempty"`
	Status   _const.BuildStatus   `gorm:"status" json:"status,omitempty"`

	StartTime   time.Time `gorm:"startTime" json:"startTime,omitempty"`
	PendingTime time.Time `gorm:"pendingTime" json:"pendingTime,omitempty"`
	ResultTime  time.Time `gorm:"resultTime" json:"resultTime,omitempty"`

	// desc
	PlanName string `gorm:"planName" json:"planName,omitempty"`
	UserName string `gorm:"userName" json:"userName,omitempty"`
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
