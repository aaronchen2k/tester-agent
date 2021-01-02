package _domain

import ()

type Iso struct {
	Name       string `json:"name"`
	Path       string `json:"path"`
	Size       int    `json:"size"`
	OsPlatform string `json:"osPlatform"`
	OsType     string `json:"osType"`
	OsVersion  string `json:"osVersion"`
	OsBuild    string `json:"osBuild"`
	OsBits     string `json:"osBits"`
	OsLanguage string `json:"naosLanguageme"`
	Status     string `json:"status"`
}
