package model

import "github.com/aaronchen2k/tester/internal/server/model/base"

type Iso struct {
	BaseModel
	base.TestEnv

	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
	Size int    `json:"size,omitempty"`

	ResolutionHeight  int `json:"resolutionHeight,omitempty"`
	ResolutionWidth   int `json:"resolutionWidth,omitempty"`
	SuggestDiskSize   int `json:"suggestDiskSize,omitempty"`
	SuggestMemorySize int `json:"suggestMemorySize,omitempty"`
}

func (Iso) TableName() string {
	return "biz_iso"
}
