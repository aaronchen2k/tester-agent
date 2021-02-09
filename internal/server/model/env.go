package model

import _const "github.com/aaronchen2k/tester/internal/pkg/const"

type EnvBase struct {
	BaseModel
	Code string `json:"code"`
	Ord  int    `json:"ord"`
}

type OsPlatform struct {
	EnvBase
	Name _const.OsPlatform `json:"name"`
}
type OsType struct {
	EnvBase
	Name       _const.OsType `json:"name"`
	OsPlatform string        `json:"osPlatform"`
}
type OsLang struct {
	EnvBase
	Name string `json:"name"`
}
type BrowserType struct {
	EnvBase
	Name _const.BrowserType `json:"name"`
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
