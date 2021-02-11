package repo

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
	"gorm.io/gorm"
	"time"
)

func NewVmRepo() *VmRepo {
	return &VmRepo{}
}

type VmRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func (r *VmRepo) Register(vm _domain.Vm) (err error) {
	r.DB.Model(&model.Vm{}).Where("name=?", vm.HostName).
		Updates(
			map[string]interface{}{"status": vm.Status, "workDir": vm.WorkDir,
				"ip": vm.PublicIp, "port": vm.PublicPort,
				"lastRegisterDate": time.Now(), "updatedAt": time.Now()})

	return
}

func (r *VmRepo) GetById(id uint) (vm model.Vm) {
	r.DB.Where("ID=?", id).First(&vm)
	return
}

func (r *VmRepo) Save(po *model.Vm) {
	r.DB.Model(po).Omit("").Create(po)
	return
}

func (r *VmRepo) UpdateStatus(vmId uint, status _const.VmStatus) {
	r.DB.Model(&model.Vm{}).Where("id=?", vmId).
		Updates(
			map[string]interface{}{"status": status})

	return
}

func (r *VmRepo) UpdateStatusByNames(vms []string, status _const.VmStatus) {
	db := r.DB.Model(&model.Vm{}).Where("name = IN (?)", vms)

	if status == "running" {
		db.Where("AND status != 'active'")
	}

	db.Updates(map[string]interface{}{"status": status, "updatedAt": time.Now()})
}

func (r *VmRepo) FailToCreate(id uint, msg string) {
	r.DB.Model(&model.Vm{}).
		Where("id=?", id).
		Updates(map[string]interface{}{"msg": _const.VmFailToCreate, "updatedAt": time.Now()})
}

func (r *VmRepo) QueryForDestroy() (vms []model.Vm) {
	tm := time.Now().Add(-time.Minute * _const.WaitForExecTime) // 60 min before

	r.DB.Where("status != ? AND created_at < ?",
		_const.VmDestroy, tm).
		Order("id").Find(&vms)
	return
}
