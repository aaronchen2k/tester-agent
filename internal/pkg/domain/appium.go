package _domain

import (
	"time"
)

type Appium struct {
	Name             string
	Version          string
	DeviceSerial     string
	ComputerIp       string
	AppiumPort       int
	LastRegisterDate time.Time
}
