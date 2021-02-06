package base

import _const "github.com/aaronchen2k/tester/internal/pkg/const"

type TestEnv struct {
	OsPlatform _const.OsPlatform `gorm:"osPlatform" json:"osPlatform,omitempty"`
	OsName     _const.OsName     `gorm:"osName" json:"osName,omitempty"`
	OsLevel    string            `gorm:"osLevel" json:"osLevel,omitempty"` // for mobile device only, e.x. android 11
	OsLang     _const.SysLang    `gorm:"osLang" json:"osLang,omitempty"`

	OsVer   string `gorm:"osVer" json:"osVer,omitempty"`
	OsBuild string `gorm:"osBuild" json:"osBuild,omitempty"`
	OsBits  string `gorm:"osBits" json:"osBits,omitempty"`

	BrowserType _const.BrowserType `gorm:"browserType" json:"browserType,omitempty"`
	BrowserVer  string             `gorm:"browserVer" json:"browserVer,omitempty"`
	BrowserLang _const.SysLang     `gorm:"browserLang" json:"browserLang,omitempty"`

	DeviceId uint   `gorm:"deviceId" json:"deviceId,omitempty"`
	Serial   string `gorm:"serial" json:"serial,omitempty"`

	VmTemplId     uint `gorm:"vmTemplId" json:"vmTemplId,omitempty"`
	DockerImageId uint `gorm:"dockerImageId" json:"dockerImageId,omitempty"`
}
