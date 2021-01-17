package _shellUtils

import (
	"bufio"
	"bytes"
	"fmt"
	_commonUtils "github.com/aaronchen2k/tester/internal/pkg/libs/common"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	_stringUtils "github.com/aaronchen2k/tester/internal/pkg/libs/string"
	"io"
	"os/exec"
	"strings"
)

func ExeShell(cmdStr string) (string, error) {
	return ExeShellInDir(cmdStr, "")
}

func ExeShellInDir(cmdStr string, dir string) (string, error) {
	var cmd *exec.Cmd
	if _commonUtils.IsWin() {
		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		cmd = exec.Command("/bin/bash", "-c", cmdStr)
	}
	if dir != "" {
		cmd.Dir = dir
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		_logUtils.Error(fmt.Sprintf("fail to exec command `%s` in dir `%s`, error `%#v`.", cmdStr, cmd.Dir, err))
	}

	str := _stringUtils.TrimAll(out.String())
	return str, err
}

func ExeShellWithOutput(cmdStr string) ([]string, error) {
	return ExeShellWithOutputInDir(cmdStr, "")
}

func ExeShellWithOutputInDir(cmdStr string, dir string) ([]string, error) {
	var cmd *exec.Cmd
	if _commonUtils.IsWin() {
		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		cmd = exec.Command("/bin/bash", "-c", cmdStr)
	}

	if dir != "" {
		cmd.Dir = dir
	}

	output := make([]string, 0)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return output, err
	}

	cmd.Start()

	if err != nil {
		return output, err
	}

	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		_logUtils.Info(strings.TrimRight(line, "\n"))
		output = append(output, line)
	}

	cmd.Wait()

	return output, nil
}
