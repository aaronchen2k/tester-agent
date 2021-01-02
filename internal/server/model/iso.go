package model

import (
	_const "github.com/aaronchen2k/openstc/internal/pkg/libs/const"
)

type Iso struct {
	BaseModel

	Name string
	Path string
	Size int

	OsPlatform _const.OsPlatform
	OsType     _const.OsType
	OsLang     _const.OsLang

	OsVersion string
	OsBuild   string
	OsBits    string

	ResolutionHeight  int
	ResolutionWidth   int
	suggestDiskSize   int
	suggestMemorySize int
}

func (Iso) TableName() string {
	return "biz_iso"
}
