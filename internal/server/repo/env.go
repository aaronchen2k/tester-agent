package repo

import (
	"github.com/aaronchen2k/tester/internal/server/model"
	"gorm.io/gorm"
)

func NewEnvRepo() *EnvRepo {
	return &EnvRepo{}
}

type EnvRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func (r *EnvRepo) ListOsPlatform() (list []model.OsPlatform) {
	r.DB.Where("1=1").Order("ord ASC").Find(&list)
	return
}
func (r *EnvRepo) ListOsType() (list []model.OsType) {
	r.DB.Where("1=1").Order("ord ASC").Find(&list)
	return
}
func (r *EnvRepo) ListOsLang() (list []model.OsLang) {
	r.DB.Where("1=1").Order("ord ASC").Find(&list)
	return
}
func (r *EnvRepo) ListBrowserType() (list []model.BrowserType) {
	r.DB.Where("1=1").Order("ord ASC").Find(&list)
	return
}
