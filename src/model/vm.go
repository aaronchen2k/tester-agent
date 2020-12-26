package model

import (
	"github.com/aaronchen2k/openstc-common/src/domain"
	_const "github.com/aaronchen2k/openstc-common/src/libs/const"
	"time"
)

type Vm struct {
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

	HostId         uint
	BackingImageId uint
	CdromSysId     uint
	CdromDriverId  uint

	HostName         string
	DefPath          string
	ImagePath        string
	BackingImagePath string
	DiskSize         int
	MemorySize       int
	CdromSys         string
	CdromDriver      string

	ResolutionHeight int
	ResolutionWidth  int

	DestroyAt        time.Time
	LastRegisterDate time.Time
}

func (Vm) TableName() string {
	return "biz_vm"
}

func GenVmTo(vmPo Vm) (to _domain.Vm) {
	to = _domain.Vm{MacAddress: vmPo.Mac, Id: int(vmPo.ID), Name: vmPo.Name,
		DiskSize: vmPo.DiskSize, MemorySize: vmPo.MemorySize,
		CdromSys: vmPo.CdromSys, CdromDriver: vmPo.CdromDriver, BackingImagePath: vmPo.BackingImagePath}

	return
}

func GenPveReq(vmPo Vm) (req _domain.PveReq) {
	req = _domain.PveReq{
		VmMacAddress: vmPo.Mac, VmId: int(vmPo.ID), VmUniqueName: vmPo.Name,
		VmDiskSize: vmPo.DiskSize, VmMemorySize: vmPo.MemorySize,
		VmCdromSys: vmPo.CdromSys, VmCdromDriver: vmPo.CdromDriver, VmBackingImage: vmPo.BackingImagePath}

	return
}
