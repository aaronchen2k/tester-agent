package base

import (
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
)

type TestObject struct {
	ScriptUrl   string `gorm:"scriptUrl" json:"scriptUrl,omitempty"`
	ScmAddress  string `gorm:"scmAddress" json:"scmAddress,omitempty"`
	ScmAccount  string `gorm:"scmAccount" json:"scmAccount,omitempty"`
	ScmPassword string `gorm:"scmPassword" json:"scmPassword,omitempty"`

	AppUrl          string         `gorm:"appUrl" json:"appUrl,omitempty"`
	BuildCommands   string         `gorm:"buildCommands" json:"buildCommands,omitempty"`
	ResultFiles     string         `gorm:"resultFiles" json:"resultFiles,omitempty"`
	KeepResultFiles _domain.MyBool `gorm:"keepResultFiles" json:"keepResultFiles,omitempty"`
}
