package handler

import (
	"parking/global"
	"parking/model"
)

func CreateUser(user *model.User) error {
	global.DB.Create(&user)
	return nil
}
