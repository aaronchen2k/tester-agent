package managerVari

import (
	"fmt"
	managerModel "github.com/aaronchen2k/tester/internal/manager/model"
	"os"
)

var (
	Verbose    = false
	WorkDir    = ""
	ConfigFile = ""
	LogFile    = ""

	Config managerModel.Config

	LogDir = fmt.Sprintf("log%s", string(os.PathSeparator))

	Actions = []string{"start", "stop", "restart", "install", "uninstall"}

	QiNiuURL           = "https://dl.cnezsoft.com/"
	VersionDownloadURL = QiNiuURL + "%s/version.txt"
	PackageDownloadURL = QiNiuURL + "%s/%s/%s/%s.zip"
)
