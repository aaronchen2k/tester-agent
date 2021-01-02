package agentUtils

import (
	"fmt"
	_commonUtils "github.com/aaronchen2k/tester/internal/pkg/libs/common"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	serverRes "github.com/aaronchen2k/tester/res/server"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

var (
	usageFile = fmt.Sprintf("res%sdoc%susage.txt", string(os.PathSeparator), string(os.PathSeparator))
)

func PrintUsage() {
	_logUtils.PrintToWithColor("Usage: ", color.FgCyan)

	usage := ReadResData(usageFile)

	app := "tester-agent"
	if _commonUtils.IsWin() {
		app += ".exe"
	}
	usage = fmt.Sprintf(usage, app)
	fmt.Printf("%s\n", usage)
}

func ReadResData(path string) string {
	isRelease := _commonUtils.IsRelease()

	var jsonStr string
	if isRelease {
		data, _ := serverRes.Asset(path)
		jsonStr = string(data)
	} else {
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			jsonStr = "fail to read " + path
		} else {
			str := string(buf)
			jsonStr = _commonUtils.RemoveBlankLine(str)
		}
	}

	return jsonStr
}
