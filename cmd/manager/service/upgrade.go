package manageService

import (
	"errors"
	"fmt"
	managerConf "github.com/aaronchen2k/tester/cmd/manager/utils/conf"
	managerConst "github.com/aaronchen2k/tester/cmd/manager/utils/const"
	managerVari "github.com/aaronchen2k/tester/cmd/manager/utils/vari"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_commonUtils "github.com/aaronchen2k/tester/internal/pkg/libs/common"
	_downloadUtils "github.com/aaronchen2k/tester/internal/pkg/libs/download"
	_fileUtils "github.com/aaronchen2k/tester/internal/pkg/libs/file"
	_i118Utils "github.com/aaronchen2k/tester/internal/pkg/libs/i118"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	_shellUtils "github.com/aaronchen2k/tester/internal/pkg/libs/shell"
	"github.com/mholt/archiver/v3"
	"log"
	"path"
	"strconv"
	"strings"
)

func CheckUpgrade(app managerConf.Client) {
	appDir := managerVari.WorkDir + app.Name + _const.PthSep
	_fileUtils.MkDirIfNeeded(appDir)

	verFile := appDir + "version.txt"
	versionUrl := fmt.Sprintf(managerConst.VersionDownloadURL, app.Name)
	_downloadUtils.Download(versionUrl, verFile)

	content := strings.TrimSpace(_fileUtils.ReadFile(verFile))
	if strings.Contains(strings.ToLower(content), "error") {
		_logUtils.Error(_i118Utils.I118Prt.Sprintf("download_fail", versionUrl, content))
		return
	}

	newVersionStr := convertVersion(content)
	newVersionNum, _ := strconv.ParseFloat(newVersionStr, 64)

	oldVersionStr := convertVersion(app.Version)
	oldVersionNum, _ := strconv.ParseFloat(oldVersionStr, 64)

	if oldVersionNum < newVersionNum {
		log.Println(_i118Utils.I118Prt.Sprintf("find_new_ver", content))

		pass, err := downloadApp(app.Name, content)
		if pass && err == nil {
			restartApp(app, content)
		}
	} else {
		log.Println(_i118Utils.I118Prt.Sprintf("no_need_to_upgrade", content))
	}
}

func downloadApp(app string, version string) (pass bool, err error) {
	appDir := managerVari.WorkDir + app + _const.PthSep

	os := _commonUtils.GetOs()
	if _commonUtils.IsWin() {
		os = fmt.Sprintf("win%d", strconv.IntSize)
	}
	url := fmt.Sprintf(managerConst.PackageDownloadURL, app, version, os, app)

	extractDir := appDir + "latest"
	pth := appDir + version + ".zip"

	err = _downloadUtils.Download(url, pth)
	if err != nil {
		return
	}

	md5Url := url + ".md5"
	md5Pth := pth + ".md5"
	err = _downloadUtils.Download(md5Url, md5Pth)
	if err != nil {
		return
	}

	pass = checkMd5(pth, md5Pth)
	if !pass {
		msg := _i118Utils.I118Prt.Sprintf("fail_md5_check", pth)
		log.Println(msg)
		err = errors.New(msg)
		return
	}

	_fileUtils.RmDir(extractDir)
	_fileUtils.MkDirIfNeeded(extractDir)
	err = archiver.Unarchive(pth, extractDir)

	if err != nil {
		log.Println(_i118Utils.I118Prt.Sprintf("fail_unzip", pth))
		return
	}

	return
}

func restartApp(app managerConf.Client, newVersion string) (err error) {
	appDir := managerVari.WorkDir + app.Name + _const.PthSep

	newExeDir := path.Join(appDir, "latest", app.Name)
	newExePath := path.Join(newExeDir, app.Name)
	if _commonUtils.IsWin() {
		newExePath += ".exe"
	}

	oldVersion := app.Version
	app.Version = newVersion

	_shellUtils.KillProcess(app.Name)
	StartApp(app)

	log.Println(_i118Utils.I118Prt.Sprintf("success_upgrade", oldVersion, newVersion))

	// update config file
	managerConf.SaveConfig(managerConf.Inst)

	return
}

func checkMd5(filePth, md5Pth string) (pass bool) {
	expectVal := _fileUtils.ReadFile(md5Pth)

	cmdStr := ""
	if _commonUtils.IsWin() {
		cmdStr = "CertUtil -hashfile " + filePth + " MD5"
	} else {
		cmdStr = "md5sum " + filePth + " | awk '{print $1}'"
	}
	actualVal, _ := _shellUtils.ExeSysCmd(cmdStr)
	if _commonUtils.IsWin() {
		arr := strings.Split(actualVal, "\n")
		if len(arr) > 1 {
			actualVal = strings.TrimSpace(strings.Split(actualVal, "\n")[1])
		}
	}

	return strings.TrimSpace(actualVal) == strings.TrimSpace(expectVal)
}

func convertVersion(str string) string {
	arr := strings.Split(str, ".")
	if len(arr) > 2 { // ignore 3th
		str = strings.Join(arr[:2], ".")
	}

	return str
}
