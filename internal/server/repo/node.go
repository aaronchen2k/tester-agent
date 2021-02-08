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

func (r *NodeRepo) Get(node string) (ret model.Node) {
	r.DB.Where("ident=?", node).First(&ret)
	return
}
