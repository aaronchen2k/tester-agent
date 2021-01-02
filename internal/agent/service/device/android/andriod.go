package androidService

import (
	"fmt"
	deviceCommon "github.com/aaronchen2k/tester/internal/agent/service/device/common"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	_shellUtils "github.com/aaronchen2k/tester/internal/pkg/libs/shell"
	"regexp"
	"strconv"
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
	cmd := "adb devices | grep  -v devices"
	out, err := _shellUtils.ExeShell(cmd)
	//a0c865fd	device
	lines := strings.Split(out, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		arr := strings.Split(line, "\t")
		serial := arr[0]

		serials = append(serials, serial)
	}

	for _, serial := range serials {
		spec := _domain.DeviceSpec{Serial: serial, Os: "android"}

		out, err = exeAdbInfoCmd(serial)
		if err != nil {
			continue
		}

		retrieveDevInfoFromCmdOutput(out, &spec)
		specs = append(specs, spec)
	}

	return specs
}

func exeAdbInfoCmd(serial string) (string, error) {
	cmd := fmt.Sprintf("adb -s %s shell \""+
		"getprop ro.product.model & "+
		"getprop ro.build.version.release & "+
		"ifconfig wlan0 | grep Mask & "+
		"cat /proc/cpuinfo | grep -E 'Hardware|Processor' & "+
		"cat /proc/meminfo | grep MemTotal & "+
		"dumpsys battery | grep level & "+
		"wm size & "+
		"wm density & "+
		"\"", serial)

	out, err := _shellUtils.ExeShell(cmd)
	if err != nil {
		_logUtils.Error("fail to exec adb command: " + cmd + ", err: " + err.Error())
	}
	return out, err
}

func retrieveDevInfoFromCmdOutput(out string, spec *_domain.DeviceSpec) {
	lines := strings.Split(out, "\n")
	// Redmi Note 8
	// 9
	// inet addr:172.16.0.2  Bcast:172.16.0.255  Mask:255.255.255.0
	// Processor	: AArch64 Processor rev 4 (aarch64)
	// Hardware	: Qualcomm Technologies, Inc TRINKET
	// MemTotal:        5812180 kB
	// level: 100
	// Physical size: 1080x2340
	// Physical density: 440
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.Index(line, "Hardware") > -1 {
			arr := strings.Split(line, ":")

			str := strings.TrimSpace(arr[1])
			if spec.Cpu == "" {
				spec.Cpu = str
			} else {
				spec.Cpu = str + " " + spec.Cpu
			}

		} else if strings.Index(line, "Processor") > -1 { // uppercase
			arr := strings.Split(line, ":")

			str := strings.TrimSpace(arr[1])
			if spec.Cpu == "" {
				spec.Cpu = str
			} else {
				spec.Cpu = spec.Cpu + " " + str
			}

		} else if strings.Index(line, "MemTotal") > -1 {
			arr := strings.Split(line, ":")
			arr = strings.Split(strings.TrimSpace(arr[1]), " ")

			str := strings.TrimSpace(arr[0])
			spec.Ram, _ = strconv.Atoi(str)

		} else if strings.Index(line, "Mask") > -1 {
			re := regexp.MustCompile(`([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+)`)
			match := re.FindString(line)

			spec.DeviceIp = match
		} else if strings.Index(line, "level") > -1 {
			arr := strings.Split(line, ":")
			spec.Battery, _ = strconv.Atoi(arr[1])

		} else if strings.Index(line, "size") > -1 {
			arr := strings.Split(line, ":")
			arr = strings.Split(arr[1], "x")

			spec.ResolutionWidth, _ = strconv.Atoi(strings.TrimSpace(arr[0]))
			spec.ResolutionHeight, _ = strconv.Atoi(strings.TrimSpace(arr[1]))

		} else if strings.Index(line, "density") > -1 {
			arr := strings.Split(line, ":")
			spec.Density, _ = strconv.Atoi(strings.TrimSpace(arr[1]))

		} else {
			re := regexp.MustCompile(`^[0-9\.]+$`)
			temp := re.FindString(line)

			if temp != "" {
				var err error
				spec.ApiLevel, err = strconv.Atoi(line)
				if err != nil {
					spec.Version = line
				}
			} else {
				spec.Model = line
			}
		}
	}
}
