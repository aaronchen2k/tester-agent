package model

import (
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"github.com/aaronchen2k/tester/internal/server/model/base"
)

type Device struct {
	BaseModel
	base.TestEnv

	// from node register
	_domain.DeviceInst

	// info to maintain
	Name   string `json:"name,omitempty"`
	Make   string `json:"make,omitempty"`
	Brand  string `json:"brand,omitempty"`
	Series string `json:"series,omitempty"`

	CpuMake         string `json:"cpuMake,omitempty"`
	CpuModel        string `json:"cpuModel,omitempty"`
	Memory          int    `json:"memory,omitempty"`
	Storage         int    `json:"storage,omitempty"`
	BatteryCapacity int    `json:"batteryCapacity,omitempty"`
}

func NewDevice() Device {
	device := Device{}

	return device
}

func (Device) TableName() string {
	return "biz_device"
}
