package model

import "github.com/aaronchen2k/tester/internal/server/model/base"

type ContainerImage struct {
	BaseModel
	base.TestEnv

	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
	Size int    `json:"size,omitempty"`

	ResolutionHeight  int `json:"resolutionHeight,omitempty"`
	ResolutionWidth   int `json:"resolutionWidth,omitempty"`
	SuggestDiskSize   int `json:"suggestDiskSize,omitempty"`
	SuggestMemorySize int `json:"suggestMemorySize,omitempty"`

	SysIsoId    uint `json:"sysIsoId,omitempty"`
	DriverIsoId uint `json:"driverIsoId,omitempty"`

	Ident    string `json:"ident"`
	Computer string `json:"computer"`
	Cluster  string `json:"cluster"`
}

func (ContainerImage) TableName() string {
	return "biz_docker_image"
}
