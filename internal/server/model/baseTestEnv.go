package model

import _const "github.com/aaronchen2k/tester/internal/pkg/const"

type TestEnv struct {
	OsPlatform _const.OsPlatform
	OsName     _const.OsName
	OsVer      string
	OsLang     _const.SysLang

	BrowserType _const.BrowserType
	BrowserVer  string
	BrowserLang _const.SysLang
}
