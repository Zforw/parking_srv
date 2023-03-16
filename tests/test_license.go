package tests

import (
	"go.uber.org/zap"
	"parking/global"
	"parking/handler"
	"parking/model"
	"testing"
)

func TestCreateLicense(t *testing.T) {
	var user model.User
	user.OpenId = "ow6HF5Y8rSlm61hl8igL1k9nOlAI"
	if result := global.DB.First(&user); result.RowsAffected == 0 {
		zap.S().Debug("用户不存在")
		return
	}
	license := &model.License{
		Number: "鄂A-1XC25",
		UserID: 1,
		User:   user,
		Status: "OUT",
	}
	err := handler.CreateLicense(license)
	if err != nil {
		zap.S().Fatal(err)
		return
	}
}
