package repo

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type BuildRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewBuildRepo() *BuildRepo {
	return &BuildRepo{}
}

func (r *BuildRepo) GetBuild(id uint) (build model.Build) {
	r.DB.Where("ID=?", id).First(&build)

	return
}

func (r *BuildRepo) Save(build *model.Build) (err error) {
	err = r.DB.Model(&build).
		Omit("StartTime", "CompleteTime").
		Create(&build).Error
	return
}

func (r *BuildRepo) Start(build model.Build) (err error) {
	r.DB.Model(&build).Where("id=?", build.ID).Updates(
		map[string]interface{}{"progress": _const.ProgressInProgress, "start_time": time.Now()})
	return
}

func (r *BuildRepo) Delete(build model.Build) (err error) {
	r.DB.Delete(&build)
	return
}

func (r *BuildRepo) SaveResult(appiumTestTo _domain.BuildTo, resultPath string,
	progress _const.BuildProgress, status _const.BuildStatus, msg string) {

	r.DB.Model(&model.Build{}).Where("id=?", appiumTestTo.ID).Updates(
		map[string]interface{}{"progress": progress, "status": status, "resultPath": resultPath, "resultMsg": msg,
			"complete_time": time.Now()})
	return
}

func (r *BuildRepo) SetTimeoutByQueueId(queueId uint) {
	r.DB.Model(&model.Build{}).Where("queue_id=?", queueId).Updates(
		map[string]interface{}{"progress": _const.ProgressTimeout})
	return
}
