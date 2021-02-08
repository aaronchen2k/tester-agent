package managerConst

import (
	"fmt"
	"os"
	"os/user"
)

const (
	AppName = "manager"
)

var (
	PthSep = string(os.PathSeparator)

	userProfile, _ = user.Current()

	LanguageDefault = "en"
	LanguageEN      = "en"
	LanguageZH      = "zh"

	LogDir = fmt.Sprintf("log%s", string(os.PathSeparator))

	Actions = []string{"start", "stop", "restart", "install", "uninstall"}

	QiNiuURL           = "https://dl.cnezsoft.com/"
	VersionDownloadURL = QiNiuURL + "%s/version.txt"
	PackageDownloadURL = QiNiuURL + "%s/%s/%s/%s.zip"
)
