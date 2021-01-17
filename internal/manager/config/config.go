package managerConfig

import (
	"fmt"
	managerModel "github.com/aaronchen2k/tester/internal/manager/model"
	managerVari "github.com/aaronchen2k/tester/internal/manager/utils/vari"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_commonUtils "github.com/aaronchen2k/tester/internal/pkg/libs/common"
	_fileUtils "github.com/aaronchen2k/tester/internal/pkg/libs/file"
	_i118Utils "github.com/aaronchen2k/tester/internal/pkg/libs/i118"
	_stdinUtils "github.com/aaronchen2k/tester/internal/pkg/libs/stdin"
	_vari "github.com/aaronchen2k/tester/internal/pkg/vari"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

func Init() {
	_vari.WorkDir = _commonUtils.GetUserHome() + _const.PthSep + _const.AppManager + _const.PthSep
	managerVari.ConfigFile = _vari.WorkDir + "conf.ini"
	managerVari.LogFile = _vari.WorkDir + _const.AppManager + ".log"

	CheckConfigPermission()
	managerVari.Config = getInst()

	_i118Utils.InitI118(managerVari.Config.Language)
}

func SaveConfig(conf managerModel.Config) error {
	_fileUtils.MkDirIfNeeded(filepath.Dir(managerVari.ConfigFile))

	cfg := ini.Empty()
	cfg.ReflectFrom(&conf)

	cfg.SaveTo(managerVari.ConfigFile)

	managerVari.Config = ReadCurrConfig()
	return nil
}

func PrintCurrConfig() {
	log.Println("\n" + _i118Utils.I118Prt.Sprintf("current_config"))

	val := reflect.ValueOf(managerVari.Config)
	typeOfS := val.Type()
	for i := 0; i < reflect.ValueOf(managerVari.Config).NumField(); i++ {
		if !_commonUtils.IsWin() && i > 4 {
			break
		}

		val := val.Field(i)
		name := typeOfS.Field(i).Name

		log.Printf("  %s: %v \n", name, val.Interface())
	}
}

func ReadCurrConfig() managerModel.Config {
	config := managerModel.Config{}

	if !_fileUtils.FileExist(managerVari.ConfigFile) {
		config := managerModel.NewConfig()
		_i118Utils.InitI118(config.Language)

		return config
	}

	ini.MapTo(&config, managerVari.ConfigFile)

	return config
}

func getInst() managerModel.Config {
	CheckConfigReady()

	ini.MapTo(&managerVari.Config, managerVari.ConfigFile)

	return managerVari.Config
}

func CheckConfigPermission() {
	d := filepath.Dir(managerVari.LogFile)
	err := _fileUtils.MkDirIfNeeded(d)
	if err != nil {
		log.Println(fmt.Sprintf("Permission denied, please change the dir %s.", d))
		os.Exit(0)
	}
}

func CheckConfigReady() {
	if !_fileUtils.FileExist(managerVari.ConfigFile) {
		InputForSet()
	}
}

func InputForSet() {
	conf := ReadCurrConfig()

	//logUtils.PrintToWithColor(i118Utils.I118Prt.Sprintf("begin_config"), color.FgCyan)

	enCheck := ""
	var numb string
	if conf.Language == "zh" {
		enCheck = "*"
		numb = "1"
	}
	zhCheck := ""
	if conf.Language == "en" {
		zhCheck = "*"
		numb = "2"
	}

	// set lang
	langNo := _stdinUtils.GetInput("(1|2)", numb, "enter_language", enCheck, zhCheck)
	if langNo == "1" {
		conf.Language = "zh"
	} else {
		conf.Language = "en"
	}

	SaveConfig(conf)
	PrintCurrConfig()
}
