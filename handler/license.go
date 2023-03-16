package handler

import (
	"errors"
	"parking/global"
	"parking/model"
)

func CreateLicense(number string, openid string) error {
	user := model.User{
		OpenId: openid,
	}
	if result := global.DB.First(&user); result.RowsAffected == 0 {
		return errors.New("用户不存在")
	}
	license := model.License{
		Number: number,
		UserID: user.ID,
		User:   user,
		Status: "OUT",
	}
	global.DB.Create(&license)
	return nil
}
