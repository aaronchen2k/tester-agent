package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"github.com/aaronchen2k/tester/internal/server/model/base"
	"time"
)

type Vm struct {
	BaseModel
	base.TestEnv

	Mac     string          `gorm:"mac" json:"mac,omitempty"`
	Name    string          `gorm:"name" json:"name,omitempty"`
	Ip      string          `gorm:"ip" json:"ip,omitempty"`
	Port    int             `gorm:"port" json:"port,omitempty"`
	SshPort int             `gorm:"sshPort" json:"sshPort,omitempty"`
	VncPort int             `gorm:"vncPort" json:"vncPort,omitempty"`
	Status  _const.VmStatus `gorm:"status" json:"status,omitempty"`
	Msg     string          `gorm:"msg" json:"msg,omitempty"`
	WorkDir string          `gorm:"workDir" json:"workDir,omitempty"`

	HostId        uint `gorm:"hostId" json:"hostId,omitempty"`
	TemplImageId  uint `gorm:"templImageId" json:"templImageId,omitempty"`
	CdromSysId    uint `gorm:"cdromSysId" json:"cdromSysId,omitempty"`
	CdromDriverId uint `gorm:"cdromDriverId" json:"cdromDriverId,omitempty"`

	HostName         string `gorm:"hostName" json:"hostName,omitempty"`
	DefPath          string `gorm:"defPath" json:"defPath,omitempty"`
	ImagePath        string `gorm:"imagePath" json:"imagePath,omitempty"`
	BackingImagePath string `gorm:"backingImagePath" json:"backingImagePath,omitempty"`
	DiskSize         int    `gorm:"diskSize" json:"diskSize,omitempty"`
	MemorySize       int    `gorm:"memorySize" json:"memorySize,omitempty"`
	CdromSys         string `gorm:"cdromSys" json:"cdromSys,omitempty"`
	CdromDriver      string `gorm:"cdromDriver" json:"cdromDriver,omitempty"`

	ResolutionHeight int `gorm:"resolutionHeight" json:"resolutionHeight,omitempty"`
	ResolutionWidth  int `gorm:"resolutionWidth" json:"resolutionWidth,omitempty"`

	DestroyAt        time.Time `gorm:"destroyAt" json:"destroyAt,omitempty"`
	LastRegisterDate time.Time `gorm:"lastRegisterDate" json:"lastRegisterDate,omitempty"`

	Ident     string `gorm:"ident" json:"ident"`
	NodeId    uint   `gorm:"nodeId" json:"nodeId"`
	ClusterId uint   `gorm:"clusterId" json:"clusterId"`
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

func (Vm) TableName() string {
	return "biz_vm"
}
