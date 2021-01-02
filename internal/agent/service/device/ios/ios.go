package iosService

import (
	"fmt"
	deviceCommon "github.com/aaronchen2k/tester/internal/agent/service/device/common"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	_shellUtils "github.com/aaronchen2k/tester/internal/pkg/libs/shell"
	"strings"
)

var (
	Devices = make([]_domain.DeviceInst, 0)
)

func GetDeviceInsts() []_domain.DeviceInst {
	specs := GetDeviceSpecs()
	Devices := deviceCommon.SpecToDevInsts(specs)
	return Devices
}

func GetDeviceSpecs() []_domain.DeviceSpec {
	specs := make([]_domain.DeviceSpec, 0)

	// get device list
	serials := make([]string, 0)
	cmd := "idevice_id -l"
	out, err := _shellUtils.ExeShell(cmd)
	//a0c865fd
	lines := strings.Split(out, "\n")
	for _, line := range lines {
		serial := strings.TrimSpace(line)
		if serial == "" {
			continue
		}

		serials = append(serials, serial)
	}

	for _, serial := range serials {
		spec := _domain.DeviceSpec{Serial: serial}

		out, err = exeIdeviceInfoCmd(serial)
		if err != nil {
			continue
		}

		retrieveDevInfoFromCmdOutput(out, &spec)
		specs = append(specs, spec)
	}

	return specs
}

func exeIdeviceInfoCmd(serial string) (string, error) {
	cmd := fmt.Sprintf("ideviceinfo -u %s", serial)

	out, err := _shellUtils.ExeShell(cmd)
	if err != nil {
		_logUtils.Error("fail to exec ideviceinfo command: " + cmd + ", err: " + err.Error())
	}
	return out, err
}

func retrieveDevInfoFromCmdOutput(out string, spec *_domain.DeviceSpec) {
	lines := strings.Split(out, "\n")
	// DeviceName: iPhone
	// ProductVersion: 13.5

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.Index(line, "DeviceName") > -1 {
			arr := strings.Split(line, ":")

			str := strings.TrimSpace(arr[1])
			spec.Version = strings.TrimSpace(str)

		} else if strings.Index(line, "ProductVersion") > -1 { // uppercase
			arr := strings.Split(line, ":")

			str := strings.TrimSpace(arr[1])
			spec.Os = "IOS_" + strings.TrimSpace(str)

		}
	}
}
