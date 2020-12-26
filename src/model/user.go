package model

type User struct {
	BaseModel

	Name     string `gorm:"not null; type:varchar(60)" json:"name" validate:"required,gte=2,lte=50" comment:"用户名"`
	Username string `gorm:"unique;not null;type:varchar(60)" json:"username" validate:"required,gte=2,lte=50"  comment:"名称"`
	Password string `gorm:"type:varchar(100)" json:"password"  comment:"密码"`
	Intro    string `gorm:"not null; type:varchar(512)" json:"introduction" comment:"简介"`
	Avatar   string `gorm:"type:longText" json:"avatar"  comment:"头像"`
	RoleIds  []uint `gorm:"-" json:"role_ids"  validate:"required" comment:"角色"`
}

type Avatar struct {
	Avatar string `gorm:"type:longText" json:"avatar" validate:"required" comment:"头像"`
}

type Token struct {
	Token string `json:"token"`
}
