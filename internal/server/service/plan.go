package service

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/repo"
)

type PlanService struct {
	TaskService *TaskService `inject:""`

	PlanRepo *repo.PlanRepo `inject:""`
}

func NewPlanService() *PlanService {
	return &PlanService{}
}

func (s *PlanService) List(keywords string, pageNo int, pageSize int) (plans []model.Plan, total int64) {
	plans, total = s.PlanRepo.Query(keywords, pageNo, pageSize)
	return
}

func (s *PlanService) Save(plan *model.Plan) (err error) {
	err = s.PlanRepo.Save(plan)

	s.TaskService.GenerateFromPlan(plan)

	return
}

func (s *PlanService) SetProgress(id uint, progress _const.BuildProgress) {
	s.PlanRepo.SetProgress(id, progress)
}
