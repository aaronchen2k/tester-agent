package _domain

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	"time"
)

type DeviceInst struct {
	DeviceSpec

	ComputerIp       string               `json:"computerIp"`
	ComputerPort     int                  `json:"computerPort"`
	AppiumPort       int                  `json:"appiumPort"`
	DeviceStatus     _const.DeviceStatus  `json:"deviceStatus"`
	AppiumStatus     _const.ServiceStatus `json:"appiumStatus"`
	LastRegisterDate time.Time            `json:"lastRegisterDate"`
}

type DeviceSpec struct {
	Serial           string `json:"serial"`
	Model            string `json:"model"`
	ApiLevel         int    `json:"apiLevel"`
	Version          string `json:"version"`
	Code             string `json:"code"`
	Os               string `json:"os"`
	Kernel           string `json:"kernel"`
	Ram              int    `json:"ram"`
	Rom              int    `json:"rom"`
	Cpu              string `json:"cpu"`
	Battery          int    `json:"battery"`
	Density          int    `json:"density"`
	DeviceIp         string `json:"deviceIp"`
	ResolutionHeight int    `json:"resolutionHeight"`
	ResolutionWidth  int    `json:"resolutionWidth"`
}
