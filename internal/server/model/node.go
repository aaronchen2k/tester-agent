package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	"time"
)

type Node struct {
	BaseModel

	Cluster   string `json:"cluster"`
	InstCount int    `gorm:"default:0" json:"instCount"`

	Ident string `json:"ident"`
	Type  string `json:"type"`

	Ip      string `json:"ip"`
	Port    int    `son:"port"`
	WorkDir string `json:"workDir,omitempty"`

	Username string `json:"username"`
	Password string `json:"password"`

	SshPort int               `json:"sshPort,omitempty"`
	VncPort int               `json:"vncPort,omitempty"`
	Status  _const.HostStatus `json:"status"`

	LastRegisterDate time.Time `json:"lastRegisterDate"`
	TaskCount        int       `gorm:"-" json:"taskCount"`
}

func NewNode() Node {
	node := Node{}
	return node
}

func (Node) TableName() string {
	return "biz_node"
}
