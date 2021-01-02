package model

import _const "github.com/aaronchen2k/openstc/internal/pkg/libs/const"

type Config struct {
	FarmServer string          `yaml:"FarmServer"`
	Ip         string          `yaml:"ip"`
	Port       int             `yaml:"port"`
	Platform   _const.Platform `yaml:"kvm"`
	MacAddress string          `yaml:"mac"`

	KvmDir  string
	WorkDir string
	LogDir  string
}
