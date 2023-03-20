package handler

import (
	"errors"
	"gorm.io/gorm"
	"parking/global"
	"parking/model"
	"parking/utils"
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

func CreateUser(auth int, openid, pass string) error {
	user := model.User{OpenId: openid, Auth: auth, Pass: pass}
	if result := global.DB.Where("open_id=?", openid).First(&user); result.RowsAffected != 0 {
		return errors.New("用户已存在")
	}
	res := global.DB.Create(&user)
	return res.Error
}

func Login(openid, pass string) error {
	user := model.User{OpenId: openid}
	if result := global.DB.Where("open_id=?", openid).First(&user); result.RowsAffected == 0 {
		return errors.New("用户不存在")
	}
	if utils.VerifyPass(pass, user.Pass) {
		return nil
	}
	return errors.New("用户名或密码错误")
}

func GetUserList(pn, psize int) ([]model.UserResp, int, error) {
	var users []model.User
	result := Paginate(pn, psize)(global.DB).Find(&users)
	mauth := map[int]string{0: "普通用户", 1: "管理员"}
	var data []model.UserResp
	for _, v := range users {
		data = append(data, model.UserResp{OpenId: v.OpenId, Auth: mauth[v.Auth]})
	}
	count := int(result.RowsAffected)
	return data, count, result.Error
}
