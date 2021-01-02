package model

import (
	_const "github.com/aaronchen2k/openstc/internal/pkg/libs/const"
	"time"
)

type Host struct {
	BaseModel

	Name string `json:"name"`

	OsPlatform _const.OsPlatform `json:"osPlatform,omitempty"`
	OsType     _const.OsType     `json:"osType,omitempty"`
	OsLang     _const.OsLang     `json:"osLang,omitempty"`

	OsVersion string `json:"osVersion,omitempty"`
	OsBuild   string `json:"osBuild,omitempty"`
	OsBits    string `json:"osBits,omitempty"`

	Ip      string `json:"ip"`
	Port    int    `json:"port"`
	WorkDir string `json:"workDir,omitempty"`

	SshPort int               `json:"sshPort,omitempty"`
	VncPort int               `json:"vncPort,omitempty"`
	Status  _const.HostStatus `json:"status"`

	TaskCount        int       `json:"taskCount"`
	LastRegisterDate time.Time `json:"lastRegisterDate"`
}

func NewHost() Host {
	host := Host{}
	return host
}

func (Host) TableName() string {
	return "biz_host"
}
