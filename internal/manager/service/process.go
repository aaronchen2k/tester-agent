package manageService

import (
	"bytes"
	"fmt"
	_commonUtils "github.com/aaronchen2k/tester/internal/pkg/libs/common"
	_fileUtils "github.com/aaronchen2k/tester/internal/pkg/libs/file"
	_vari "github.com/aaronchen2k/tester/internal/pkg/vari"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func GetAgentProcess(app string) (string, error) {
	var cmd *exec.Cmd

	tmpl := ""
	cmdStr := ""
	if _commonUtils.IsWin() {
		tmpl = `tasklist`
		cmdStr = fmt.Sprintf(tmpl)

		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		tmpl = `ps -ef | grep "%s" | grep -v "grep" | awk '{print $2}'`
		cmdStr = fmt.Sprintf(tmpl, app)

		cmd = exec.Command("/bin/bash", "-c", cmdStr)
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	output := ""
	if _commonUtils.IsWin() {
		arr := strings.Split(out.String(), "\n")
		for _, line := range arr {
			if strings.Index(line, app+".exe") > -1 {
				arr2 := regexp.MustCompile(`\s+`).Split(line, -1)
				output = arr2[1]
				break
			}
		}
	} else {
		output = out.String()
	}

	return output, err
}

func KillAgentProcess(app string) (string, error) {
	var cmd *exec.Cmd

	tmpl := ""
	cmdStr := ""
	if _commonUtils.IsWin() {
		// tasklist | findstr ztf.exe
		tmpl = `taskkill.exe /f /im %s.exe`
		cmdStr = fmt.Sprintf(tmpl, app)

		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		tmpl = `ps -ef | grep "%s" | grep -v "grep" | awk '{print $2}' | xargs kill -9`
		cmdStr = fmt.Sprintf(tmpl, app)

		cmd = exec.Command("/bin/bash", "-c", cmdStr)
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	output := out.String()

	return output, err
}

func StartAgentProcess(execPath string, app string) (string, error) {
	execDir := _fileUtils.GetAbosutePath(filepath.Dir(execPath))

	portTag := "p"
	portNum := 8848

	tmpl := ""
	cmdStr := ""
	var cmd *exec.Cmd
	if _commonUtils.IsWin() {
		tmpl = `start cmd /c %s -%s %d ^1^> %snohup.%s.log ^2^>^&^1`
		cmdStr = fmt.Sprintf(tmpl, execPath, portTag, portNum, _vari.WorkDir, app)

		cmd = exec.Command("cmd", "/C", cmdStr)

	} else {
		cmd = exec.Command("nohup", execPath, "-"+portTag, strconv.Itoa(portNum))

		log := filepath.Join(_vari.WorkDir, "nohup."+app+".log")
		f, _ := os.Create(log)

		cmd.Stdout = f
		cmd.Stderr = f
	}

	cmd.Dir = execDir
	err := cmd.Start()
	return "", err
}
