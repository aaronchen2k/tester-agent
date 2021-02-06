package handler

import (
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
)

type VmTemplCtrl struct {
	Ctx          iris.Context
	ImageService *service.DockerImageService `inject:""`
}

func NewVmTemplCtrl() *VmTemplCtrl {
	return &VmTemplCtrl{}
}
