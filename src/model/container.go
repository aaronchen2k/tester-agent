package model

import (
	"github.com/aaronchen2k/openstc-common/src/domain"
	_const "github.com/aaronchen2k/openstc-common/src/libs/const"
	"time"
)

type Container struct {
	BaseModel

	Mac     string
	Name    string
	Ip      string
	Port    int
	SshPort int
	VncPort int
	Status  _const.VmStatus
	Msg     string
	WorkDir string

	HostId  uint
	ImageId uint

	HostName   string
	DefPath    string
	ImagePath  string
	DiskSize   int
	MemorySize int

	DestroyAt        time.Time
	LastRegisterDate time.Time
}

func (Container) TableName() string {
	return "biz_vm"
}

func (Container) GenContainerTo(containerPo Container) (to _domain.Container) {
	to = _domain.Container{MacAddress: containerPo.Mac, Id: int(containerPo.ID), Name: containerPo.Name,
		DiskSize: containerPo.DiskSize, MemorySize: containerPo.MemorySize, ImagePath: containerPo.ImagePath}

	return
}

func (Container) GenPveReq(vmPo Container) (req _domain.PveReq) {
	req = _domain.PveReq{
		VmMacAddress: vmPo.Mac, VmId: int(vmPo.ID), VmUniqueName: vmPo.Name,
		VmDiskSize: vmPo.DiskSize, VmMemorySize: vmPo.MemorySize, VmBackingImage: vmPo.ImagePath}

	return
}
