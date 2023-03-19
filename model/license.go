package model

// License 车牌
type License struct {
	BaseModel
	Number string `gorm:"type:varchar(20)"` //车牌号
	UserID int32
	User   User   `gorm:"foreignKey:UserID"`
	Status string `gorm:"type:varchar(20)  comment 'IN(进入), OUT(离开)'"`
}

type LicenseResp struct {
	Number string `json:"number"`
	OpenId string `json:"id"`
	Status string `json:"status"`
}

type UserLicenseResp struct {
	Number string `json:"number"`
	Status string `json:"status"`
}

func (License) TableName() string {
	return "license"
}
