package tests

import (
	"fmt"
	"github.com/aaronchen2k/openstc/src/libs/proxmox"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func setup() {
	var err error
	proxmox.Proxmox, err = proxmox.NewProxMox("172.16.0.56:8006", "root", "P2ssw0rd")
	if err != nil {
		log.Println("fail to init proxmox, error: " + err.Error())
		return
	}

	log.Println("success to init proxmox")
}

func teardown() {

}

func TestNodes(t *testing.T) {
	nodes, e := proxmox.Proxmox.Nodes()
	if e != nil {
		t.Errorf("fail to get proxmox nodes, error: " + e.Error())
		return
	}

	assert.Equal(t, 1, len(nodes), "proxmox nodes number")

	var node proxmox.Node
	for key := range nodes {
		node = nodes[key]
	}

	vms, _ := node.Qemu()
	assert.Equal(t, 2, len(vms), "proxmox vms number")

	var vm proxmox.QemuVM
	for key, val := range vms {
		t.Log(fmt.Sprintf("find vm %s: %s", key, val.Name))
		if vm.Name == "" {
			vm = vms[key]
		}
	}

	name := vm.Name
	assert.Equal(t, "win10-tpl", name, "proxmox vm name")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}