package handler

import (
	"errors"
	"go.uber.org/zap"
	"parking/global"
	"parking/model"
)

func LStatus2Chn(status string) string {
	if status == "IN" {
		return "入内"
	} else {
		return "离开"
	}
}

func CreateLicense(number string, openid string) error {
	user := model.User{
		OpenId: openid,
	}
	license := model.License{
		Number: number,
	}
	if result := global.DB.Where("open_id=?", openid).First(&user); result.RowsAffected == 0 {
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

func DeleteLicense(number string, openid string) error {
	u := model.User{}
	if result := global.DB.Where("open_id=?", openid).First(&u); result.RowsAffected == 0 {
		return errors.New("用户不存在")
	}
	if result := global.DB.Where("number=? and user_id=?", number, u.ID).Delete(&model.License{}); result.RowsAffected == 0 {
		return errors.New("车牌不存在")
	}
	return nil
}

/*
func UpdateLicense(number string, status string) error {
	license := model.License{
		Number: number,
	}
	if result := global.DB.Where("number=?", number).First(&license); result.RowsAffected == 0 {
		return errors.New("车牌不存在")
	}
	license.Status = status
	res := global.DB.Save(&license)
	return res.Error
}
*/

func GetLicenseList(pn, psize int) ([]model.LicenseResp, int, error) {
	var licenses []model.License
	result := global.DB.Preload("User").Scopes(Paginate(pn, psize)).Find(&licenses)
	var data []model.LicenseResp
	for _, v := range licenses {
		data = append(data, model.LicenseResp{
			Number: v.Number,
			OpenId: v.User.OpenId,
			Status: LStatus2Chn(v.Status),
		})
	}
	count := int(result.RowsAffected)
	return data, count, result.Error
}

func GetUserLicenseList(id string, pn, psize int) ([]model.UserLicenseResp, int, error) {
	var user model.User
	if result := global.DB.Where("open_id=?", id).First(&user); result.RowsAffected == 0 {
		return nil, 0, errors.New("用户不存在")
	}
	zap.S().Debug(user)
	localDB := global.DB
	var licenses []model.License
	result := localDB.Where("user_id=?", user.ID).Scopes(Paginate(pn, psize)).Find(&licenses)
	zap.S().Debug(licenses)
	var data []model.UserLicenseResp
	for _, v := range licenses {
		data = append(data, model.UserLicenseResp{
			Number: v.Number,
			Status: LStatus2Chn(v.Status),
		})
	}
	count := int(result.RowsAffected)
	return data, count, result.Error
}
