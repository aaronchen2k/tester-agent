package serverUtils

import (
	"fmt"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
)

func GenVmHostName(queueId uint, osPlatform _const.OsPlatform, osName _const.OsType, osLang _const.SysLang) (ret string) {
	ret = fmt.Sprintf("queue%d-%s-%s-%s", queueId, osPlatform, osName, osLang)

	return
}
