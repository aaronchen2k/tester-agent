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
	Name   string `gorm:"name" json:"name,omitempty"`
	Make   string `gorm:"make" json:"make,omitempty"`
	Brand  string `gorm:"brand" json:"brand,omitempty"`
	Series string `gorm:"series" json:"series,omitempty"`

	CpuMake         string `gorm:"cpuMake" json:"cpuMake,omitempty"`
	CpuModel        string `gorm:"cpuModel" json:"cpuModel,omitempty"`
	Memory          int    `gorm:"memory" json:"memory,omitempty"`
	Storage         int    `gorm:"storage" json:"storage,omitempty"`
	BatteryCapacity int    `gorm:"batteryCapacity" json:"batteryCapacity,omitempty"`
}

func NewDevice() Device {
	device := Device{}

	return device
}

func (Device) TableName() string {
	return "biz_device"
}
