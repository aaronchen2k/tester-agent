package managerConfig

import (
	"fmt"
	"github.com/easysoft/zmanager/pkg/model"
	commonUtils "github.com/easysoft/zmanager/pkg/utils/common"
	constant "github.com/easysoft/zmanager/pkg/utils/const"
	fileUtils "github.com/easysoft/zmanager/pkg/utils/file"
	i118Utils "github.com/easysoft/zmanager/pkg/utils/i118"
	stdinUtils "github.com/easysoft/zmanager/pkg/utils/stdin"
	"github.com/easysoft/zmanager/pkg/utils/vari"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

func Init() {
	vari.WorkDir = commonUtils.GetUserHome() + constant.PthSep + constant.AppName + constant.PthSep
	vari.ConfigFile = vari.WorkDir + "conf.ini"
	vari.LogFile = vari.WorkDir + constant.AppName + ".log"

	CheckConfigPermission()
	vari.Config = getInst()

	i118Utils.InitI118(vari.Config.Language)
}

func SaveConfig(conf model.Config) error {
	fileUtils.MkDirIfNeeded(filepath.Dir(vari.ConfigFile))

	cfg := ini.Empty()
	cfg.ReflectFrom(&conf)

	cfg.SaveTo(vari.ConfigFile)

	vari.Config = ReadCurrConfig()
	return nil
}

func PrintCurrConfig() {
	log.Println("\n" + i118Utils.I118Prt.Sprintf("current_config"))

	val := reflect.ValueOf(vari.Config)
	typeOfS := val.Type()
	for i := 0; i < reflect.ValueOf(vari.Config).NumField(); i++ {
		if !commonUtils.IsWin() && i > 4 {
			break
		}

		val := val.Field(i)
		name := typeOfS.Field(i).Name

		fmt.Printf("  %s: %v \n", name, val.Interface())
	}
}

func ReadCurrConfig() model.Config {
	config := model.Config{}

	if !fileUtils.FileExist(vari.ConfigFile) {
		config := model.NewConfig()
		i118Utils.InitI118(config.Language)

		return config
	}

	ini.MapTo(&config, vari.ConfigFile)

	return config
}

func getInst() model.Config {
	CheckConfigReady()

	ini.MapTo(&vari.Config, vari.ConfigFile)

	return vari.Config
}

func CheckConfigPermission() {
	d := filepath.Dir(vari.LogFile)
	err := fileUtils.MkDirIfNeeded(d)
	if err != nil {
		log.Println(fmt.Sprintf("Permission denied, please change the dir %s.", d))
		os.Exit(0)
	}
}

func CheckConfigReady() {
	if !fileUtils.FileExist(vari.ConfigFile) {
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
	langNo := stdinUtils.GetInput("(1|2)", numb, "enter_language", enCheck, zhCheck)
	if langNo == "1" {
		conf.Language = "zh"
	} else {
		conf.Language = "en"
	}

	SaveConfig(conf)
	PrintCurrConfig()
}
