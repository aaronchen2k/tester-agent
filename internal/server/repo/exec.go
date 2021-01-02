package repo

import (
	"github.com/aaronchen2k/tester/internal/server/model"
	"gorm.io/gorm"
	"strings"
)

func NewExecRepo() *ExecRepo {
	return &ExecRepo{}
}

type ExecRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func (r *ExecRepo) Save(queue model.Queue) (err error) {
	err = r.DB.Model(&queue).Omit("StartTime", "PendingTime").Create(&queue).Error
	return
}

func (r *ExecRepo) DeleteInSameGroup(groupId uint, serials []string) (err error) {
	r.DB.Where("group_id=? AND serial IN (?)", groupId, strings.Join(serials, ",")).Delete(&model.Queue{})
	return
}
