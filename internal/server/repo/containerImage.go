package repo

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/model/base"
	"gorm.io/gorm"
)

func NewContainerImageRepo() *ContainerImageRepo {
	return &ContainerImageRepo{}
}

type ContainerImageRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func (r *ContainerImageRepo) Get(id uint) (image model.ContainerImage) {
	r.DB.Where("id=?", id).First(&image)
	return
}

func (r *ContainerImageRepo) QueryByOs(platform _const.OsPlatform, osType _const.OsType, osLang _const.SysLang) (ids []int) {
	var db = r.DB.Model(model.ContainerImage{}).Where("NOT disabled AND NOT deleted")
	if platform != "" {
		db.Where("osPlatform = ?", platform)
	}
	if osType != "" {
		db.Where("osType = ?", osType)
	}
	if platform != "" {
		db.Where("osLang = ?", osLang)
	}

	db.Order("id ASC, createdAt ASC").Find(&ids)

	return
}

func (r *ContainerImageRepo) QueryByBrowser(browserType _const.BrowserType, ver string) (ids []int) {

	return
}

func (r *ContainerImageRepo) GetByEnv(env base.TestEnv) (image model.ContainerImage) {
	condition := r.convertEnvToVmTempl(env)
	r.DB.Where(&condition).First(&image)
	return
}

func (r *ContainerImageRepo) convertEnvToVmTempl(env base.TestEnv) (image model.ContainerImage) {
	if env.OsPlatform != "" {
		image.OsPlatform = env.OsPlatform
	}
	if env.OsType != "" {
		image.OsType = env.OsType
	}
	if env.OsLevel != "" {
		image.OsLevel = env.OsLevel
	}
	if env.OsLang != "" {
		image.OsLang = env.OsLang
	}

	if env.OsVer != "" {
		image.OsVer = env.OsVer
	}
	if env.OsBuild != "" {
		image.OsBuild = env.OsBuild
	}
	if env.OsBits != "" {
		image.OsBits = env.OsBits
	}

	if env.BrowserType != "" {
		image.BrowserType = env.BrowserType
	}
	if env.BrowserVer != "" {
		image.BrowserVer = env.BrowserVer
	}
	if env.BrowserLang != "" {
		image.BrowserLang = env.BrowserLang
	}

	return
}
