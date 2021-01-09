package vmService

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"time"
)

var (
	vmMap = map[string]_domain.Vm{}
	tm    = time.Now().Unix()
)

func updateVms(vms []_domain.Vm) {
	names := map[string]bool{}

	for _, vm := range vms {
		name := vm.Name
		names[name] = true

		if _, ok := vmMap[name]; ok {
			v := vmMap[name]
			v.Status = vm.Status
			vmMap[name] = v
		} else {
			add(name, vm)
		}
	}

	keys := getKeys(vmMap)
	for _, key := range keys {
		// remove vms in map but not found this time
		if !names[key] {
			delete(vmMap, key)
			continue
		}

		// destroy and remove timeout vm
		v := vmMap[key]
		if time.Now().Unix()-v.FirstDetectedTime.Unix() > _const.VmTimeout*60 { // timeout
			Remove(v)
			delete(vmMap, key)
		}
	}
}

func add(name string, vm _domain.Vm) {
	if vm.FirstDetectedTime.IsZero() {
		vm.FirstDetectedTime = time.Now()
	}
	vmMap[name] = vm
}

func getKeys(m map[string]_domain.Vm) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
