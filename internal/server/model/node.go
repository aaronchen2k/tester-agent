package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	"time"
)

type Node struct {
	BaseModel

	Ident string `gorm:"ident" json:"ident"`
	Type  string `gorm:"type" json:"type"`

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

func NewNode() Node {
	node := Node{}
	return node
}

func (Node) TableName() string {
	return "biz_node"
}
