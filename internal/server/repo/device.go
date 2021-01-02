package repo

import (
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
	"gorm.io/gorm"
)

func NewDeviceRepo() *DeviceRepo {
	return &DeviceRepo{}
}

type DeviceRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func (r *DeviceRepo) Register(device _domain.DeviceInst) (err error) {
	code := 1
	tx := r.DB.Begin()
	defer r.Defer(tx, &code)

	var po model.Device
	r.DB.Where("serial = ?", device.Serial).First(&po)

	if po.ID == 0 {
		err := r.DB.Model(&device).Omit("Ip").Create(&device).Error
		return err
	} else {
		r.DB.Model(&device).Updates(device)
		return nil
	}
}

func (r *DeviceRepo) GetBySerial(serial string) (device model.Device) {
	r.DB.Where("serial=?", serial).First(&device)
	return
}
