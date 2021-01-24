package service

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/repo"
)

type PlanService struct {
	PlanRepo  *repo.PlanRepo  `inject:""`
	QueueRepo *repo.QueueRepo `inject:""`
}

func NewPlanService() *PlanService {
	return &PlanService{}
}

func (s *PlanService) Save(plan model.Plan) model.Plan {
	s.PlanRepo.Save(&plan)
	return plan
}

func (s *PlanService) SetProgress(id uint, progress _const.BuildProgress) {
	s.PlanRepo.SetProgress(id, progress)
}
