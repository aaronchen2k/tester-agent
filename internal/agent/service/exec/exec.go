package execService

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/openstc/internal/agent/cfg"
	_domain "github.com/aaronchen2k/openstc/internal/pkg/domain"
	_fileUtils "github.com/aaronchen2k/openstc/internal/pkg/libs/file"
	_httpUtils "github.com/aaronchen2k/openstc/internal/pkg/libs/http"
	_logUtils "github.com/aaronchen2k/openstc/internal/pkg/libs/log"
	_shellUtils "github.com/aaronchen2k/openstc/internal/pkg/libs/shell"
	"github.com/satori/go.uuid"
	"strings"
)

func ExcCommand(build *_domain.BuildTo) _domain.RpcResult {
	cmdStr := build.BuildCommands
	out, err := _shellUtils.ExeShellWithOutputInDir(cmdStr, build.ProjectDir)

	result := _domain.RpcResult{}
	if err == nil {
		result.Success(strings.Join(out, "\n"))
	} else {
		result.Fail(fmt.Sprintf("fail to exec command, out %s, errpr %#v", out, err))
	}

	return result
}

func UploadResult(build _domain.BuildTo, result _domain.RpcResult) {
	zipFile := build.WorkDir + "testResult.zip"
	err := _fileUtils.ZipFiles(zipFile, build.ProjectDir, strings.Split(build.ResultFiles, ","))
	if err != nil {
		_logUtils.Errorf("fail to zip test results, dist '%s', dir %s, files '%s', error %#v",
			zipFile, build.ProjectDir, build.ResultFiles, err)
	}

	result.Payload = build

	uploadResultUrl := _httpUtils.GenUrl(agentConf.Inst.FarmServer, "build/upload")
	files := []string{zipFile}
	extraParams := map[string]string{}
	json, _ := json.Marshal(result)
	extraParams["result"] = string(json)

	_fileUtils.Upload(uploadResultUrl, files, extraParams)
}

func GetTestApp(build *_domain.BuildTo) _domain.RpcResult {
	result := _domain.RpcResult{}

	if strings.Index(build.AppUrl, "http://") == 0 {
		DownloadApp(build)
	} else {
		build.AppPath = build.AppUrl
	}

	if build.AppPath == "" {
		result.Fail(fmt.Sprintf("App获取错误，%s", build.AppUrl))
	} else {
		result.Success("")
	}

	return result
}

func DownloadApp(build *_domain.BuildTo) {
	path := build.WorkDir + uuid.NewV4().String() + _fileUtils.GetExtName(build.AppUrl)
	_fileUtils.Download(build.AppUrl, path)
	build.AppPath = path
}
