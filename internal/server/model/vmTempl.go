package model

import "github.com/aaronchen2k/tester/internal/server/model/base"

type VmTempl struct {
	BaseModel
	base.TestEnv

	Name string
	Path string
	Size int

	ResolutionHeight  int
	ResolutionWidth   int
	SuggestDiskSize   int
	SuggestMemorySize int

	SysIsoId    uint
	DriverIsoId uint
}

func (VmTempl) TableName() string {
	return "biz_docker_templ"
}
