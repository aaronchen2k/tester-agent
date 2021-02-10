package service

import (
	"fmt"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_stringUtils "github.com/aaronchen2k/tester/internal/pkg/libs/string"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/repo"
	"math/rand"
	"strings"
)

type VmService struct {
	RpcService *RpcService `inject:""`
	ResService *ResService `inject:""`

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

func (s *VmService) CreateByQueue(queue model.Queue) (err error) {
	templ := s.VmTemplRepo.Get(queue.VmTemplId)
	node := s.NodeRepo.GetByIndent(templ.Node)
	cluster := s.ClusterRepo.Get(templ.Cluster)

	vmName := fmt.Sprintf("vm-%d", queue.ID)
	vm, err := s.ResService.CreateVm(vmName, templ, node, cluster)

	if err != nil {
		return
	}

	s.VmRepo.Save(vm)
	queue.VmId = vm.ID
	s.NodeRepo.LaunchVm(queue)

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
