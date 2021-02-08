package model

import _const "github.com/aaronchen2k/tester/internal/pkg/const"

type EnvBase struct {
	BaseModel
	Code string `gorm:"code" json:"code"`
}

type OsPlatform struct {
	EnvBase
	Name _const.OsPlatform `gorm:"name" json:"name"`
}
type OsType struct {
	EnvBase
	Name       _const.OsType `gorm:"name" json:"name"`
	OsPlatform string        `gorm:"osPlatform" json:"osPlatform"`
}
type OsLang struct {
	EnvBase
	Name string `gorm:"name" json:"name"`
}
type BrowserType struct {
	EnvBase
	Name _const.BrowserType `gorm:"name" json:"name"`
}

func (OsPlatform) TableName() string {
	return "biz_os_platform"
}
func (OsType) TableName() string {
	return "biz_os_type"
}
func (OsLang) TableName() string {
	return "biz_os_lang"
}
func (BrowserType) TableName() string {
	return "biz_browser_type"
}
