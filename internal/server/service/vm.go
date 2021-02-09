package service

import (
	"fmt"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	_stringUtils "github.com/aaronchen2k/tester/internal/pkg/libs/string"
	"github.com/aaronchen2k/tester/internal/server/repo"
	"math/rand"
	"strings"
)

type VmService struct {
	RpcService     *RpcService `inject:""`
	MachineService *ResService `inject:""`

	VmRepo      *repo.VmRepo      `inject:""`
	VmTemplRepo *repo.VmTemplRepo `inject:""`
	ClusterRepo *repo.ClusterRepo `inject:""`
	NodeRepo    *repo.NodeRepo    `inject:""`

	IsoRepo   *repo.IsoRepo   `inject:""`
	QueueRepo *repo.QueueRepo `inject:""`
}

func NewVmService() *VmService {
	return &VmService{}
}

func (s *VmService) Create(vmTemplId uint) {
	templ := s.VmTemplRepo.Get(vmTemplId)
	node := s.NodeRepo.Get(templ.Node)
	cluster := s.ClusterRepo.Get(templ.Cluster)

	vm, err := s.MachineService.CreateVm(templ, node, cluster)

	// TODO: save to db?
	_logUtils.Info(fmt.Sprintf("%#v, %s", vm, err.Error()))
}

func (s *VmService) CreateRemote(hostId, templImageId, queueId uint) (result _domain.RpcResult) {
	//host := s.ClusterRepo.Get(hostId)
	//vmTempl := s.VmTemplRepo.Get(templImageId)
	//sysIso := s.IsoRepo.Get(vmTempl.SysIsoId)
	//sysIsoPath := sysIso.Path
	//
	//driverIsoPath := ""
	//if vmTempl.OsPlatform == _const.OsWindows {
	//	driverIso := s.IsoRepo.Get(vmTempl.DriverIsoId)
	//	driverIsoPath = driverIso.Path
	//}
	//
	//mac := s.genValidMacAddress() // get a unique mac address
	//vmName := s.genVmName(vmTempl.Name)
	//
	//vmPo := model.Vm{Mac: mac, Name: vmName, HostName: host.Name,
	//	DiskSize: vmTempl.SuggestDiskSize, MemorySize: vmTempl.SuggestMemorySize,
	//	CdromSys: sysIsoPath, CdromDriver: driverIsoPath, BackingImagePath: vmTempl.Path,
	//	HostId: uint(hostId), TemplImageId: uint(templImageId),
	//	CdromSysId: vmTempl.SysIsoId, CdromDriverId: vmTempl.DriverIsoId}
	//
	//s.VmRepo.Save(vmPo) // save vm to db
	//
	//kvmRequest := model.GenPveReq(vmPo)
	//result = s.RpcService.CreateVm(kvmRequest)
	//
	//if result.IsSuccess() { // success to create vm
	//	vmInResp := result.Payload.(_domain.Vm)
	//	s.VmRepo.Launch(vmInResp) // update vm status, mac address
	//
	//	s.QueueRepo.UpdateVm(uint(queueId), vmPo.ID, _const.ProgressLaunchVm)
	//} else {
	//	s.VmRepo.FailToCreate(vmPo.ID, result.Msg)
	//
	//	s.QueueRepo.Pending(uint(queueId))
	//}

	return
}

func (s *VmService) Register(vm _domain.Vm) (result _domain.RpcResult) {
	err := s.VmRepo.Register(vm)
	if err != nil {
		result.Fail(fmt.Sprintf("fail to register host %s ", vm.MacAddress))
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
