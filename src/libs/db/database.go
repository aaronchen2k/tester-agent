package db

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/openstc/src/libs/common"
	logger "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"

	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func InitDb() {
	var err error
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
		logger.Println(errors.New("not supported database adapter"))
	}

	Db, err = gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   common.Config.DB.Prefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: false,                   // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})

	if err != nil {
		logger.Println(err)
	}

	_ = Db.Use(
		dbresolver.Register(
			dbresolver.Config{ /* xxx */ }).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200),
	)
	Db.Session(&gorm.Session{FullSaveAssociations: true, AllowGlobalUpdate: false})
}
