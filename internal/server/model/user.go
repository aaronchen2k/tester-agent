package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	Name             string    `gorm:"name;not null; type:varchar(60)" json:"name" validate:"required,gte=2,lte=50" comment:"用户名"`
	Username         string    `gorm:"username;unique;not null;type:varchar(60)" json:"username" validate:"required,gte=2,lte=50"  comment:"名称"`
	Password         string    `gorm:"password;type:varchar(100)" json:"password"  comment:"密码"`
	Intro            string    `gorm:"intro;not null; type:varchar(512)" json:"introduction" comment:"简介"`
	Avatar           string    `gorm:"avatar;type:longText" json:"avatar"  comment:"头像"`
	Token            string    `gorm:"token" json:"token" comment:"令牌"`
	TokenUpdatedTime time.Time `gorm:"tokenUpdatedTime" json:"tokenUpdatedTime" comment:"令牌更新时间"`

	RoleIds []uint `gorm:"-" json:"role_ids"  validate:"required" comment:"角色"`
}

type Avatar struct {
	Avatar string `gorm:"avatar;type:longText" json:"avatar" validate:"required" comment:"头像"`
}

type Token struct {
	Token      string `json:"token"`
	RememberMe bool   `json:"rememberMe"`
}
