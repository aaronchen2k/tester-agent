package repo

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	"github.com/aaronchen2k/tester/internal/server/model"
	"gorm.io/gorm"
	"time"
)

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{}
}

type TaskRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func (r *TaskRepo) Save(task *model.Task) (err error) {
	err = r.DB.Model(&task).Omit("").Create(&task).Error
	return
}

func (r *TaskRepo) SetProgress(taskId uint, progress _const.BuildProgress) (err error) {
	var data map[string]interface{}
	if progress == _const.ProgressInProgress {
		data = map[string]interface{}{"progress": progress, "start_time": time.Now()}
	} else {
		data = map[string]interface{}{"progress": progress, "pending_time": time.Now()}
	}

	r.DB.Model(model.Task{}).Where("id=?", taskId).Updates(data)
	return
}

func (r *TaskRepo) SetResult(taskId uint, progress _const.BuildProgress, status _const.BuildStatus) (err error) {
	var data = map[string]interface{}{"progress": progress, "result": status, "updatedTime": time.Now()}
	r.DB.Model(model.Task{}).Where("id=?", taskId).Updates(data)
	return
}
