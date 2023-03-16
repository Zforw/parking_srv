package handler

import (
	"parking/global"
	"parking/model"
)

func CreateLicense(license *model.License) error {
	global.DB.Create(&license)
	return nil
}
