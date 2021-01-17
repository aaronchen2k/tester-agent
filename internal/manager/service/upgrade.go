package manageService

import (
	"errors"
	"fmt"
	managerConfig "github.com/aaronchen2k/tester/internal/manager/config"
	_managerConst "github.com/aaronchen2k/tester/internal/manager/utils/const"
	_managerVari "github.com/aaronchen2k/tester/internal/manager/utils/vari"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_commonUtils "github.com/aaronchen2k/tester/internal/pkg/libs/common"
	_downloadUtils "github.com/aaronchen2k/tester/internal/pkg/libs/download"
	_fileUtils "github.com/aaronchen2k/tester/internal/pkg/libs/file"
	_i118Utils "github.com/aaronchen2k/tester/internal/pkg/libs/i118"
	_shellUtils "github.com/aaronchen2k/tester/internal/pkg/libs/shell"
	_vari "github.com/aaronchen2k/tester/internal/pkg/vari"
	"github.com/mholt/archiver/v3"
	"log"
	"strconv"
	"strings"
)

func CheckUpgrade(app string) {
	appDir := _vari.WorkDir + app + _const.PthSep
	_fileUtils.MkDirIfNeeded(appDir)

	versionFile := appDir + "version.txt"
	versionUrl := fmt.Sprintf(_managerConst.VersionDownloadURL, app)
	_downloadUtils.Download(versionUrl, versionFile)

	content := strings.TrimSpace(_fileUtils.ReadFile(versionFile))
	newVersionStr := convertVersion(content)
	newVersionNum, _ := strconv.ParseFloat(newVersionStr, 64)

	oldVersionStr := convertVersion(_managerVari.Config.AgentVersion)
	oldVersionNum, _ := strconv.ParseFloat(oldVersionStr, 64)

	if oldVersionNum < newVersionNum {
		log.Println(_i118Utils.I118Prt.Sprintf("find_new_ver", content))

		pass, err := downloadApp(app, content)
		if pass && err == nil {
			restartApp(app, content)
		}
	} else {
		log.Println(_i118Utils.I118Prt.Sprintf("no_need_to_upgrade", content))
	}
}

func downloadApp(app string, version string) (pass bool, err error) {
	appDir := _vari.WorkDir + app + _const.PthSep

	os := _commonUtils.GetOs()
	if _commonUtils.IsWin() {
		os = fmt.Sprintf("win%d", strconv.IntSize)
	}
	url := fmt.Sprintf(_managerConst.PackageDownloadURL, app, version, os, app)

	extractDir := appDir + version

	pth := extractDir + ".zip"
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

func restartApp(app string, newVersion string) (err error) {
	appDir := _vari.WorkDir + app + _const.PthSep

	newExePath := appDir + newVersion + _const.PthSep + app + _const.PthSep + app
	if _commonUtils.IsWin() {
		newExePath += ".exe"
	}

	var oldVersion = _managerVari.Config.AgentVersion
	_managerVari.Config.AgentVersion = newVersion

	KillAgentProcess(app)
	StartAgentProcess(newExePath, app)

	log.Println(_i118Utils.I118Prt.Sprintf("success_upgrade", oldVersion, newVersion))

	// update config file
	managerConfig.SaveConfig(_managerVari.Config)

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
	actualVal, _ := _shellUtils.ExeShell(cmdStr)
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
