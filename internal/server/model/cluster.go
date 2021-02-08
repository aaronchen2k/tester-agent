package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	"github.com/aaronchen2k/tester/internal/server/model/base"
	"time"
)

type Cluster struct {
	BaseModel
	base.TestEnv

	Name string `json:"name"`
	Type string `json:"type"`

	Ip      string `json:"ip"`
	Port    int    `json:"port"`
	WorkDir string `json:"workDir,omitempty"`

	Username string `json:"username"`
	Password string `json:"password"`

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
