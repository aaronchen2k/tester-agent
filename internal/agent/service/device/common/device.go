package deviceCommon

import (
	"github.com/aaronchen2k/tester/internal/agent/cfg"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"github.com/jinzhu/copier"
)

func SpecToDevInsts(specs []_domain.DeviceSpec) []_domain.DeviceInst {
	insts := make([]_domain.DeviceInst, 0)

	for _, spec := range specs {
		inst := _domain.DeviceInst{}
		copier.Copy(&inst, spec)

		inst.NodeIp = agentConf.Inst.Ip
		inst.NodePort = agentConf.Inst.Port

		insts = append(insts, inst)
	}

	return insts
}
