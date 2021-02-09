package service

import (
	"github.com/aaronchen2k/tester/internal/server/repo"
)

type DockerImageService struct {
	DockerImageRepo *repo.ContainerImageRepo `inject:""`
}

func NewDockerImageService() *DockerImageService {
	return &DockerImageService{}
}
