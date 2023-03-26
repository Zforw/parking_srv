package handler

import (
	"errors"
	"fmt"
	"math/rand"
	"parking/global"
	"parking/model"
	"time"
)

func GenerateOrderSn(userId int32) string {
	//订单号的生成规则
	/*
		年月日时分秒+用户id+2位随机数
	*/
	now := time.Now()
	rand.Seed(time.Now().UnixNano())
	orderSn := fmt.Sprintf("%d%d%d%d%d%d%d%d",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Nanosecond(),
		userId, rand.Intn(90)+10,
	)
	return orderSn
}

func CreateOrder(number string, start time.Time) error {
	l := model.License{}
	if result := global.DB.Where("number=?", number).First(&l); result.RowsAffected == 0 {
		return errors.New("车牌不存在")
	}
	o := model.OrderInfo{
		UserID:    l.UserID,
		User:      l.User,
		OrderSn:   GenerateOrderSn(l.UserID),
		Status:    "WAIT_BUYER_PAY",
		StartTime: &start,
		LicenseID: l.ID,
		License:   l,
	}
	res := global.DB.Create(&o)
	return res.Error
}

func UpdateOrder(number, pay_type string) error {
	l := model.License{}
	if result := global.DB.Where("number=?", number).First(&l); result.RowsAffected == 0 {
		return errors.New("车牌不存在")
	}
	o := model.OrderInfo{
		UserID:    l.UserID,
		User:      l.User,
		OrderSn:   GenerateOrderSn(l.UserID),
		Status:    "WAIT_BUYER_PAY",
		LicenseID: l.ID,
		License:   l,
	}
	res := global.DB.Create(&o)
	return res.Error
}

func GetMoney(number string, end time.Time) (float64, error) {
	o := model.OrderInfo{}
	l := model.License{}
	if result := global.DB.Where("number=?", number).First(&l); result.RowsAffected == 0 {
		return 0, errors.New("车牌不存在")
	}
	if result := global.DB.Where("license_id=?", l.ID).First(&o); result.RowsAffected == 0 {
		return 0, errors.New("订单不存在")
	}
	o.PayTime = &end
	dur := time.Since(*o.StartTime)
	res := global.DB.Save(&o)
	return dur.Hours() * 4, res.Error
}
