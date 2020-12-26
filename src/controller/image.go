package controller

import (
	"github.com/aaronchen2k/openstc/src/service"
	"github.com/kataras/iris/v12"
)

type ImageController struct {
	Ctx          iris.Context
	ImageService *service.ImageService `inject:""`
}

func NewImageController() *ImageController {
	return &ImageController{ImageService: service.NewImageService()}
}
