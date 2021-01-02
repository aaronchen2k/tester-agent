package handler

import (
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
)

type ImageController struct {
	Ctx          iris.Context
	ImageService *service.ImageService `inject:""`
}

func NewImageController() *ImageController {
	return &ImageController{}
}
