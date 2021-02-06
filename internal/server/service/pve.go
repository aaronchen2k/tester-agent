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

func (s *PveService) ListVm(clusterNode *domain.ResNode) (vms []*model.Vm, err error) {
	s.GetNodeTree(clusterNode)

	return
}

func (s *PveService) CreateVm(templ model.VmTempl, node model.Node, cluster model.Cluster) (vm model.Vm, err error) {
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
	vmHostName := serverUtils.GenVmHostName(templ.OsPlatform, templ.OsName, templ.OsLang)
	task, err := templVm.Clone(newVmId, vmHostName, node.Tag)
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

	vm.Ident = newVmIdStr
	vm.NodeId = node.ID
	vm.ClusterId = cluster.ID

	return
}

func (s *PveService) GetNodeTree(clusterNode *domain.ResNode) (root domain.ResNode, err error) {
	address := fmt.Sprintf("%s:%d", clusterNode.Ip, clusterNode.Port)
	go_proxmox.Proxmox, err = go_proxmox.NewProxMox(address, clusterNode.Username, clusterNode.Password)
	if err != nil {
		_logUtils.Print("fail to connect proxmox, error: " + err.Error())
		return
	}

	nodes, _ := go_proxmox.Proxmox.Nodes()
	for _, node := range nodes {
		id := node.Id

		nodeNode := &domain.ResNode{Name: node.Node + "(节点)", Type: _const.ResNode,
			Id: id, HostId: clusterNode.Id, Key: string(_const.ResNode) + "-" + id}
		clusterNode.Children = append(clusterNode.Children, nodeNode)

		vmFolderNode := &domain.ResNode{Name: "实例", Type: _const.ResFolder,
			Id: id + "-folder-vms", Key: id + "-folder-vms"}
		nodeNode.Children = append(nodeNode.Children, vmFolderNode)

		templFolderNode := &domain.ResNode{Name: "模板", Type: _const.ResFolder,
			Id: id + "-folder-templs", Key: id + "-folder-templs"}
		nodeNode.Children = append(nodeNode.Children, templFolderNode)

		vms, _ := node.Qemu()
		for _, vm := range vms {
			vmId := strconv.FormatFloat(vm.VMId, 'f', 0, 64)
			isTemplate := false
			if vm.Template == 1 {
				isTemplate = true
			}

			vmNode := &domain.ResNode{Name: vm.Name, Type: _const.ResVm, IsTemplate: isTemplate,
				Id: vmId, HostId: clusterNode.Id, NodeId: nodeNode.Id, Key: string(_const.ResVm) + "-" + vmId}

			if !isTemplate {
				vmFolderNode.Children = append(vmFolderNode.Children, vmNode)
			} else {
				templFolderNode.Children = append(templFolderNode.Children, vmNode)
			}
		}
	}

	return
}
