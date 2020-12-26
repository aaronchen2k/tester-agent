package model

import (
	"github.com/aaronchen2k/openstc-common/src/domain"
	"github.com/aaronchen2k/openstc-common/src/libs/const"
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
