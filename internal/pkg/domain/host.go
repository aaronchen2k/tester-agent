package _domain

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	"time"
)

type Host struct {
	Name string

	OsPlatform _const.OsPlatform
	OsType     _const.OsType
	OsLang     _const.SysLang

	OsVersion string
	OsBuild   string
	OsBits    string

	Ip      string
	Port    int
	WorkDir string

	SshPort int
	VncPort int
	Status  _const.HostStatus

	taskCount        int
	LastRegisterDate time.Time

	Vms []Vm
}
