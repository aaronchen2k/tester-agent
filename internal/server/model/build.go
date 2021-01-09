package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"gorm.io/gorm"
	"time"
)

type Build struct {
	gorm.Model

	QueueId uint
	Queue   `sql:"-", gorm:"foreignkey:QueueId"`

	BuildType _const.BuildType
	VmId      uint

	Serial   string
	Priority int
	NodeIp   string
	NodePort int

	AppiumPort int

	StartTime    time.Time
	CompleteTime time.Time

	Progress _const.BuildProgress
	Status   _const.BuildStatus
}

func NewBuild() Build {
	build := Build{
		Progress: _const.ProgressCreated,
		Status:   _const.StatusCreated,
	}

	return build
}

func NewBuildDetail(queueId uint, vmId uint, buildType _const.BuildType,
	serial string, priority int, nodeIp string, nodePort int) Build {
	build := Build{
		QueueId:   queueId,
		VmId:      vmId,
		BuildType: buildType,
		Serial:    serial,
		Priority:  priority,
		NodeIp:    nodeIp,
		NodePort:  nodePort,

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
