package service

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_numbUtils "github.com/aaronchen2k/tester/internal/pkg/libs/numb"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/repo"
)

type ClusterService struct {
	HostRepo  *repo.ClusterRepo `inject:""`
	ImageRepo *repo.ImageRepo   `inject:""`
	VmRepo    *repo.VmRepo      `inject:""`
}

func NewHostService() *ClusterService {
	return &ClusterService{}
}

func (s *ClusterService) GetValidForQueue(queue model.Queue) (hostId, backingImageId int) {
	imageIds1 := s.ImageRepo.QueryByOs(queue.OsPlatform, queue.OsName, queue.OsLang)
	imageIds2 := s.ImageRepo.QueryByBrowser(queue.BrowserType, queue.BrowserVer)

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

func (s *ClusterService) getIdleHost() (ids []int) {
	// keys: hostId, vmCount
	hostToVmCountList := s.HostRepo.QueryIdle(_const.MaxVmOnHost)

	hostIds := make([]int, 0)
	for _, mp := range hostToVmCountList {
		hostId := mp["hostId"]
		hostIds = append(hostIds, hostId)
	}

	return hostIds
}

func (s *ClusterService) updateVmsStatus(host _domain.Host, hostId uint) {
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

func (s *ClusterService) getVmsByStatus(host _domain.Host, vmNames []string) (runningVms, destroyVms, unknownVms []string) {
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

func (s *ClusterService) ListByType(tp string) (hosts []model.Cluster) {
	hosts, _ = s.HostRepo.QueryByType(tp)

	return
}

func (s *ClusterService) ListAll(keywords string, pageNo, pageSize int) (hosts []model.Cluster, total int64) {
	hosts, total, _ = s.HostRepo.Query(keywords, pageNo, pageSize)

	return
}
