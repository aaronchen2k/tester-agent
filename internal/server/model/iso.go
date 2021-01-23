package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
)

type Iso struct {
	BaseModel

	Name string
	Path string
	Size int

	OsPlatform _const.OsPlatform
	OsType     _const.OsName
	OsLang     _const.SysLang

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
