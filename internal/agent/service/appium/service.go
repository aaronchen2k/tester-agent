package appiumService

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/openstc/internal/agent/cfg"
	deviceService "github.com/aaronchen2k/openstc/internal/agent/service/device"
	_domain "github.com/aaronchen2k/openstc/internal/pkg/domain"
	_const "github.com/aaronchen2k/openstc/internal/pkg/libs/const"
	_httpUtils "github.com/aaronchen2k/openstc/internal/pkg/libs/http"
	_logUtils "github.com/aaronchen2k/openstc/internal/pkg/libs/log"
	_shellUtils "github.com/aaronchen2k/openstc/internal/pkg/libs/shell"
	"strconv"
	"strings"
)

func CheckService(devices []_domain.DeviceInst) {
	serials := deviceService.GetDeviceSerials(devices)
	appiumObjMap, appiumPortMap := getServices(serials)

	for i := 0; i < len(devices); i++ {
		serial := devices[i].Serial

		isAppiumOn := false
		if _, ok := appiumObjMap[serial]; ok {
			isAppiumOn = true
		}

		if isAppiumOn {
			devices[i].AppiumPort = appiumObjMap[serial].AppiumPort
			devices[i].AppiumStatus = _const.ServiceOn
		} else {
			validPort := getValidPort(appiumPortMap)
			appium := _domain.Appium{DeviceSerial: serial, AppiumPort: validPort, NodeIp: devices[0].NodeIp}
			appiumObjMap[serial] = appium

			result := restartService(serial, validPort)
			if result {
				devices[i].AppiumStatus = _const.ServiceOn
				devices[i].AppiumPort = validPort
			} else {
				devices[i].AppiumStatus = _const.ServiceOff
			}
		}

		checkService(&devices[i])
	}
}

func getServices(serials []string) (map[string]_domain.Appium, map[int]bool) {
	appiumMap := map[string]_domain.Appium{}
	portMap := map[int]bool{}

	processMap := getAppiumProcesses(serials)

	for serial, port := range processMap {
		appium := _domain.Appium{DeviceSerial: serial, AppiumPort: port, NodeIp: agentConf.Inst.Ip}
		appiumMap[serial] = appium

		portMap[port] = true
	}

	return appiumMap, portMap
}

func restartService(serial string, port int) bool {
	_logUtils.Info("start appium")

	killCmd := fmt.Sprintf("ps -ef | grep appium | grep %s | grep -v 'grep' | awk '{print $2}' | xargs kill -9", serial)
	_, err := _shellUtils.ExeShell(killCmd)
	if err != nil {
		return false
	}

	startCmd := fmt.Sprintf("nohup appium -p %d --default-capabilities '{\"udid\":\"%s\"}' "+
		"--log-timestamp --log-level debug:error "+
		"> appium.log 2>&1 &", port, serial)

	_, err = _shellUtils.ExeShell(startCmd)
	if err != nil {
		return false
	}

	return true
}

func getAppiumProcesses(serials []string) map[string]int {
	cmd := "ps -ef | grep appium | grep default-capabilities | grep -v 'grep' | awk '{print $11,$13}'"
	output, err := _shellUtils.ExeShell(cmd) // 4723 {"udid":"a0c865fd"}

	ret := map[string]int{}
	if err != nil || output == "" {
		return ret
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		cols := strings.Split(line, " ")

		port, _ := strconv.Atoi(cols[0])

		jsonObj := map[string]string{}
		json.Unmarshal([]byte(cols[1]), &jsonObj)

		udid := jsonObj["udid"]
		ret[udid] = port
	}

	return ret
}

func getValidPort(appiumPortMap map[int]bool) int {
	newPort := 0

	for i := 1; i < 100; i++ {
		port := 50000 + i

		if _, ok := appiumPortMap[port]; !ok {
			newPort = port
			break
		}
	}

	appiumPortMap[newPort] = true
	return newPort
}

func checkService(dev *_domain.DeviceInst) bool {
	if dev.AppiumStatus != _const.ServiceOn {
		return false
	}

	url := fmt.Sprintf("http://%s:%d/wd/hub/status", dev.NodeIp, dev.AppiumPort)
	result, ok := _httpUtils.GetObj(url, "appium")

	if ok && result != nil {
		resp := result.(map[string]interface{})

		status := resp["status"]
		if status.(float64) == 0 {
			dev.DeviceStatus = _const.DeviceActive
			dev.AppiumStatus = _const.ServiceActive

			return true
		}
	}

	return false
}
