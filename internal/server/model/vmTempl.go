package model

import "github.com/aaronchen2k/tester/internal/server/model/base"

type VmTempl struct {
	BaseModel
	base.TestEnv

	Name string `gorm:"name" json:"name,omitempty"`
	Path string `gorm:"path" json:"path,omitempty"`
	Size int    `gorm:"size" json:"size,omitempty"`

	ResolutionHeight  int `gorm:"resolutionHeight" json:"resolutionHeight,omitempty"`
	ResolutionWidth   int `gorm:"resolutionWidth" json:"resolutionWidth,omitempty"`
	SuggestDiskSize   int `gorm:"suggestDiskSize" json:"suggestDiskSize,omitempty"`
	SuggestMemorySize int `gorm:"suggestMemorySize" json:"suggestMemorySize,omitempty"`

	SysIsoId    uint `gorm:"sysIsoId" json:"sysIsoId,omitempty"`
	DriverIsoId uint `gorm:"driverIsoId" json:"driverIsoId,omitempty"`

	Ident   string `gorm:"ident" json:"ident"`
	Node    string `gorm:"node" json:"node"`
	Cluster string `gorm:"cluster" json:"cluster"`
}

func (VmTempl) TableName() string {
	return "biz_vm_templ"
}
