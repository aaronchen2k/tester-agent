package db

import (
	"fmt"
	_logUtils "github.com/aaronchen2k/openstc-common/src/libs/log"
	"github.com/aaronchen2k/openstc/src/libs/common"
	"github.com/aaronchen2k/openstc/src/model"
	"github.com/casbin/gorm-adapter/v2"
	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"

	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	inst *Instance
)

func GetInst() *Instance {
	if inst == nil {
		InitDB()
	}

	return inst
}

func InitDB() {
	var dialector gorm.Dialector

	if common.Config.DB.Adapter == "sqlite3" {
		conn := common.DBFile()
		dialector = sqlite.Open(conn)

	} else if common.Config.DB.Adapter == "mysql" {
		conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local",
			common.Config.DB.User, common.Config.DB.Password, common.Config.DB.Host, common.Config.DB.Port, common.Config.DB.Name)
		dialector = mysql.Open(conn)

	} else if common.Config.DB.Adapter == "postgres" {
		conn := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable",
			common.Config.DB.User, common.Config.DB.Password, common.Config.DB.Host, common.Config.DB.Name)
		dialector = postgres.Open(conn)

	} else {
		_logUtils.Info("not supported database adapter")
	}

	DB, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   common.Config.DB.Prefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: false,                   // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
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
	inst.config = &common.Config.DB
}

func (*Instance) DB() *gorm.DB {
	return inst.db
}

type Instance struct {
	config *common.DBConfig
	db     *gorm.DB
}

func (i *Instance) Close() error {
	if i.db != nil {
		sqlDB, _ := i.db.DB()
		return sqlDB.Close()
	}
	return nil
}

func (i *Instance) Migrate() {
	err := i.DB().AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Permission{},
		&gormadapter.CasbinRule{},
	)

	if err != nil {
		color.Yellow(fmt.Sprintf("初始化数据表错误 ：%+v", err))
	}
}
