package db

import (
	"fmt"
	_fileUtils "github.com/aaronchen2k/tester/internal/pkg/libs/file"
	"github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"github.com/aaronchen2k/tester/internal/pkg/utils"
	"github.com/aaronchen2k/tester/internal/server/cfg"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"path/filepath"
	"strings"
	"time"

	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	inst *Instance

	FlagVarDBFile string
)

func GetInst() *Instance {
	if inst == nil {
		InitDB()
	}

	return inst
}

func InitDB() {
	var dialector gorm.Dialector

	if serverConf.Config.DB.Adapter == "sqlite3" {
		conn := DBFile()
		dialector = sqlite.Open(conn)

	} else if serverConf.Config.DB.Adapter == "mysql" {
		conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local",
			serverConf.Config.DB.User, serverConf.Config.DB.Password, serverConf.Config.DB.Host, serverConf.Config.DB.Port, serverConf.Config.DB.Name)
		dialector = mysql.Open(conn)

	} else if serverConf.Config.DB.Adapter == "postgres" {
		conn := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable",
			serverConf.Config.DB.User, serverConf.Config.DB.Password, serverConf.Config.DB.Host, serverConf.Config.DB.Name)
		dialector = postgres.Open(conn)

	} else {
		_logUtils.Info("not supported database adapter")
	}

	DB, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   serverConf.Config.DB.Prefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: false,                       // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})

	if err != nil {
		_logUtils.Info(err.Error())
	}

	_ = DB.Use(
		dbresolver.Register(
			dbresolver.Config{ /* xxx */ }).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200),
	)

	DB.Session(&gorm.Session{FullSaveAssociations: true, AllowGlobalUpdate: false})

	inst = &Instance{}
	inst.db = DB
	inst.config = &serverConf.Config.DB
}

func (*Instance) DB() *gorm.DB {
	return inst.db
}

type Instance struct {
	config *serverConf.DBConfig
	db     *gorm.DB
}

func (i *Instance) Close() error {
	if i.db != nil {
		sqlDB, _ := i.db.DB()
		return sqlDB.Close()
	}
	return nil
}

func DBFile() string {
	if FlagVarDBFile != "" {
		return FlagVarDBFile
	}
	if _fileUtils.FileExist(DBFileDev()) {
		return DBFileDev()
	}

	path := filepath.Join(utils.GetExeDir(), strings.ToLower(serverConf.Config.DB.Name+".db"))
	return path
}

func DBFileDev() string {
	return filepath.Join(utils.CWD(), strings.ToLower(utils.EXEName())+".dev.db")
}
