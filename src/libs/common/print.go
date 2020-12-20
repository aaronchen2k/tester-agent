package libs

import (
	"fmt"
	"github.com/aaronchen2k/openstc-agent/res"
	commonUtils "github.com/aaronchen2k/openstc-common/src/utils/common"
	logUtils "github.com/aaronchen2k/openstc-common/src/utils/log"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

var (
	usageFile = fmt.Sprintf("res%sdoc%susage.txt", string(os.PathSeparator), string(os.PathSeparator))
)

func PrintUsage() {
	logUtils.PrintToWithColor("Usage: ", color.FgCyan)

	usage := ReadResData(usageFile)

	app := "openstc-agent"
	if commonUtils.IsWin() {
		app += ".exe"
	}
	usage = fmt.Sprintf(usage, app)
	fmt.Printf("%s\n", usage)
}

func ReadResData(path string) string {
	isRelease := commonUtils.IsRelease()

	var jsonStr string
	if isRelease {
		data, _ := res.Asset(path)
		jsonStr = string(data)
	} else {
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			jsonStr = "fail to read " + path
		} else {
			str := string(buf)
			jsonStr = commonUtils.RemoveBlankLine(str)
		}
	}

	return jsonStr
}
