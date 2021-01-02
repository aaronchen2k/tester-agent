package model

import (
	_domain "github.com/aaronchen2k/openstc/internal/pkg/domain"
	_const "github.com/aaronchen2k/openstc/internal/pkg/libs/const"
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
