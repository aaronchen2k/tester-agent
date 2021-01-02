package service

import (
	"fmt"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_const "github.com/aaronchen2k/tester/internal/pkg/libs/const"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/repo"
	"time"
)

type DeviceService struct {
	DeviceRepo *repo.DeviceRepo `inject:""`
}

func NewDeviceService() *DeviceService {
	return &DeviceService{}
}

func (s *DeviceService) Register(devices []_domain.DeviceInst) (result _domain.RpcResult) {
	for _, device := range devices {
		device.LastRegisterDate = time.Now()
		err := s.DeviceRepo.Register(device)

		if err != nil {
			result.Fail(fmt.Sprintf("fail to register device %s ", device.Serial))
			break
		}
	}

	result.Success(fmt.Sprintf("success to register %d devices", len(devices)))
	return
}

func (s *DeviceService) IsDeviceReady(device model.Device) bool {
	if device.ID == 0 {
		return false
	}

	deviceStatus := device.DeviceStatus
	appiumStatus := device.AppiumStatus

	registerExpire := time.Now().Unix()-device.LastRegisterDate.Unix() > _const.RegisterExpireTime*60*1000

	ret := deviceStatus == _const.DeviceActive && appiumStatus == _const.ServiceActive && !registerExpire
	return ret
}
