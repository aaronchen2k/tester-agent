package repo

import (
	"github.com/aaronchen2k/tester/internal/server/model"
	"gorm.io/gorm"
)

type ComputerRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewComputerRepo() *ComputerRepo {
	return &ComputerRepo{}
}

func (r *ComputerRepo) Get(id uint) (ret model.Computer) {
	r.DB.Where("id=?", id).First(&ret)
	return
}
func (r *ComputerRepo) GetByIndent(ident string) (ret model.Computer) {
	r.DB.Where("ident=?", ident).First(&ret)
	return
}

func (r *ComputerRepo) Create(computer *model.Computer) {
	r.DB.Model(&computer).Create(computer)
	return
}

func (r *ComputerRepo) AddInstCount(id uint) {
	r.DB.Model(&model.Computer{}).Where("id = ?", id).
		UpdateColumn("instCount", gorm.Expr("instCount + 1"))
}
