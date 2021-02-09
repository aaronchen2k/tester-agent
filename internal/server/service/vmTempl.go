package service

import (
	v1 "github.com/aaronchen2k/tester/cmd/server/router/v1"
	"github.com/aaronchen2k/tester/internal/server/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/model/base"
	"github.com/aaronchen2k/tester/internal/server/repo"
)

type VmTemplService struct {
	VmTemplRepo *repo.VmTemplRepo `inject:""`
}

func NewVmTemplService() *VmTemplService {
	return &VmTemplService{}
}

func (s *VmTemplService) GetByIdent(ident, node, cluster string) (templ model.VmTempl) {
	templ = s.VmTemplRepo.GetByIdent(ident, node, cluster)

	return
}

func (s *VmTemplService) CreateByNode(node domain.ResItem) (templ model.VmTempl) {
	templ = model.VmTempl{
		Name:    node.Name,
		Ident:   node.Ident,
		Node:    node.Node,
		Cluster: node.Cluster,
	}

	s.VmTemplRepo.Create(&templ)

	return
}

func (s *VmTemplService) Update(data v1.VmData) (err error) {
	po := model.VmTempl{
		Name:      data.Name,
		Ident:     data.Ident,
		BaseModel: model.BaseModel{ID: data.Id},
		TestEnv: base.TestEnv{
			OsPlatform: data.OsPlatform,
			OsType:     data.OsType,
			OsLang:     data.OsLang,
			OsVer:      data.OsVer,
			OsBits:     data.OsBits,
		},
	}

	if data.UpdateAll {
		err = s.VmTemplRepo.UpdateAllSameName(&po)
	} else {
		err = s.VmTemplRepo.Update(&po)
	}

	return
}
