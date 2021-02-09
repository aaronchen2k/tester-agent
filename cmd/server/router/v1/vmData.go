package v1

import _const "github.com/aaronchen2k/tester/internal/pkg/const"

type VmData struct {
	Id    uint
	Name  string
	Ident string

	OsPlatform _const.OsPlatform
	OsType     _const.OsType
	OsLang     _const.SysLang
	OsBits     string

	OsLevel string
	OsVer   string
	OsBuild string

	UpdateAll bool
}
