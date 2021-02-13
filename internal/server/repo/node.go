package repo

import (
	"github.com/aaronchen2k/tester/internal/server/model"
	"gorm.io/gorm"
)

type NodeRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewNodeRepo() *NodeRepo {
	return &NodeRepo{}
}

func (r *NodeRepo) Get(id uint) (ret model.Node) {
	r.DB.Where("id=?", id).First(&ret)
	return
}
func (r *NodeRepo) GetByIndent(ident string) (ret model.Node) {
	r.DB.Where("ident=?", ident).First(&ret)
	return
}

func (r *NodeRepo) Create(node *model.Node) {
	r.DB.Model(&node).Create(node)
	return
}

func (r *NodeRepo) AddInstCount(id uint) {
	r.DB.Model(&model.Node{}).Where("id = ?", id).
		UpdateColumn("instCount", gorm.Expr("instCount + 1"))
}
