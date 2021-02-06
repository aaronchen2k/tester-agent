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

func (r *NodeRepo) Get(id uint) (node model.Node) {
	r.DB.Where("id=?", id).First(&node)
	return
}
