package model

// License 车牌
type License struct {
	BaseModel
	Number string `gorm:"type:varchar(20)"` //车牌号
	User   User   //用户对象
	Status string `gorm:"type:varchar(20)  comment 'IN(进入), OUT(离开)'"`
}

func (License) TableName() string {
	return "license"
}
