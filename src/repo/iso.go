package repo

import (
	"github.com/aaronchen2k/openstc/src/model"
	"gorm.io/gorm"
)

func NewIsoRepo() *IsoRepo {
	return &IsoRepo{}
}

type IsoRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func (r IsoRepo) Get(id uint) (iso model.Iso) {
	r.DB.Where("id=?", id).First(&iso)
	return
}
