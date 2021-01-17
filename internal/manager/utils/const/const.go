package managerConst

var (
	Actions = []string{"start", "stop", "restart", "install", "uninstall"}

	QiNiuURL           = "https://dl.cnezsoft.com/"
	VersionDownloadURL = QiNiuURL + "%s/version.txt"
	PackageDownloadURL = QiNiuURL + "%s/%s/%s/%s.zip"
)
