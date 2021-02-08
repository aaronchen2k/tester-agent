package base

import (
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
)

type TestObject struct {
	ScriptUrl   string `json:"scriptUrl,omitempty"`
	ScmAddress  string `json:"scmAddress,omitempty"`
	ScmAccount  string `json:"scmAccount,omitempty"`
	ScmPassword string `json:"scmPassword,omitempty"`

	AppUrl          string         `json:"appUrl,omitempty"`
	BuildCommands   string         `json:"buildCommands,omitempty"`
	ResultFiles     string         `json:"resultFiles,omitempty"`
	KeepResultFiles _domain.MyBool `json:"keepResultFiles,omitempty"`
}
