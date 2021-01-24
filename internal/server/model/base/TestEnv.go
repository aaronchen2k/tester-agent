package base

import _const "github.com/aaronchen2k/tester/internal/pkg/const"

type TestEnv struct {
	OsPlatform _const.OsPlatform `json:"osPlatform,omitempty"`
	OsName     _const.OsName     `json:"osName,omitempty"`
	OsLevel    string            `json:"osLevel,omitempty"` // for mobile device only, e.x. android 11
	OsLang     _const.SysLang    `json:"osLang,omitempty"`

	OsVer   string `json:"osVer,omitempty"`
	OsBuild string `json:"osBuild,omitempty"`
	OsBits  string `json:"osBits,omitempty"`

	BrowserType _const.BrowserType `json:"browserType,omitempty"`
	BrowserVer  string             `json:"browserVer,omitempty"`
	BrowserLang _const.SysLang     `json:"browserLang,omitempty"`

	VmTemplId uint `json:"vmTemplId,omitempty"`
	VmId      uint `json:"vmId,omitempty"`

	DeviceId uint   `json:"deviceId,omitempty"`
	Serial   string `json:"serial,omitempty"`
}
