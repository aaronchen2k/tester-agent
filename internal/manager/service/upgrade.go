package manageService

import (
	"errors"
	"fmt"
	configUtils "github.com/easysoft/zmanager/pkg/config"
	commonUtils "github.com/easysoft/zmanager/pkg/utils/common"
	constant "github.com/easysoft/zmanager/pkg/utils/const"
	"github.com/easysoft/zmanager/pkg/utils/download"
	fileUtils "github.com/easysoft/zmanager/pkg/utils/file"
	i118Utils "github.com/easysoft/zmanager/pkg/utils/i118"
	shellUtils "github.com/easysoft/zmanager/pkg/utils/shell"
	"github.com/easysoft/zmanager/pkg/utils/vari"
	"github.com/mholt/archiver/v3"
	"log"
	"strconv"
	"strings"
)

func CheckUpgrade(app string) {
	appDir := vari.WorkDir + app + constant.PthSep
	fileUtils.MkDirIfNeeded(appDir)

	versionFile := appDir + "version.txt"
	versionUrl := fmt.Sprintf(constant.VersionDownloadURL, app)
	downloadUtils.Download(versionUrl, versionFile)

	content := strings.TrimSpace(fileUtils.ReadFile(versionFile))
	newVersionStr := convertVersion(content)
	newVersionNum, _ := strconv.ParseFloat(newVersionStr, 64)

	var oldVersionNum float64
	if app == constant.ZTF {
		oldVersionStr := convertVersion(vari.Config.ZTFVersion)
		oldVersionNum, _ = strconv.ParseFloat(oldVersionStr, 64)
	} else if app == constant.ZenData {
		oldVersionStr := convertVersion(vari.Config.ZDVersion)
		oldVersionNum, _ = strconv.ParseFloat(oldVersionStr, 64)
	}

	if oldVersionNum < newVersionNum {
		log.Println(i118Utils.I118Prt.Sprintf("find_new_ver", content))

		pass, err := downloadApp(app, content)
		if pass && err == nil {
			restartApp(app, content)
		}
	} else {
		log.Println(i118Utils.I118Prt.Sprintf("no_need_to_upgrade", content))
	}
}

func downloadApp(app string, version string) (pass bool, err error) {
	appDir := vari.WorkDir + app + constant.PthSep

	os := commonUtils.GetOs()
	if commonUtils.IsWin() {
		os = fmt.Sprintf("win%d", strconv.IntSize)
	}
	url := fmt.Sprintf(constant.PackageDownloadURL, app, version, os, app)

	extractDir := appDir + version

	pth := extractDir + ".zip"
	err = downloadUtils.Download(url, pth)
	if err != nil {
		return
	}

	md5Url := url + ".md5"
	md5Pth := pth + ".md5"
	err = downloadUtils.Download(md5Url, md5Pth)
	if err != nil {
		return
	}

	pass = checkMd5(pth, md5Pth)
	if !pass {
		msg := i118Utils.I118Prt.Sprintf("fail_md5_check", pth)
		log.Println(msg)
		err = errors.New(msg)
		return
	}

	fileUtils.RmDir(extractDir)
	fileUtils.MkDirIfNeeded(extractDir)
	err = archiver.Unarchive(pth, extractDir)

	if err != nil {
		log.Println(i118Utils.I118Prt.Sprintf("fail_unzip", pth))
		return
	}

	return
}

func restartApp(app string, newVersion string) (err error) {
	appDir := vari.WorkDir + app + constant.PthSep

	newExePath := appDir + newVersion + constant.PthSep + app + constant.PthSep + app
	if commonUtils.IsWin() {
		newExePath += ".exe"
	}

	var oldVersion string
	if app == constant.ZTF {
		oldVersion = vari.Config.ZTFVersion
		vari.Config.ZTFVersion = newVersion
	} else if app == constant.ZenData {
		oldVersion = vari.Config.ZDVersion
		vari.Config.ZDVersion = newVersion
	}

	shellUtils.KillProcess(app)
	shellUtils.StartProcess(newExePath, app)

	log.Println(i118Utils.I118Prt.Sprintf("success_upgrade", oldVersion, newVersion))

	// update config file
	configUtils.SaveConfig(vari.Config)

	return
}

func checkMd5(filePth, md5Pth string) (pass bool) {
	expectVal := fileUtils.ReadFile(md5Pth)

	cmdStr := ""
	if commonUtils.IsWin() {
		cmdStr = "CertUtil -hashfile " + filePth + " MD5"
	} else {
		cmdStr = "md5sum " + filePth + " | awk '{print $1}'"
	}
	actualVal, _ := shellUtils.ExeSysCmd(cmdStr)
	if commonUtils.IsWin() {
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
