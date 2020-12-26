package service

import (
	"github.com/aaronchen2k/openstc/src/repo"
)

type ImageService struct {
	ImageRepo *repo.ImageRepo `inject:""`
}

func NewImageService() *ImageService {
	return &ImageService{}
}
