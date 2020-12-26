package model

import (
	"github.com/aaronchen2k/openstc-common/src/libs/const"
	"time"
)

type Host struct {
	BaseModel

	Name string

	OsPlatform _const.OsPlatform
	OsType     _const.OsType
	OsLang     _const.OsLang

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
}

func NewHost() Host {
	host := Host{}
	return host
}

func (Host) TableName() string {
	return "biz_host"
}
