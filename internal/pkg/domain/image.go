package _domain

import (
	"time"
)

type Image struct {
	Name             string    `json:"name"`
	Path             string    `json:"path"`
	Size             int       `json:"size"`
	OsPlatform       string    `json:"osPlatform"`
	OsType           string    `json:"osType"`
	OsVersion        string    `json:"osVersion"`
	OsBuild          string    `json:"osBuild"`
	OsBits           string    `json:"osBits"`
	OsLanguage       string    `json:"naosLanguageme"`
	ResolutionHeight int       `json:"resolutionHeight"`
	ResolutionWidth  int       `json:"resolutionWidth"`
	Status           string    `json:"status"`
	DestroyAt        time.Time `json:"distroyAt"`

	IsoId int `json:"isoId"`
}
