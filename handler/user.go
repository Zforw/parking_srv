package handler

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"parking/global"
	"parking/model"
)

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		// Limit  指定要查询的最大记录数
		// Offset 指定开始返回记录前要跳过的记录数
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func CreateUser(openid string) error {
	zap.S().Info("创建用户")
	user := model.User{OpenId: openid}
	if result := global.DB.Where("open_id=?", openid).First(&user); result.RowsAffected != 0 {
		return errors.New("用户已存在")
	}
	res := global.DB.Create(&user)
	return res.Error
}

func GetUserList(pn, psize int) ([]model.UserResp, int) {
	zap.S().Info("用户列表")
	var users []model.User
	result := Paginate(pn, psize)(global.DB).Find(&users)
	var data []model.UserResp
	for _, v := range users {
		data = append(data, model.UserResp{OpenId: v.OpenId})
	}
	count := int(result.RowsAffected)
	return data, count
}
