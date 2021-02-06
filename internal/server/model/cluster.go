package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	"github.com/aaronchen2k/tester/internal/server/model/base"
	"time"
)

type Cluster struct {
	BaseModel
	base.TestEnv

	Name string `gorm:"name" json:"name"`
	Type string `gorm:"type" json:"type"`

	Ip      string `gorm:"ip" json:"ip"`
	Port    int    `gorm:"port" json:"port"`
	WorkDir string `gorm:"workDir" json:"workDir,omitempty"`

	Username string `gorm:"username" json:"username"`
	Password string `gorm:"password" json:"password"`

	SshPort int               `gorm:"sshPort" json:"sshPort,omitempty"`
	VncPort int               `gorm:"vncPort" json:"vncPort,omitempty"`
	Status  _const.HostStatus `gorm:"status" json:"status"`

	LastRegisterDate time.Time `gorm:"lastRegisterDate" json:"lastRegisterDate"`
	TaskCount        int       `gorm:"-" json:"taskCount"`
}

func NewCluster() Cluster {
	cluster := Cluster{}
	return cluster
}

func (Cluster) TableName() string {
	return "biz_cluster"
}
