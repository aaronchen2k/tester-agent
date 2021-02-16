package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"github.com/aaronchen2k/tester/internal/server/model/base"
	"time"
)

type Build struct {
	BaseModel
	base.TestObject
	base.TestEnv
	QueueId uint `json:"queueId,omitempty"`

	// info
	Name string `json:"name,omitempty"`

	// job
	BuildType _const.BuildType `json:"buildType,omitempty"`
	Priority  int              `json:"priority,omitempty"`

	// env
	ComputerIp   string `json:"computerIp,omitempty"`
	ComputerPort int    `json:"computerPort,omitempty"`
	AppiumPort   int    `json:"appiumPort,omitempty"`

	// status
	StartTime    time.Time `json:"startTime,omitempty"`
	CompleteTime time.Time `json:"completeTime,omitempty"`

	Progress _const.BuildProgress `json:"progress,omitempty"`
	Status   _const.BuildStatus   `json:"status,omitempty"`
}

func NewBuild() Build {
	build := Build{
		Progress: _const.ProgressCreated,
		Status:   _const.StatusCreated,
	}

	return build
}

func NewBuildDetail(queueId uint, vmId uint, computerIp string, computerPort int) Build {
	build := Build{
		QueueId: queueId,

		ComputerIp:   computerIp,
		ComputerPort: computerPort,

		Progress: _const.ProgressCreated,
		Status:   _const.StatusCreated,

		TestEnv: base.TestEnv{VmId: vmId},
	}

	return build
}

func NewBuildTo(build Build) _domain.BuildTo {
	toValue := _domain.BuildTo{
		ID:           build.ID,
		QueueId:      build.QueueId,
		BuildType:    build.BuildType,
		Serial:       build.Serial,
		Priority:     build.Priority,
		ComputerIp:   build.ComputerIp,
		ComputerPort: build.ComputerPort,

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

		BrowserType: build.BrowserType,
		BrowserVer:  build.BrowserVer,
	}

	return toValue
}

func (Build) TableName() string {
	return "biz_build"
}
