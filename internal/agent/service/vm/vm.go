package vmService

import (
	"fmt"
	"github.com/aaronchen2k/tester/internal/agent/cfg"
	constanct "github.com/aaronchen2k/tester/internal/agent/libs/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_fileUtils "github.com/aaronchen2k/tester/internal/pkg/libs/file"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	_shellUtils "github.com/aaronchen2k/tester/internal/pkg/libs/shell"
	_stringUtils "github.com/aaronchen2k/tester/internal/pkg/libs/string"
	"github.com/go-errors/errors"
	"strconv"
	"strings"
)

func Create(vm *_domain.Vm) (err error) {

	// create def file
	defContent, macAddress := genDefFromTempl(vm.Name, vm.MemorySize*1000, vm.CdromSys, vm.CdromDriver)
	if defContent == "" {
		return
	}

	vm.MacAddress = macAddress
	defPath := fmt.Sprintf("%s%s%s.xml", agentConf.Inst.WorkDir, constanct.FolderDef, vm.Name)
	_fileUtils.WriteFile(defPath, defContent)

	// create image file
	var cmd string
	rawPath := fmt.Sprintf("%s%s.qcow2", constanct.FolderImage, vm.Name)
	if vm.BackingImagePath == "" {
		cmd = fmt.Sprintf("qemu-img create -f qcow2 %s %dG", rawPath, vm.DiskSize/1000)
	} else {
		cmd = fmt.Sprintf("qemu-img create -f qcow2 -o cluster_size=2M,backing_file=%s %s %dG",
			vm.BackingImagePath, rawPath, vm.DiskSize/1000)
	}
	_, err = _shellUtils.ExeShellInDir(cmd, agentConf.Inst.WorkDir)
	if err != nil {
		_logUtils.Errorf("fail to generate vm, cmd %s, err %s.", cmd, err.Error())
		return
	}

	cmd = fmt.Sprintf("virsh define %s.xml", constanct.FolderDef+vm.Name)
	_, err = _shellUtils.ExeShellInDir(cmd, agentConf.Inst.WorkDir)
	if err != nil {
		_logUtils.Errorf("fail to define vm, cmd %s, err %s.", cmd, err.Error())
		return
	}

	cmd = fmt.Sprintf("virsh start %s", defPath)
	_, err = _shellUtils.ExeShellInDir(cmd, agentConf.Inst.WorkDir)
	if err != nil {
		_logUtils.Errorf("fail to start vm, cmd %s, err %s.", cmd, err.Error())
		return
	}

	return
}

func Define(vmUniqueName string) (err error) {
	defPath := constanct.FolderDef + vmUniqueName
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
	vmTemplate := fmt.Sprintf("%s%s/common.xml", agentConf.Inst.WorkDir, constanct.FolderTempl)

	templContent = _fileUtils.ReadFile(agentConf.Inst.WorkDir + constanct.FolderTempl + vmTemplate)

	templContent = strings.Replace(templContent, "${name}", vmName, -1)
	templContent = strings.Replace(templContent, "${memory}", strconv.Itoa(memory), -1)

	rawPath := fmt.Sprintf("%s%s%s.qcow2", agentConf.Inst.WorkDir, constanct.FolderImage, vmName)
	templContent = strings.Replace(templContent, "${rawPath}", rawPath, -1)

	cdromPath := fmt.Sprintf("%s%s%s", agentConf.Inst.WorkDir, constanct.FolderIso, vmCdrom)
	templContent = strings.Replace(templContent, "${cdrom}", cdromPath, -1)

	cdromPath2 := fmt.Sprintf("%s%s%s", agentConf.Inst.WorkDir, constanct.FolderIso, vmCdrom2)
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
	defPath := fmt.Sprintf("%s/%s.xml", constanct.FolderDef, vmUniqueName)
	cmd := fmt.Sprintf("rm -rf %s", defPath)
	_, err = _shellUtils.ExeShellInDir(cmd, agentConf.Inst.WorkDir)
	if err != nil {
		_logUtils.Errorf("fail to remove define file, cmd %s, err %s.", cmd, err.Error())
		return
	}
	return
}
