package handler

import (
	"errors"
	"go.uber.org/zap"
	"parking/global"
	"parking/model"
)

func CreateLicense(number string, openid string) error {
	user := model.User{
		OpenId: openid,
	}
	license := model.License{
		Number: number,
	}
	if result := global.DB.First(&user); result.RowsAffected == 0 {
		return errors.New("用户不存在")
	}
	if result := global.DB.Where("number=?", number).First(&license); result.RowsAffected != 0 {
		return errors.New("车牌已存在")
	}
	license.UserID = user.ID
	license.User = user
	license.Status = "OUT"
	res := global.DB.Create(&license)
	return res.Error
}

func GetLicenseList(pn, psize int) ([]model.LicenseResp, int) {
	zap.S().Info("车牌列表")
	var licenses []model.License
	result := global.DB.Preload("User").Scopes(Paginate(pn, psize)).Find(&licenses)
	var data []model.LicenseResp
	for _, v := range licenses {
		data = append(data, model.LicenseResp{
			Number: v.Number,
			OpenId: v.User.OpenId,
			Status: v.Status,
		})
	}
	count := int(result.RowsAffected)
	return data, count
}
