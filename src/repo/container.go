package repo

import (
	_domain "github.com/aaronchen2k/openstc-common/src/domain"
	_const "github.com/aaronchen2k/openstc-common/src/libs/const"
	"github.com/aaronchen2k/openstc/src/model"
	"gorm.io/gorm"
	"time"
)

func NewContainerRepo() *ContainerRepo {
	return &ContainerRepo{}
}

type ContainerRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func (r *ContainerRepo) Register(container _domain.Container) (err error) {
	// just update status by mac for exist container
	r.DB.Model(&model.Container{}).Where("mac=?", container.MacAddress).
		Updates(
			map[string]interface{}{"status": container.Status,
				"ip": container.PublicIp, "port": container.PublicPort, "workDir": container.WorkDir,
				"lastRegisterDate": time.Now(), "updatedAt": time.Now()})

	return
}

func (r *ContainerRepo) GetById(id uint) (container model.Container) {
	r.DB.Where("ID=?", id).First(&container)
	return
}

func (r *ContainerRepo) Save(po model.Container) {
	r.DB.Model(&po).Omit("").Create(&po)
	return
}

func (r *ContainerRepo) Launch(container _domain.Container) {
	r.DB.Model(&container).Where("id=?", container.Id).
		Updates(
			map[string]interface{}{"status": "launch", "imagePath": container.ImagePath,
				"defPath": container.DefPath, "updatedAt": time.Now()})

	return
}

func (r *ContainerRepo) UpdateStatusByNames(containers []string, status _const.VmStatus) {
	db := r.DB.Model(&model.Container{}).Where("name = IN (?)", containers)

	if status == "running" {
		db.Where("AND status != 'active'")
	}

	db.Updates(map[string]interface{}{"status": status, "updatedAt": time.Now()})
}

func (r *ContainerRepo) DestroyMissedContainersStatus(containers []string, hostId uint) {
	db := r.DB.Model(&model.Container{}).Where("hostId=? AND status!=?", hostId, "destroy")

	if len(containers) > 0 {
		db.Where("AND name NOT IN (?)", containers)
	}

	db.Updates(map[string]interface{}{"status": "destroy", "updatedAt": time.Now()})
}

func (r *ContainerRepo) FailToCreate(id uint, msg string) {
	r.DB.Model(&model.Container{}).
		Where("id=?", id).
		Updates(map[string]interface{}{"msg": _const.VmFailToCreate, "updatedAt": time.Now()})
}
