package model

import (
	_const "github.com/aaronchen2k/openstc-common/src/libs/const"
)

type BackingImage struct {
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
	SuggestDiskSize   int
	SuggestMemorySize int

	SysIsoId    uint
	DriverIsoId uint
}

func (BackingImage) TableName() string {
	return "biz_backing_image"
}
