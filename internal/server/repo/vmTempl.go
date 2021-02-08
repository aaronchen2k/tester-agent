package repo

import (
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/model/base"
	"gorm.io/gorm"
)

func NewVmTemplRepo() *VmTemplRepo {
	return &VmTemplRepo{}
}

type VmTemplRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func (r *VmTemplRepo) Get(id uint) (templ model.VmTempl) {
	r.DB.Where("id=?", id).First(&templ)
	return
}
func (r *VmTemplRepo) GetByIdent(ident string) (templ model.VmTempl) {
	r.DB.Model(&templ).Where("Ident=?", ident).First(&templ)
	return
}

func (r *VmTemplRepo) GetByEnv(env base.TestEnv) (templ model.VmTempl) {
	condition := r.convertEnvToVmTempl(env)
	r.DB.Where(&condition).First(&templ)
	return
}

func (r *VmTemplRepo) CreateTempl(templ *model.VmTempl) {
	r.DB.Model(&templ).Updates(templ)
	return
}

func (r *VmTemplRepo) UpdateTempl(templ *model.VmTempl) (err error) {
	r.DB.Model(&templ).Updates(templ)
	return
}

func (r *VmTemplRepo) convertEnvToVmTempl(env base.TestEnv) (templ model.VmTempl) {
	if env.OsPlatform != "" {
		templ.OsPlatform = env.OsPlatform
	}
	if env.OsType != "" {
		templ.OsType = env.OsType
	}
	if env.OsLevel != "" {
		templ.OsLevel = env.OsLevel
	}
	if env.OsLang != "" {
		templ.OsLang = env.OsLang
	}

	if env.OsVer != "" {
		templ.OsVer = env.OsVer
	}
	if env.OsBuild != "" {
		templ.OsBuild = env.OsBuild
	}
	if env.OsBits != "" {
		templ.OsBits = env.OsBits
	}

	if env.BrowserType != "" {
		templ.BrowserType = env.BrowserType
	}
	if env.BrowserVer != "" {
		templ.BrowserVer = env.BrowserVer
	}
	if env.BrowserLang != "" {
		templ.BrowserLang = env.BrowserLang
	}

	return
}
