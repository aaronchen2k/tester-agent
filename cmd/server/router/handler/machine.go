package handler

import (
	"fmt"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"github.com/aaronchen2k/tester/internal/server/domain"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/aaronchen2k/tester/internal/server/utils"
	go_proxmox "github.com/aaronchen2k/tester/vendors/github.com/joernott/go-proxmox"
	"github.com/kataras/iris/v12"
	"strconv"
)

type MachineController struct {
	Ctx            iris.Context
	HostService    *service.HostService    `inject:""`
	ProxmoxService *service.ProxmoxService `inject:""`
}

func NewMachineController() *MachineController {
	return &MachineController{}
}
func (c *MachineController) PostRegister() (result _domain.RpcResult) {
	var host _domain.Host
	if err := c.Ctx.ReadJSON(&host); err != nil {
		_logUtils.Error(err.Error())
		result.Fail("wrong request data")
		return
	}

	result = c.HostService.Register(host)
	return result
}

func (c *MachineController) List(ctx iris.Context) {
	rootNode := &domain.ResNode{Name: "测试机", Type: _const.ResRoot, Id: "0"}

	hosts, _ := c.HostService.ListAll("", 0, 0)
	for _, host := range hosts {
		id := strconv.Itoa(int(host.ID))

		hostNode := &domain.ResNode{Name: host.Name, Type: _const.ResHost,
			Id: id, Key: string(_const.ResHost) + "-" + id}
		rootNode.Children = append(rootNode.Children, hostNode)

		var err error
		address := fmt.Sprintf("%s:%d", host.Ip, host.Port)
		go_proxmox.Proxmox, err = go_proxmox.NewProxMox(address, host.Username, host.Password)
		if err != nil {
			_logUtils.Print("fail to connect proxmox, error: " + err.Error())
			break
		}

		nodes, _ := go_proxmox.Proxmox.Nodes()
		for _, node := range nodes {
			id := node.Id

			nodeNode := &domain.ResNode{Name: node.Node, Type: _const.ResNode,
				Id: node.Id, HostId: hostNode.Id, Key: string(_const.ResNode) + "-" + id}
			hostNode.Children = append(hostNode.Children, nodeNode)

			vms, _ := node.Qemu()
			for _, vm := range vms {
				vmId := strconv.FormatFloat(vm.VMId, 'f', 0, 64)

				vmNode := &domain.ResNode{Name: vm.Name, Type: _const.ResVm,
					Id: vmId, HostId: hostNode.Id, NodeId: nodeNode.Id, Key: string(_const.ResVm) + "-" + vmId}
				nodeNode.Children = append(nodeNode.Children, vmNode)
			}
		}
	}

	_, _ = ctx.JSON(agentUtils.ApiRes(200, "请求成功", rootNode))
}

func (c *MachineController) Get(ctx iris.Context) {

	_, _ = ctx.JSON(agentUtils.ApiRes(200, "请求成功", nil))
}
