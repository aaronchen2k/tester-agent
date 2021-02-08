package service

import (
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/repo"
)

type EnvService struct {
	EnvRepo *repo.EnvRepo `inject:""`
}

func NewEnvService() *EnvService {
	return &EnvService{}
}

func (s *EnvService) List() (
	osPlatforms []model.OsPlatform, osTypes []model.OsType, osLangs []model.OsLang,
	browserTypes []model.BrowserType) {

	osPlatforms = s.EnvRepo.ListOsPlatform()
	osTypes = s.EnvRepo.ListOsType()
	osLangs = s.EnvRepo.ListOsLang()
	browserTypes = s.EnvRepo.ListBrowserType()

	return
}
