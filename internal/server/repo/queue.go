package repo

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/libs/const"
	"github.com/aaronchen2k/tester/internal/server/model"
	"gorm.io/gorm"
	"strings"
	"time"
)

func NewQueueRepo() *QueueRepo {
	return &QueueRepo{}
}

type QueueRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func (r *QueueRepo) QueryForExec() (queues []model.Queue) {
	queues = make([]model.Queue, 0)

	r.DB.Where("progress=? OR progress=?",
		_const.ProgressCreated, _const.ProgressPending).Order("priority").Find(&queues)

	return
}
func (r *QueueRepo) QueryByTask(taskID uint) (queues []model.Queue) {
	queues = make([]model.Queue, 0)

	r.DB.Where("task_id=?", taskID).Order("id").Find(&queues)

	return
}

func (r *QueueRepo) GetQueue(id uint) (queue model.Queue) {
	r.DB.Where("id=?", id).First(&queue)

	return
}

func (r *QueueRepo) Save(queue *model.Queue) (err error) {
	err = r.DB.Model(&queue).
		Omit("StartTime", "PendingTime", "ResultTime", "TimeoutTime").
		Create(&queue).Error
	return
}

func (r *QueueRepo) DeleteInSameGroup(groupId uint, serials []string) (err error) {
	r.DB.Where("group_id=? AND serial IN (?)", groupId, strings.Join(serials, ",")).Delete(&model.Queue{})
	return
}

func (r *QueueRepo) Start(queue model.Queue) (err error) {
	r.DB.Model(&queue).Where("id=?", queue.ID).Updates(
		map[string]interface{}{"progress": _const.ProgressInProgress, "start_time": time.Now(), "retry": gorm.Expr("retry +1")})
	return
}
func (r *QueueRepo) Pending(queueId uint) (err error) {
	r.DB.Model(&model.Queue{}).Where("id=?", queueId).Updates(
		map[string]interface{}{"progress": _const.ProgressPending, "pending_time": time.Now()})
	return
}

func (r *QueueRepo) SetTimeout(id uint) (err error) {
	r.DB.Model(&model.Queue{}).Where("id=?", id).Updates(
		map[string]interface{}{"progress": _const.ProgressTimeout, "timeout_time": time.Now()})
	return
}

func (r *QueueRepo) QueryTimeout() (queues []model.Queue) {
	queues = make([]model.Queue, 0)

	r.DB.Where("(progress = ? AND unix_timestamp(NOW()) - unix_timestamp(pending_time) > ?)"+
		" OR (progress = ? AND unix_timestamp(NOW()) - unix_timestamp(start_time) > ?)",
		_const.ProgressPending, _const.WaitForExecTime*60*1000,
		_const.ProgressInProgress, _const.WaitForResultTime*60*1000).
		Order("priority").Find(&queues)
	return
}
func (r *QueueRepo) QueryTimeoutOrFailedForRetry() (queues []model.Queue) {
	queues = make([]model.Queue, 0)

	r.DB.Where("retry < ?"+" AND (progress = ? OR status = ? )",
		_const.RetryTime, _const.ProgressTimeout, _const.StatusFail).
		Order("priority").Find(&queues)
	return
}

func (r *QueueRepo) SetQueueStatus(queueId uint, progress _const.BuildProgress, status _const.BuildStatus) {
	r.DB.Model(&model.Queue{}).Where("id=?", queueId).Updates(
		map[string]interface{}{"progress": progress, "status": status, "result_time": time.Now(), "updated_at": time.Now()})
	return
}

func (r *QueueRepo) UpdateVm(queueId, vmId uint, status _const.BuildProgress) {
	r.DB.Model(&model.Queue{}).Where("id=?", queueId).Updates(
		map[string]interface{}{"vmId": vmId, "progress": status, "updated_at": time.Now()})
	return
}
