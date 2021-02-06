package handler

import (
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
)

type ContainerImageCtrl struct {
	Ctx          iris.Context
	ImageService *service.DockerImageService `inject:""`
}

func NewContainerImageCtrl() *ContainerImageCtrl {
	return &ContainerImageCtrl{}
}
