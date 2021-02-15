package service

import (
	"fmt"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"github.com/aaronchen2k/tester/internal/server/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
	go_proxmox "github.com/aaronchen2k/tester/vendors/github.com/joernott/go-proxmox"
	"strconv"
)

type PveService struct {
}

func NewPveService() *PveService {
	return &PveService{}
}

func (s *PveService) ListVm(clusterItem *domain.ResItem) (vms []*model.Vm, err error) {
	s.GetNodeTree(clusterItem)

	return
}

func (s *PveService) CreateVm(hostName string, templ model.VmTempl, computer model.Computer, cluster model.Cluster) (
	vmIdent string, err error) {
	address := fmt.Sprintf("%s:%d", cluster.Ip, cluster.Port)
	pve, err := go_proxmox.NewProxMox(address, cluster.Username, cluster.Password)
	if err != nil {
		_logUtils.Info("fail to connect proxmox, error: " + err.Error())
		return
	}

	vmIdent, _ = pve.NextVMId()
	templVm, err := pve.FindVM(templ.Ident)
	if err != nil {
		_logUtils.Info("fail to find vm templ, error: " + err.Error())
		return
	}

	newVmId, _ := strconv.ParseFloat(vmIdent, 64)

	task, err := templVm.Clone(newVmId, hostName, computer.Ident)
	if err != nil {
		_logUtils.Info("fail to clone vm, error: " + err.Error())
		return
	}

	newVm, err := pve.FindVM(vmIdent)
	err = newVm.Start()
	if err != nil {
		_logUtils.Info("fail to start vm, error: " + err.Error())
		return
	}

	_logUtils.Info("success to clone vm, task: " + task.ID)

	return
}

func (s *PveService) DestroyVm(ident string, cluster model.Cluster) (err error) {
	address := fmt.Sprintf("%s:%d", cluster.Ip, cluster.Port)
	pve, err := go_proxmox.NewProxMox(address, cluster.Username, cluster.Password)
	if err != nil {
		_logUtils.Info("fail to connect proxmox, error: " + err.Error())
		return
	}

	vm, err := pve.FindVM(ident)
	if err != nil {
		_logUtils.Info("fail to find vm templ, error: " + err.Error())
		return
	}

	_, err = vm.Delete()
	if err != nil {
		_logUtils.Info("fail to delete vm, error: " + err.Error())
		return
	}

	return
}

func (s *PveService) GetNodeTree(clusterItem *domain.ResItem) (err error) {
	address := fmt.Sprintf("%s:%d", clusterItem.Ip, clusterItem.Port)
	go_proxmox.Proxmox, err = go_proxmox.NewProxMox(address, clusterItem.Username, clusterItem.Password)
	if err != nil {
		_logUtils.Print("fail to connect proxmox, error: " + err.Error())
		return
	}

	computers, _ := go_proxmox.Proxmox.Nodes()
	for _, computer := range computers {
		ident := computer.Id

		computerItem := &domain.ResItem{Name: computer.Node + "(节点)", Type: _const.ResComputer,
			Ident: ident, Cluster: clusterItem.Ident, Key: string(_const.ResComputer) + "-" + ident}
		clusterItem.Children = append(clusterItem.Children, computerItem)

		vmFolderItem := &domain.ResItem{Name: "实例", Type: _const.ResFolder,
			Ident: ident + "-folder-vms", Key: ident + "-folder-vms"}
		computerItem.Children = append(computerItem.Children, vmFolderItem)

		templFolderItem := &domain.ResItem{Name: "模板", Type: _const.ResFolder,
			Ident: ident + "-folder-templs", Key: ident + "-folder-templs"}
		computerItem.Children = append(computerItem.Children, templFolderItem)

		vms, _ := computer.Qemu()
		for _, vm := range vms {
			vmId := strconv.FormatFloat(vm.VMId, 'f', 0, 64)
			isTemplate := false
			if vm.Template == 1 {
				isTemplate = true
			}

			vmItem := &domain.ResItem{Name: vm.Name, Type: _const.ResVm, IsTemplate: isTemplate,
				Ident: vmId, Computer: computerItem.Ident, Cluster: clusterItem.Ident, Key: string(_const.ResVm) + "-" + vmId}

			if !isTemplate {
				vmFolderItem.Children = append(vmFolderItem.Children, vmItem)
			} else {
				templFolderItem.Children = append(templFolderItem.Children, vmItem)
			}
		}
	}

	return
}
