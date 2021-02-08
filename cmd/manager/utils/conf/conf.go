package managerConf

import (
	managerConst "github.com/aaronchen2k/tester/cmd/manager/utils/const"
	managerVari "github.com/aaronchen2k/tester/cmd/manager/utils/vari"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_commonUtils "github.com/aaronchen2k/tester/internal/pkg/libs/common"
	_fileUtils "github.com/aaronchen2k/tester/internal/pkg/libs/file"
	_i118Utils "github.com/aaronchen2k/tester/internal/pkg/libs/i118"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
	"path/filepath"
)

var (
	Inst Config
)

type Config struct {
	Interval int64  `yaml:"interval"`
	Language string `yaml:"language"`

	Clients []Client `yaml:"clients,flow"`
}
type Client struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Params  string `yaml:"params"`
}

func NewConfig() Config {
	return Config{
		Interval: 6,
		Language: "en",
	}
}

func Init() {
	managerVari.WorkDir = path.Join(_commonUtils.GetUserHome(), _const.AppName)
	// just for ide debug
	if !_commonUtils.IsRelease() {
		managerVari.WorkDir = _fileUtils.AddSepIfNeeded(path.Join("cmd", "manager"))
	}

	managerVari.ConfFile = managerVari.WorkDir + "conf.yml"
	managerVari.LogFile = managerVari.WorkDir + managerConst.AppName + ".log"

	initInst()

	_i118Utils.InitI118(Inst.Language, managerConst.AppName)
	PrintCurrConfig()
}

func SaveConfig(conf Config) error {
	_fileUtils.MkDirIfNeeded(filepath.Dir(managerVari.ConfFile))

	cfg := ini.Empty()
	cfg.ReflectFrom(&conf)

	cfg.SaveTo(managerVari.ConfFile)

	Inst = ReadCurrConfig()
	return nil
}

func PrintCurrConfig() {
	log.Println("\n" + _i118Utils.I118Prt.Sprintf("current_config"))
	log.Printf("%#v \n", Inst)
}

func ReadCurrConfig() Config {
	config := Config{}

	if !_fileUtils.FileExist(managerVari.ConfFile) {
		config := NewConfig()
		_i118Utils.InitI118(config.Language, "")

		return config
	}

	ini.MapTo(&config, managerVari.ConfFile)

	return config
}

func initInst() {
	_fileUtils.MkDirIfNeeded(managerVari.WorkDir)
	content := _fileUtils.ReadFileBuf(managerVari.ConfFile)

	err := yaml.Unmarshal(content, &Inst)
	if err != nil {
		_logUtils.Error(_i118Utils.I118Prt.Sprintf("fail_to_read_file", managerVari.ConfFile))
		os.Exit(1)
	}
}
