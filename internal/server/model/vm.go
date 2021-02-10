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

	Mac     string          `json:"mac,omitempty"`
	Name    string          `json:"name,omitempty"`
	Ip      string          `json:"ip,omitempty"`
	Port    int             `json:"port,omitempty"`
	SshPort int             `json:"sshPort,omitempty"`
	VncPort int             `json:"vncPort,omitempty"`
	Status  _const.VmStatus `json:"status,omitempty"`
	Msg     string          `json:"msg,omitempty"`
	WorkDir string          `json:"workDir,omitempty"`

	HostId        uint `json:"hostId,omitempty"`
	TemplImageId  uint `json:"templImageId,omitempty"`
	CdromSysId    uint `json:"cdromSysId,omitempty"`
	CdromDriverId uint `json:"cdromDriverId,omitempty"`

	HostName         string `json:"hostName,omitempty"`
	DefPath          string `json:"defPath,omitempty"`
	ImagePath        string `json:"imagePath,omitempty"`
	BackingImagePath string `json:"backingImagePath,omitempty"`
	DiskSize         int    `json:"diskSize,omitempty"`
	MemorySize       int    `json:"memorySize,omitempty"`
	CdromSys         string `json:"cdromSys,omitempty"`
	CdromDriver      string `json:"cdromDriver,omitempty"`

	ResolutionHeight int `json:"resolutionHeight,omitempty"`
	ResolutionWidth  int `json:"resolutionWidth,omitempty"`

	DestroyAt        time.Time `json:"destroyAt,omitempty"`
	LastRegisterDate time.Time `gorm:"json:"lastRegisterDate,omitempty"`

	Ident     string `json:"ident"`
	Node      string `json:"node"`
	Cluster   string `json:"cluster"`
	NodeId    uint   `json:"nodeId"`
	ClusterId uint   `json:"clusterId"`
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
