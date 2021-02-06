package vmService

import (
	"fmt"
	"github.com/aaronchen2k/tester/internal/agent/cfg"
	agentConst "github.com/aaronchen2k/tester/internal/agent/utils/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_fileUtils "github.com/aaronchen2k/tester/internal/pkg/libs/file"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	_shellUtils "github.com/aaronchen2k/tester/internal/pkg/libs/shell"
	_stringUtils "github.com/aaronchen2k/tester/internal/pkg/libs/string"
	"github.com/go-errors/errors"
	"strconv"
	"strings"
)

func Define(vmUniqueName string) (err error) {
	defPath := agentConst.FolderDef + vmUniqueName
	cmd := fmt.Sprintf("virsh define %s.xml", defPath)
	_, err = _shellUtils.ExeShellInDir(cmd, agentConf.Inst.WorkDir)
	if err != nil {
		_logUtils.Errorf("fail to define vm, cmd %s, err %s.", cmd, err.Error())
		return
	}
	return
}

func Start(vmName string) (vncPort int, err error) {
	cmd := fmt.Sprintf("virsh start %s", vmName)
	_, err = _shellUtils.ExeShellInDir(cmd, agentConf.Inst.WorkDir)
	if err != nil {
		_logUtils.Errorf("fail to start vm, cmd %s, err %s.", cmd, err.Error())
		vncPort = -1
		return
	}

	vncPort, err = GetVncPort(vmName)
	return
}

func GetVncPort(vmName string) (port int, err error) {
	cmd := fmt.Sprintf("virsh vncdisplay %s", vmName)
	out, err := _shellUtils.ExeShellInDir(cmd, agentConf.Inst.WorkDir)
	if err != nil {
		_logUtils.Errorf("fail to start vm, cmd %s, err %s.", cmd, err.Error())
		return
	}

	if strings.Contains(out, "error: Domain is not running") {
		return -1, errors.New("error: Domain is not running")
	}

	arr := strings.Split(out, ":")
	if len(arr) > 1 {
		port, _ := strconv.Atoi(arr[1])
		return port, nil
	}

	return -1, errors.New(fmt.Sprintf("error: No port field, output is %s", out))
}

func genDefFromTempl(vmName string, memory int, vmCdrom string, vmCdrom2 string) (templContent, macAddress string) {
	vmTemplate := fmt.Sprintf("%s%s/common.xml", agentConf.Inst.WorkDir, agentConst.FolderTempl)

	templContent = _fileUtils.ReadFile(agentConf.Inst.WorkDir + agentConst.FolderTempl + vmTemplate)

	templContent = strings.Replace(templContent, "${name}", vmName, -1)
	templContent = strings.Replace(templContent, "${memory}", strconv.Itoa(memory), -1)

	rawPath := fmt.Sprintf("%s%s%s.qcow2", agentConf.Inst.WorkDir, agentConst.FolderImage, vmName)
	templContent = strings.Replace(templContent, "${rawPath}", rawPath, -1)

	cdromPath := fmt.Sprintf("%s%s%s", agentConf.Inst.WorkDir, agentConst.FolderIso, vmCdrom)
	templContent = strings.Replace(templContent, "${cdrom}", cdromPath, -1)

	cdromPath2 := fmt.Sprintf("%s%s%s", agentConf.Inst.WorkDir, agentConst.FolderIso, vmCdrom2)
	templContent = strings.Replace(templContent, "${cdrom2}", cdromPath2, -1)

	macAddress = genMacAddress()
	templContent = strings.Replace(templContent, "${mac}", macAddress, -1)

	machine := genMachine()
	templContent = strings.Replace(templContent, "${machine}", machine, -1)

	return
}

func genMacAddress() string {
	cmd := "dd if=/dev/urandom count=1 2>/dev/null | md5sum | sed 's/^\\(..\\)\\(..\\)\\(..\\)\\(..\\).*$/\\1:\\2:\\3:\\4/'"
	output, err := _shellUtils.ExeShell(cmd)
	if err != nil {
		_logUtils.Errorf("fail to exec cmd %s, err %s.", cmd, err.Error())
		return ""
	}

	mac := "fa:92:" + _stringUtils.TrimAll(output)
	return mac
}

func genMachine() string {
	cmd := "qemu-system-x86_64 -M ? | grep default | awk '{print $1}'"
	output, err := _shellUtils.ExeShell(cmd)
	if err != nil {
		_logUtils.Errorf("fail to exec cmd %s, err %s.", cmd, err.Error())
		return ""
	}

	machine := _stringUtils.TrimAll(output)
	return machine
}

func Remove(vm _domain.Vm) (err error) {
	err = Destroy(vm.Name)
	if err != nil {
		return
	}

	err = Undefine(vm.Name)
	if err != nil {
		return
	}

	err = RemoveDefImage(vm.Name)
	if err != nil {
		return
	}

	err = RemoveDefFile(vm.Name)

	return err
}

func Destroy(vmUniqueName string) (err error) {
	cmd := fmt.Sprintf("virsh destroy %s", vmUniqueName)
	_, err = _shellUtils.ExeShellInDir(cmd, agentConf.Inst.WorkDir)
	if err != nil {
		_logUtils.Errorf("fail to start vm, cmd %s, err %s.", cmd, err.Error())
		return
	}
	return
}

func Undefine(vmUniqueName string) (err error) {
	cmd := fmt.Sprintf("virsh undefine %s", vmUniqueName)
	_, err = _shellUtils.ExeShellInDir(cmd, agentConf.Inst.WorkDir)
	if err != nil {
		_logUtils.Errorf("fail to undefine vm, cmd %s, err %s.", cmd, err.Error())
		return
	}
	return
}

func RemoveDefFile(vmUniqueName string) (err error) {
	defPath := fmt.Sprintf("%s/%s.xml", agentConst.FolderDef, vmUniqueName)
	cmd := fmt.Sprintf("rm -rf %s", defPath)
	_, err = _shellUtils.ExeShellInDir(cmd, agentConf.Inst.WorkDir)
	if err != nil {
		_logUtils.Errorf("fail to remove define file, cmd %s, err %s.", cmd, err.Error())
		return
	}
	return
}
