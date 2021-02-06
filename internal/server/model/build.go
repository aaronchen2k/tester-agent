package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"time"
)

type Build struct {
	BaseModel

	QueueId uint `gorm:"serial" json:"serial,omitempty"`
	Queue   `sql:"-", gorm:"foreignkey:QueueId" json:"-"`

	// env
	NodeIp     string `gorm:"nodeIp" json:"nodeIp,omitempty"`
	NodePort   int    `gorm:"nodePort" json:"nodePort,omitempty"`
	AppiumPort int    `gorm:"appiumPort" json:"appiumPort,omitempty"`

	// status
	StartTime    time.Time `gorm:"startTime" json:"startTime,omitempty"`
	CompleteTime time.Time `gorm:"completeTime" json:"completeTime,omitempty"`

	Progress _const.BuildProgress `gorm:"progress" json:"progress,omitempty"`
	Status   _const.BuildStatus   `gorm:"status" json:"status,omitempty"`
}

func NewBuild() Build {
	build := Build{
		Progress: _const.ProgressCreated,
		Status:   _const.StatusCreated,
	}

	return build
}

func NewBuildDetail(queueId uint, nodeIp string, nodePort int) Build {
	build := Build{
		QueueId: queueId,

		NodeIp:   nodeIp,
		NodePort: nodePort,

		Progress: _const.ProgressCreated,
		Status:   _const.StatusCreated,
	}

	return build
}

func NewTestTo(build Build) _domain.BuildTo {
	toValue := _domain.BuildTo{
		ID:        build.ID,
		QueueId:   build.QueueId,
		BuildType: build.BuildType,
		Serial:    build.Serial,
		Priority:  build.Priority,
		NodeIp:    build.NodeIp,
		NodePort:  build.NodePort,

		AppUrl: build.AppUrl,

		Progress: _const.ProgressCreated,
		Status:   _const.StatusCreated,

		ScriptUrl:   build.ScriptUrl,
		ScmAddress:  build.ScmAddress,
		ScmAccount:  build.ScmAccount,
		ScmPassword: build.ScmPassword,

		BuildCommands:   build.BuildCommands,
		ResultFiles:     build.ResultFiles,
		KeepResultFiles: build.KeepResultFiles,
	}

	return toValue
}

func (Build) TableName() string {
	return "biz_build"
}
