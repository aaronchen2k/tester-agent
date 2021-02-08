package model

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"time"
)

type Container struct {
	BaseModel

	Mac     string          `json:"mac,omitempty"`
	Name    string          `json:"name,omitempty"`
	Ip      string          `json:"ip,omitempty"`
	Port    int             `json:"port,omitempty"`
	SshPort int             `json:"sshPort,omitempty"`
	VncPort int             `json:"vncPort,omitempty"`
	Status  _const.VmStatus `json:"status,omitempty"`
	Msg     string          `json:"msg,omitempty"`
	WorkDir string          `json:"workDir,omitempty"`

	HostId  uint `json:"hostId,omitempty"`
	ImageId uint `json:"imageId,omitempty"`

	HostName   string `json:"hostName,omitempty"`
	DefPath    string `json:"defPath,omitempty"`
	ImagePath  string `json:"imagePath,omitempty"`
	DiskSize   int    `json:"diskSize,omitempty"`
	MemorySize int    `json:"memorySize,omitempty"`

	DestroyAt        time.Time `json:"destroyAt,omitempty"`
	LastRegisterDate time.Time `json:"lastRegisterDate,omitempty"`

	Ident     string `json:"ident"`
	NodeId    uint   `json:"nodeId"`
	ClusterId uint   `json:"clusterId"`
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
