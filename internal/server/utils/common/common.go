package serverUtils

import (
	"fmt"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
)

func GenVmHostName(name string, osPlatform _const.OsPlatform, osName _const.OsType, osLang _const.SysLang) (ret string) {
	ret = fmt.Sprintf("%s-%s-%s-%s", name, osPlatform, osName, osLang)

	return
}
