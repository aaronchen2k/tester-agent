package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
)

type Image struct {
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
	SuggestDiskSize   int
	SuggestMemorySize int

	SysIsoId    uint
	DriverIsoId uint
}

func (Image) TableName() string {
	return "biz_backing_image"
}
