package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"time"
)

type Container struct {
	BaseModel

	Mac     string          `gorm:"mac" json:"mac,omitempty"`
	Name    string          `gorm:"name" json:"name,omitempty"`
	Ip      string          `gorm:"ip" json:"ip,omitempty"`
	Port    int             `gorm:"port" json:"port,omitempty"`
	SshPort int             `gorm:"sshPort" json:"sshPort,omitempty"`
	VncPort int             `gorm:"vncPort" json:"vncPort,omitempty"`
	Status  _const.VmStatus `gorm:"status" json:"status,omitempty"`
	Msg     string          `gorm:"msg" json:"msg,omitempty"`
	WorkDir string          `gorm:"workDir" json:"workDir,omitempty"`

	HostId  uint `gorm:"hostId" json:"hostId,omitempty"`
	ImageId uint `gorm:"imageId" json:"imageId,omitempty"`

	HostName   string `gorm:"hostName" json:"hostName,omitempty"`
	DefPath    string `gorm:"defPath" json:"defPath,omitempty"`
	ImagePath  string `gorm:"imagePath" json:"imagePath,omitempty"`
	DiskSize   int    `gorm:"diskSize" json:"diskSize,omitempty"`
	MemorySize int    `gorm:"memorySize" json:"memorySize,omitempty"`

	DestroyAt        time.Time `gorm:"destroyAt" json:"destroyAt,omitempty"`
	LastRegisterDate time.Time `gorm:"lastRegisterDate" json:"lastRegisterDate,omitempty"`

	Ident     string `gorm:"ident" json:"ident"`
	NodeId    uint   `gorm:"nodeId" json:"nodeId"`
	ClusterId uint   `gorm:"clusterId" json:"clusterId"`
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

func (Container) TableName() string {
	return "biz_container"
}
