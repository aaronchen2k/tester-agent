package service

import (
	v1 "github.com/aaronchen2k/tester/cmd/server/router/v1"
	"github.com/aaronchen2k/tester/internal/server/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/model/base"
	"github.com/aaronchen2k/tester/internal/server/repo"
)

type VmTemplService struct {
	VmTemplRepo  *repo.VmTemplRepo  `inject:""`
	ComputerRepo *repo.ComputerRepo `inject:""`
}

func NewVmTemplService() *VmTemplService {
	return &VmTemplService{}
}

func (s *VmTemplService) GetByIdent(ident, computer, cluster string) (templ model.VmTempl) {
	templ = s.VmTemplRepo.GetByIdent(ident, computer, cluster)

	return
}

func (s *VmTemplService) CreateByComputer(item domain.ResItem) (templ model.VmTempl) {
	templ = model.VmTempl{
		Name:     item.Name,
		Ident:    item.Ident,
		Computer: item.Computer,
		Cluster:  item.Cluster,
	}
	s.VmTemplRepo.Create(&templ)

	// create parent computer
	computer := s.ComputerRepo.GetByIndent(item.ComputerObj.Ident)
	if computer.ID == 0 {
		newComputer := model.Computer{
			Ident:   item.ComputerObj.Ident,
			Name:    item.ComputerObj.Name,
			Cluster: item.Cluster,
		}
		s.ComputerRepo.Create(&newComputer)
	}

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
