package execService

import (
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_fileUtils "github.com/aaronchen2k/tester/internal/pkg/libs/file"
	_gitUtils "github.com/aaronchen2k/tester/internal/pkg/libs/git"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/mholt/archiver/v3"
	"github.com/satori/go.uuid"
	"os"
	"strings"
)

func GetTestScript(build *_domain.BuildTo) _domain.RpcResult {
	if build.ScmAddress != "" {
		CheckoutCodes(build)
	} else if strings.Index(build.ScriptUrl, "http://") == 0 {
		DownloadCodes(build)
	} else {
		build.ProjectDir = _fileUtils.UpdateDir(build.ScriptUrl)
	}

	result := _domain.RpcResult{}
	result.Success("")
	return result
}

func CheckoutCodes(task *_domain.BuildTo) {
	url := task.ScmAddress
	userName := task.ScmAccount
	password := task.ScmPassword

	projectDir := task.WorkDir + _gitUtils.GetGitProjectName(url) + string(os.PathSeparator)

	_fileUtils.MkDirIfNeeded(projectDir)

	options := git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	}
	if userName != "" {
		options.Auth = &http.BasicAuth{
			Username: userName,
			Password: password,
		}
	}
	_, err := git.PlainClone(projectDir, false, &options)
	if err != nil {
		return
	}

	task.ProjectDir = projectDir
}

func DownloadCodes(task *_domain.BuildTo) {
	zipPath := task.WorkDir + uuid.NewV4().String() + _fileUtils.GetExtName(task.ScriptUrl)
	_fileUtils.Download(task.ScriptUrl, zipPath)

	scriptFolder := _fileUtils.GetZipSingleDir(zipPath)
	if scriptFolder != "" { // single dir in zip
		task.ProjectDir = task.WorkDir + scriptFolder
		archiver.Unarchive(zipPath, task.WorkDir)
	} else { // more then one dir, unzip to a folder
		fileNameWithoutExt := _fileUtils.GetFileNameWithoutExt(zipPath)
		task.ProjectDir = task.WorkDir + fileNameWithoutExt + string(os.PathSeparator)
		archiver.Unarchive(zipPath, task.ProjectDir)
	}
}
