package domain

import _const "github.com/aaronchen2k/tester/internal/pkg/const"

type ResNode struct {
	Id       string         `json:"id"`
	Name     string         `json:"name"`
	Path     string         `json:"path"`
	Type     _const.ResType `json:"type"`
	Key      string         `json:"key"`
	Children []*ResNode     `json:"children"`

	IsTemplate bool   `json:"isTemplate"`
	HostId     string `json:"hostId,omitempty"`
	NodeId     string `json:"nodeId,omitempty"`

	Ip       string `json:"ip,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"-"`
	Password string `json:"-"`
}
