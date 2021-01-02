package middleware

import (
	"errors"
	"fmt"
	_fileUtils "github.com/aaronchen2k/tester/internal/pkg/libs/file"
	"github.com/aaronchen2k/tester/internal/server/cfg"
	"github.com/aaronchen2k/tester/internal/server/db"
	"github.com/aaronchen2k/tester/internal/server/utils"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v2"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

func NewEnforcer() *casbin.Enforcer {
	var err error
	var conn string
	if serverConf.Config.DB.Adapter == "mysql" {
		conn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local",
			serverConf.Config.DB.User, serverConf.Config.DB.Password, serverConf.Config.DB.Host, serverConf.Config.DB.Port, serverConf.Config.DB.Name)
	} else if serverConf.Config.DB.Adapter == "postgres" {
		conn = fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable",
			serverConf.Config.DB.User, serverConf.Config.DB.Password, serverConf.Config.DB.Host, serverConf.Config.DB.Name)
	} else if serverConf.Config.DB.Adapter == "sqlite3" {
		conn = db.DBFile()
	} else {
		logrus.Println(errors.New("not supported database adapter"))
	}

	if len(conn) == 0 {
		logrus.Println(fmt.Sprintf("数据链接不可用: %s", conn))
	}

	c, err := gormadapter.NewAdapter(serverConf.Config.DB.Adapter, conn, true) // Your driver and data source.
	if err != nil {
		logrus.Println(fmt.Sprintf("NewAdapter 错误: %v,Path: %s", err, conn))
	}

	exeDir := agentUtils.GetExeDir()
	pth := filepath.Join(exeDir, "rbac_model.conf")
	if !_fileUtils.FileExist(pth) { // debug mode
		pth = filepath.Join(exeDir, "cmd", "server", "rbac_model.conf")
	}

	enforcer, err := casbin.NewEnforcer(pth, c)
	if err != nil {
		logrus.Println(fmt.Sprintf("NewEnforcer 错误: %v", err))
	}

	_ = enforcer.LoadPolicy()

	return enforcer
}
