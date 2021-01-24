package model

import "github.com/aaronchen2k/tester/internal/server/model/base"

type DockerImage struct {
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

func (DockerImage) TableName() string {
	return "biz_docker_image"
}
