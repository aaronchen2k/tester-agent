package service

import (
	"github.com/aaronchen2k/tester/internal/server/repo"
)

type ImageService struct {
	ImageRepo *repo.ImageRepo `inject:""`
}

func NewImageService() *ImageService {
	return &ImageService{}
}
