package service

import (
	"crypto/rand"
	"fmt"
	_domain "github.com/aaronchen2k/openstc-common/src/domain"
	_const "github.com/aaronchen2k/openstc-common/src/libs/const"
	_stringUtils "github.com/aaronchen2k/openstc-common/src/libs/string"
	"github.com/aaronchen2k/openstc/src/model"
	"github.com/aaronchen2k/openstc/src/repo"
	"strings"
)

type VmService struct {
	RpcService *RpcService `inject:""`

	VmRepo    *repo.VmRepo    `inject:""`
	HostRepo  *repo.HostRepo  `inject:""`
	ImageRepo *repo.ImageRepo `inject:""`
	IsoRepo   *repo.IsoRepo   `inject:""`
	QueueRepo *repo.QueueRepo `inject:""`
}

func NewVmService() *VmService {
	return &VmService{}
}

func (s *VmService) Register(vm _domain.Vm) (result _domain.RpcResult) {
	err := s.VmRepo.Register(vm)
	if err != nil {
		result.Fail(fmt.Sprintf("fail to register host %s ", vm.MacAddress))
	}
	return
}

func (s *VmService) CreateRemote(hostId, backingImageId, queueId int) (result _domain.RpcResult) {
	host := s.HostRepo.Get(hostId)
	backingImage := s.ImageRepo.Get(backingImageId)
	sysIso := s.IsoRepo.Get(backingImage.SysIsoId)
	sysIsoPath := sysIso.Path

	driverIsoPath := ""
	if backingImage.OsPlatform == _const.OsWindows {
		driverIso := s.IsoRepo.Get(backingImage.DriverIsoId)
		driverIsoPath = driverIso.Path
	}

	mac := s.genValidMacAddress() // get a unique mac address
	vmName := s.genVmName(backingImage.Name)

	vmPo := model.Vm{Mac: mac, Name: vmName, HostName: host.Name,
		DiskSize: backingImage.SuggestDiskSize, MemorySize: backingImage.SuggestMemorySize,
		CdromSys: sysIsoPath, CdromDriver: driverIsoPath, BackingImagePath: backingImage.Path,
		HostId: uint(hostId), BackingImageId: uint(backingImageId),
		CdromSysId: backingImage.SysIsoId, CdromDriverId: backingImage.DriverIsoId}

	s.VmRepo.Save(vmPo) // save vm to db

	kvmRequest := model.GenPveReq(vmPo)
	result = s.RpcService.CreateVm(kvmRequest)

	if result.IsSuccess() { // success to create vm
		vmInResp := result.Payload.(_domain.Vm)
		s.VmRepo.Launch(vmInResp) // update vm status, mac address

		s.QueueRepo.UpdateVm(uint(queueId), vmPo.ID, _const.ProgressLaunchVm)
	} else {
		s.VmRepo.FailToCreate(vmPo.ID, result.Msg)

		s.QueueRepo.Pending(uint(queueId))
	}

	return
}

func (s *VmService) genVmName(imageName string) (name string) {
	uuid := strings.Replace(_stringUtils.NewUUID(), "-", "", -1)
	name = strings.Replace(imageName, "backing", uuid, -1)

	return
}

func (s *VmService) genValidMacAddress() (mac string) {
	for i := 0; i < 10; i++ {
		mac := s.genRandomMac()
		vm := s.VmRepo.GetByMac(mac)
		if vm.ID == 0 {
			return mac
		}
	}

	return "N/A"
}

func (s *VmService) genRandomMac() (mac string) {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	buf[0] |= 2
	mac = fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x\n", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])
	return
}
