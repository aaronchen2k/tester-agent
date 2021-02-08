package repo

import (
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/model/base"
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

func (r *DeviceRepo) Get(id uint) (device model.Device) {
	r.DB.Where("id=?", id).First(&device)
	return
}

func (r *DeviceRepo) GetBySerial(serial string) (device model.Device) {
	r.DB.Where("serial=?", serial).First(&device)
	return
}

func (r *DeviceRepo) GetByEnv(env base.TestEnv) (dev model.Device) {
	condition := r.convertEnvToVmTempl(env)
	r.DB.Where(&condition).First(&dev)
	return
}

func (r *DeviceRepo) convertEnvToVmTempl(env base.TestEnv) (dev model.Device) {
	if env.OsPlatform != "" {
		dev.OsPlatform = env.OsPlatform
	}
	if env.OsType != "" {
		dev.OsType = env.OsType
	}
	if env.OsLevel != "" {
		dev.OsLevel = env.OsLevel
	}
	if env.OsLang != "" {
		dev.OsLang = env.OsLang
	}

	if env.OsVer != "" {
		dev.OsVer = env.OsVer
	}
	if env.OsBuild != "" {
		dev.OsBuild = env.OsBuild
	}
	if env.OsBits != "" {
		dev.OsBits = env.OsBits
	}

	if env.BrowserType != "" {
		dev.BrowserType = env.BrowserType
	}
	if env.BrowserVer != "" {
		dev.BrowserVer = env.BrowserVer
	}
	if env.BrowserLang != "" {
		dev.BrowserLang = env.BrowserLang
	}

	return
}
