package _domain

type PveReq struct {
	StartAfterCreated bool `json:"startAfterCreated"`

	ComputerIp   string `json:"-"`
	ComputerPort int    `json:"-"`

	VmPublicIp   string `json:"vmPublicIp"`
	VmPrivateIp  string `json:"vmPrivateIp"`
	VmPublicPort int    `json:"vmPublicPort"`

	VmMacAddress   string `json:"vmMacAddress"`
	VmTemplate     string `json:"vmTemplate"`
	VmUniqueName   string `json:"vmUniqueName"`
	VmMemorySize   int    `json:"vmMemorySize"`
	VmDiskSize     int    `json:"vmDiskSize"`
	VmCdromSys     string `json:"vmCdromSys"`
	VmCdromDriver  string `json:"vmCdromDriver"`
	VmBackingImage string `json:"vmBackingImage"`

	HostId           int `json:"hostId"`
	VmId             int `json:"vmId"`
	VmBackingImageId int `json:"vmBackingImageId"`
	VmCdromSysId     int `json:"vmCdromSysId"`
	VmCdromDriverId  int `json:"vmCdromDriverId"`
}

type PveResp struct {
	Code    int    `json:"code"`
	Msg     int    `json:"msg"`
	Name    string `json:"name"`
	VncPort int    `json:"vncPort"`
	Path    string `json:"path"`
	Mac     string `json:"mac"`
}
