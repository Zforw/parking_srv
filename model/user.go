package model

// User 用户
type User struct {
	BaseModel
	OpenId string `gorm:"type:varchar(20)"` //用户的OpenId
}

type UserResp struct {
	OpenId string `json:"id"` //用户的OpenId
}

func (User) TableName() string {
	return "user"
}
