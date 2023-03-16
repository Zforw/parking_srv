package handler

import (
	"gorm.io/gorm"
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

func GetUserList() ([]model.User, error) {

	return nil, nil
}
