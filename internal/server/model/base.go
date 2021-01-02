package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BaseModel struct {
	gorm.Model

	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
