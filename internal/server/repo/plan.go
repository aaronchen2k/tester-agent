package repo

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"github.com/aaronchen2k/tester/internal/server/model"
	"gorm.io/gorm"
	"time"
)

func NewPlanRepo() *PlanRepo {
	return &PlanRepo{}
}

type PlanRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func (r *PlanRepo) Query(keywords string, pageNo int, pageSize int) (models []model.Plan, total int64) {
	query := r.DB.Select("id,title,referName,fileName,folder,path").Order("id ASC")
	if keywords != "" {
		query = query.Where("title LIKE ?", "%"+keywords+"%")
	}
	if pageNo > 0 {
		query = query.Offset((pageNo) * pageSize).Limit(pageSize)
	}

	err := query.Find(&models).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}
	err = r.DB.Model(&model.Plan{}).Count(&total).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *PlanRepo) Save(plan *model.Plan) (err error) {
	err = r.DB.Model(&plan).Omit("").Create(&plan).Error
	return
}

func (r *PlanRepo) SetProgress(planId uint, progress _const.BuildProgress) (err error) {
	var data map[string]interface{}
	if progress == _const.ProgressInProgress {
		data = map[string]interface{}{"progress": progress, "start_time": time.Now()}
	} else {
		data = map[string]interface{}{"progress": progress, "pending_time": time.Now()}
	}

	r.DB.Model(model.Plan{}).Where("id=?", planId).Updates(data)
	return
}

func (r *PlanRepo) SetResult(planId uint, progress _const.BuildProgress, status _const.BuildStatus) (err error) {
	var data = map[string]interface{}{"progress": progress, "result": status, "updatedTime": time.Now()}
	r.DB.Model(model.Plan{}).Where("id=?", planId).Updates(data)
	return
}
