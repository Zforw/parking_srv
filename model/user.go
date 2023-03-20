package model

// User 用户
type User struct {
	BaseModel
	Auth   int    `gorm:"comment '0(普通用户), 1(收费人员), 2(管理人员)'"` //用户级别, 暂时将1、2合并
	OpenId string `gorm:"type:varchar(40)"`                    //用户id
	Pass   string `gorm:"type:varchar(100)"`                   //管理员密码
}

type UserResp struct {
	OpenId string `json:"id"`   //用户的OpenId
	Auth   string `json:"auth"` //用户级别
}

func (User) TableName() string {
	return "user"
}
