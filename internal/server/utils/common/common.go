package serverUtils

import (
	"fmt"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_stringUtils "github.com/aaronchen2k/tester/internal/pkg/libs/string"
)

func GenVmHostName(osPlatform _const.OsPlatform, osName _const.OsType, osLang _const.SysLang) (name string) {
	uuid := _stringUtils.NewUUID()
	name = fmt.Sprintf("%s-%s-%s-%s", osPlatform, osName, osLang, uuid)

	return
}
