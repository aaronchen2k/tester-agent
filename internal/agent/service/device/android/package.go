package androidService

import (
	"fmt"
	"github.com/aaronchen2k/tester/internal/agent/agentModel"
	_shellUtils "github.com/aaronchen2k/tester/internal/pkg/libs/shell"
	"regexp"
	"strings"
)

func GetAppInfo(apkPath string) agentModel.AndroidAppInfo {
	info := agentModel.AndroidAppInfo{ApkFile: apkPath}
	cmd := fmt.Sprintf("aapt dump badging %s | grep  -E 'package:|application-label:|launchable-activity:'", apkPath)
	out, err := _shellUtils.ExeShell(cmd)

	if err == nil {
		retrieveAppInfoFromAaptCmdOutput(out, &info)
	}

	return info
}

func retrieveAppInfoFromAaptCmdOutput(out string, info *agentModel.AndroidAppInfo) {
	lines := strings.Split(out, "\n")
	// package: name='com.applitools.helloworld.android' versionCode='2' versionName='1.1'
	// application-label:'ApplitoolsHelloWorld'
	// launchable-activity: name='com.applitools.helloworld.android.MainActivity'  label='' icon=''
	for _, line := range lines {
		line = strings.TrimSpace(line)
		lineLow := strings.ToLower(line)

		if lineLow == "" {
			continue
		}

		if strings.Index(lineLow, "package:") > -1 {
			re := regexp.MustCompile(`name='(.+)'.*versionCode='(.+)'.*versionName='(.+)'`)
			match := re.FindStringSubmatch(line)

			info.MainPackage = match[1]
			info.VersionCode = match[2]
			info.VersionName = match[3]

		} else if strings.Index(line, "application-label:") > -1 { // uppercase
			arr := strings.Split(line, ":")

			str := strings.TrimSpace(arr[1])
			appName := strings.ReplaceAll(str, "'", "")
			info.AppName = appName

		} else if strings.Index(lineLow, "launchable-activity:") > -1 {
			arr := strings.Split(line, "name='")
			arr = strings.Split(strings.TrimSpace(arr[1]), "'")

			info.MainActivity = strings.TrimSpace(arr[0])
		}
	}
}
