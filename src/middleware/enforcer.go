package middleware

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/openstc/src/libs/common"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v2"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

func NewEnforcer() *casbin.Enforcer {
	var err error
	var conn string
	if common.Config.DB.Adapter == "mysql" {
		conn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local",
			common.Config.DB.User, common.Config.DB.Password, common.Config.DB.Host, common.Config.DB.Port, common.Config.DB.Name)
	} else if common.Config.DB.Adapter == "postgres" {
		conn = fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable",
			common.Config.DB.User, common.Config.DB.Password, common.Config.DB.Host, common.Config.DB.Name)
	} else if common.Config.DB.Adapter == "sqlite3" {
		conn = common.DBFile()
	} else {
		logrus.Println(errors.New("not supported database adapter"))
	}

	if len(conn) == 0 {
		logrus.Println(fmt.Sprintf("数据链接不可用: %s", conn))
	}

	c, err := gormadapter.NewAdapter(common.Config.DB.Adapter, conn, true) // Your driver and data source.
	if err != nil {
		logrus.Println(fmt.Sprintf("NewAdapter 错误: %v,Path: %s", err, conn))
	}

	casbinModelPath := filepath.Join(common.GetExeDir(), "rbac_model.conf")
	enforcer, err := casbin.NewEnforcer(casbinModelPath, c)
	if err != nil {
		logrus.Println(fmt.Sprintf("NewEnforcer 错误: %v", err))
	}

	_ = enforcer.LoadPolicy()

	return enforcer
}
