package _domain

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/libs/const"
	"time"
)

type Container struct {
	Id                int
	Name              string
	DiskSize          int
	MemorySize        int
	CdromSys          string
	CdromDriver       string
	DefPath           string
	ImagePath         string
	WorkDir           string
	PublicIp          string
	PublicPort        int
	MacAddress        string
	ResolutionHeight  int
	ResolutionWidth   int
	RpcPort           int
	SshPort           int
	VncPort           int
	DestroyAt         time.Time
	FirstDetectedTime time.Time
	HostId            int
	ImageId           int
	Status            _const.VmStatus
}
