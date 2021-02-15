package model

import (
	"time"
)

type BaseModel struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

var (
	Models = []interface{}{
		&User{},
		&Role{},
		&Permission{},

		&OsPlatform{},
		&OsType{},
		&OsLang{},
		&BrowserType{},
		&Device{},

		&Iso{},
		&Queue{},

		&Plan{},
		&Task{},
		&Queue{},
		&Build{},

		&Cluster{},
		&Computer{},
		&ContainerImage{},
		&Container{},
		&VmTempl{},
		&Vm{},
	}
)
