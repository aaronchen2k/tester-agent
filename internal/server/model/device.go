package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
)

type Device struct {
	BaseModel

	// from node register
	_domain.DeviceInst

	// info to maintain
	Name   string
	Make   string
	Brand  string
	Series string

	CpuMake         string
	CpuModel        string
	Memory          int
	Storage         int
	BatteryCapacity int

	OsType    _const.Platform
	OsLevel   string
	OsVersion string
}

func NewDevice() Device {
	device := Device{}

	return device
}

func (Device) TableName() string {
	return "biz_device"
}
