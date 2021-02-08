package model

import (
	"time"
)

type BaseModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `gorm:"createdAt" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"updatedAt" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"deletedAt" sql:"index" json:"deletedAt"`
}
