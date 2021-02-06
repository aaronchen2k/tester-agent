package model

import "github.com/aaronchen2k/tester/internal/server/model/base"

type Iso struct {
	BaseModel
	base.TestEnv

	Name string `gorm:"name" json:"name,omitempty"`
	Path string `gorm:"path" json:"path,omitempty"`
	Size int    `gorm:"size"`

	ResolutionHeight  int `gorm:"resolutionHeight" json:"resolutionHeight,omitempty"`
	ResolutionWidth   int `gorm:"resolutionWidth" json:"resolutionWidth,omitempty"`
	SuggestDiskSize   int `gorm:"suggestDiskSize" json:"suggestDiskSize,omitempty"`
	SuggestMemorySize int `gorm:"suggestMemorySize" json:"suggestMemorySize,omitempty"`
}

func (Iso) TableName() string {
	return "biz_iso"
}
