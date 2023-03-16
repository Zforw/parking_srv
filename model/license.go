package model

// License 车牌
type License struct {
	BaseModel
	Number string `gorm:"type:varchar(20)"` //车牌号
	UserID int32
	User   User   `gorm:"foreignKey:UserID"`
	Status string `gorm:"type:varchar(20)  comment 'IN(进入), OUT(离开)'"`
}

func (License) TableName() string {
	return "license"
}
