package _domain

import (
	"time"
)

type Appium struct {
	Name             string
	Version          string
	DeviceSerial     string
	NodeIp           string
	AppiumPort       int
	LastRegisterDate time.Time
}
