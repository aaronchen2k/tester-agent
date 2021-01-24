package base

import (
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
)

type TestObject struct {
	ScriptUrl   string
	ScmAddress  string
	ScmAccount  string
	ScmPassword string

	AppUrl          string
	BuildCommands   string
	ResultFiles     string
	KeepResultFiles _domain.MyBool
	UserName        string
}
