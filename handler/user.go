package handler

import (
	"parking/global"
	"parking/model"
)

func CreateUser(openid string) error {
	user := model.User{OpenId: openid}
	global.DB.Create(&user)
	return nil
}
