package service

import (
	"fmt"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"github.com/aaronchen2k/tester/internal/server/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
	serverUtils "github.com/aaronchen2k/tester/internal/server/utils/common"
	go_proxmox "github.com/aaronchen2k/tester/vendors/github.com/joernott/go-proxmox"
	"strconv"
)

type PveService struct {
}

func NewPveService() *PveService {
	return &PveService{}
}

func (s *PveService) ListVm(clusterNode *domain.ResItem) (vms []*model.Vm, err error) {
	s.GetNodeTree(clusterNode)

	return
}

func (s *PveService) CreateVm(name string, templ model.VmTempl, node model.Node, cluster model.Cluster) (vm model.Vm, err error) {
	address := fmt.Sprintf("%s:%d", cluster.Ip, cluster.Port)
	pve, err := go_proxmox.NewProxMox(address, cluster.Username, cluster.Password)
	if err != nil {
		_logUtils.Info("fail to connect proxmox, error: " + err.Error())
		return
	}

	newVmIdStr, _ := pve.NextVMId()
	templVm, err := pve.FindVM(templ.Ident)
	if err != nil {
		_logUtils.Info("fail to find vm, error: " + err.Error())
		return
	}

	newVmId, _ := strconv.ParseFloat(newVmIdStr, 64)
	vmHostName := serverUtils.GenVmHostName(name, templ.OsPlatform, templ.OsType, templ.OsLang)
	task, err := templVm.Clone(newVmId, vmHostName, node.Ident)
	if err != nil {
		_logUtils.Info("fail to clone vm, error: " + err.Error())
		return
	}

	newVm, err := pve.FindVM(newVmIdStr)
	err = newVm.Start()
	if err != nil {
		_logUtils.Info("fail to start vm, error: " + err.Error())
		return
	}

	_logUtils.Info("success to clone vm, task: " + task.ID)

	vm.Name = vmHostName
	vm.Ident = newVmIdStr
	vm.Node = node.Ident
	vm.Cluster = node.Cluster
	vm.NodeId = node.ID
	vm.ClusterId = cluster.ID

	return
}

func (s *PveService) GetNodeTree(clusterItem *domain.ResItem) (err error) {
	address := fmt.Sprintf("%s:%d", clusterItem.Ip, clusterItem.Port)
	go_proxmox.Proxmox, err = go_proxmox.NewProxMox(address, clusterItem.Username, clusterItem.Password)
	if err != nil {
		_logUtils.Print("fail to connect proxmox, error: " + err.Error())
		return
	}

	nodes, _ := go_proxmox.Proxmox.Nodes()
	for _, node := range nodes {
		ident := node.Id

		nodeItem := &domain.ResItem{Name: node.Node + "(节点)", Type: _const.ResNode,
			Ident: ident, Cluster: clusterItem.Ident, Key: string(_const.ResNode) + "-" + ident}
		clusterItem.Children = append(clusterItem.Children, nodeItem)

		vmFolderItem := &domain.ResItem{Name: "实例", Type: _const.ResFolder,
			Ident: ident + "-folder-vms", Key: ident + "-folder-vms"}
		nodeItem.Children = append(nodeItem.Children, vmFolderItem)

		templFolderItem := &domain.ResItem{Name: "模板", Type: _const.ResFolder,
			Ident: ident + "-folder-templs", Key: ident + "-folder-templs"}
		nodeItem.Children = append(nodeItem.Children, templFolderItem)

		vms, _ := node.Qemu()
		for _, vm := range vms {
			vmId := strconv.FormatFloat(vm.VMId, 'f', 0, 64)
			isTemplate := false
			if vm.Template == 1 {
				isTemplate = true
			}

			vmItem := &domain.ResItem{Name: vm.Name, Type: _const.ResVm, IsTemplate: isTemplate,
				Ident: vmId, Node: nodeItem.Ident, Cluster: clusterItem.Ident, Key: string(_const.ResVm) + "-" + vmId}

			if !isTemplate {
				vmFolderItem.Children = append(vmFolderItem.Children, vmItem)
			} else {
				templFolderItem.Children = append(templFolderItem.Children, vmItem)
			}
		}
	}

	return
}
