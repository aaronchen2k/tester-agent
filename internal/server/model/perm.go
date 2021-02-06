package model

type Permission struct {
	BaseModel

	Name        string `gorm:"name;not null;type:varchar(256)" json:"name" validate:"required,gte=4,lte=50" comment:"名称"`
	DisplayName string `gorm:"displayName;type:varchar(256)" json:"displayName" comment:"显示名称"`
	Description string `gorm:"description;type:varchar(256)" json:"description" comment:"描述"`
	Act         string `gorm:"act;type:varchar(256)" json:"act" comment:"Act"`
}
