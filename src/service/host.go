package service

import (
	"fmt"
	_domain "github.com/aaronchen2k/openstc-common/src/domain"
	_const "github.com/aaronchen2k/openstc-common/src/libs/const"
	_numbUtils "github.com/aaronchen2k/openstc-common/src/libs/numb"
	"github.com/aaronchen2k/openstc/src/model"
	"github.com/aaronchen2k/openstc/src/repo"
)

type HostService struct {
	HostRepo  *repo.HostRepo  `inject:""`
	ImageRepo *repo.ImageRepo `inject:""`
	VmRepo    *repo.VmRepo    `inject:""`
}

func NewHostService() *HostService {
	return &HostService{}
}

func (s *HostService) Register(host _domain.Host) (result _domain.RpcResult) {
	hostPo, err := s.HostRepo.Register(host)
	if err != nil {
		result.Fail(fmt.Sprintf("fail to register host %s ", host.Ip))
	}

	s.updateVmsStatus(host, hostPo.ID)

	return
}

func (s *HostService) GetValidForQueue(queue model.Queue) (hostId, backingImageId int) {
	imageIds1 := s.ImageRepo.QueryByOs(queue.OsPlatform, queue.OsType, queue.OsLang)
	imageIds2 := s.ImageRepo.QueryByBrowser(queue.BrowserType, queue.BrowserVersion)

	images := make([]int, 0)
	for id := range imageIds1 {
		if _numbUtils.FindIntInArr(id, imageIds2) {
			images = append(images, id)
		}
	}

	if len(images) == 0 {
		return
	}

	hostIds := s.getIdleHost()
	if len(hostIds) == 0 {
		return
	}

	hostId, backingImageId = s.HostRepo.QueryByImages(images, hostIds)

	return
}

func (s *HostService) getIdleHost() (ids []int) {
	// keys: hostId, vmCount
	hostToVmCountList := s.HostRepo.QueryIdle(_const.MaxVmOnHost)

	hostIds := make([]int, 0)
	for _, mp := range hostToVmCountList {
		hostId := mp["hostId"]
		hostIds = append(hostIds, hostId)
	}

	return hostIds
}

func (s *HostService) updateVmsStatus(host _domain.Host, hostId uint) {
	vmNames := make([]string, 0)
	runningVms, destroyVms, unknownVms := s.getVmsByStatus(host, vmNames)

	if len(runningVms) > 0 {
		s.VmRepo.UpdateStatusByNames(runningVms, _const.VmRunning)
	}
	if len(destroyVms) > 0 {
		s.VmRepo.UpdateStatusByNames(destroyVms, _const.VmDestroy)
	}
	if len(unknownVms) > 0 {
		s.VmRepo.UpdateStatusByNames(unknownVms, _const.VmUnknown)
	}

	// destroy vms already removed on agent side
	s.VmRepo.DestroyMissedVmsStatus(vmNames, hostId)

	return
}

func (s *HostService) getVmsByStatus(host _domain.Host, vmNames []string) (runningVms, destroyVms, unknownVms []string) {
	vms := host.Vms

	for _, vm := range vms {
		name := vm.Name
		status := vm.Status
		vmNames = append(vmNames, name)

		if status == _const.VmRunning {
			runningVms = append(runningVms, name)
		} else if status == _const.VmDestroy {
			destroyVms = append(destroyVms, name)
		} else if status == _const.VmUnknown {
			unknownVms = append(unknownVms, name)
		}
	}

	return
}

func (s *HostService) ListAll(keywords string, pageNo, pageSize int) (hosts []model.Host, total int64) {
	hosts, total, _ = s.HostRepo.Query(keywords, pageNo, pageSize)

	return
}
