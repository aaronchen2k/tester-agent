package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	"time"
)

type Cluster struct {
	BaseModel

	Name string `json:"name"`
	Type string `json:"type"`

	OsPlatform _const.OsPlatform `json:"osPlatform,omitempty"`
	OsType     _const.OsName     `json:"osType,omitempty"`
	OsLang     _const.SysLang    `json:"osLang,omitempty"`

	OsVersion string `json:"osVersion,omitempty"`
	OsBuild   string `json:"osBuild,omitempty"`
	OsBits    string `json:"osBits,omitempty"`

	Ip      string `json:"ip"`
	Port    int    `json:"port"`
	WorkDir string `json:"workDir,omitempty"`

	Username string `gorm:"username" json:"username"`
	Password string `gorm:"password" json:"password"`

	SshPort int               `json:"sshPort,omitempty"`
	VncPort int               `json:"vncPort,omitempty"`
	Status  _const.HostStatus `json:"status"`

	LastRegisterDate time.Time `json:"lastRegisterDate"`
	TaskCount        int       `gorm:"-" json:"taskCount"`
}

func NewCluster() Cluster {
	cluster := Cluster{}
	return cluster
}

func (Cluster) TableName() string {
	return "biz_cluster"
}
