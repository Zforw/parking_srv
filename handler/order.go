package handler

import (
	"errors"
	"fmt"
	"math"
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

func Status2Chn(status string) string {
	switch status {
	case "WAIT_BUYER_PAY":
		return "待支付"
	case "PAYING":
		return "支付中"
	case "TRADE_SUCCESS":
		return "交易成功"
	}
	return "未知"
}

func PayType2Chn(status string) string {
	switch status {
	case "cash":
		return "现金"
	case "alipay":
		return "支付宝"
	case "wechat":
		return "微信支付"
	}
	return "未知"
}

func CreateOrder(number string, start time.Time) error {
	l := model.License{}
	if result := global.DB.Where("number=?", number).First(&l); result.RowsAffected == 0 {
		return errors.New("车牌不存在")
	}
	o0 := model.OrderInfo{}
	if result := global.DB.Where("license_id=? AND status <> ?", l.ID, "TRADE_SUCCESS").First(&o0); result.RowsAffected != 0 {
		if o0.Status != "TRADE_SUCCESS" {
			return errors.New("订单已存在")
		}
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
	l.Status = "IN"
	res := global.DB.Save(&l)
	res = global.DB.Create(&o)
	return res.Error
}

func GetOrderList(pn, psize int) ([]model.OrderResp, int, error) {
	var oo []model.OrderInfo
	var data []model.OrderResp
	result := global.DB.Preload("License").Scopes(Paginate(pn, psize)).Find(&oo)
	for _, v := range oo {
		data = append(data, model.OrderResp{
			OrderSn:       v.OrderSn,
			PayType:       PayType2Chn(v.PayType),
			Status:        Status2Chn(v.Status),
			OrderMount:    v.OrderMount,
			StartTime:     v.StartTime,
			PayTime:       v.PayTime,
			LicenseNumber: v.License.Number,
		})
	}
	count := int(result.RowsAffected)
	return data, count, result.Error
}

func GetUserOrderList(id string) ([]model.OrderResp, int, error) {
	var oo []model.OrderInfo
	u := model.User{}
	if result := global.DB.Where("open_id=?", id).First(&u); result.RowsAffected == 0 {
		return nil, 0, errors.New("用户不存在")
	}
	var data []model.OrderResp
	localDB := global.DB
	localDB = localDB.Where("user_id=?", u.ID)
	result := localDB.Preload("License").Scopes(Paginate(0, 90)).Find(&oo)
	for _, v := range oo {
		data = append(data, model.OrderResp{
			OrderSn:       v.OrderSn,
			PayType:       PayType2Chn(v.PayType),
			Status:        Status2Chn(v.Status),
			OrderMount:    v.OrderMount,
			StartTime:     v.StartTime,
			PayTime:       v.PayTime,
			LicenseNumber: v.License.Number,
		})
	}
	count := int(result.RowsAffected)
	return data, count, result.Error
}

func GetLicenseOrderList(number string) ([]model.OrderResp, int, error) {
	var oo []model.OrderInfo
	l := model.License{}
	if result := global.DB.Where("number=?", number).First(&l); result.RowsAffected == 0 {
		return nil, 0, errors.New("车牌不存在")
	}
	var data []model.OrderResp
	localDB := global.DB
	localDB = localDB.Where("license_id=?", l.ID)
	result := localDB.Preload("License").Scopes(Paginate(0, 90)).Find(&oo)
	for _, v := range oo {
		data = append(data, model.OrderResp{
			OrderSn:       v.OrderSn,
			PayType:       PayType2Chn(v.PayType),
			Status:        Status2Chn(v.Status),
			OrderMount:    v.OrderMount,
			StartTime:     v.StartTime,
			PayTime:       v.PayTime,
			LicenseNumber: v.License.Number,
		})
	}
	count := int(result.RowsAffected)
	return data, count, result.Error
}

// UpdateOrder 查找最新的待支付订单
func UpdateOrder(number, pay_type string) error {
	l := model.License{}
	if result := global.DB.Where("number=?", number).First(&l); result.RowsAffected == 0 {
		return errors.New("车牌不存在")
	}
	o := model.OrderInfo{}
	if result := global.DB.Where("license_id=? AND status=?", l.ID, "PAYING").First(&o); result.RowsAffected == 0 {
		return errors.New("订单不存在")
	}
	o.PayType = pay_type
	o.Status = "TRADE_SUCCESS"
	res := global.DB.Save(&o)
	return res.Error
}

func CalcMoney(number string, end time.Time) (float64, error) {
	o := model.OrderInfo{}
	l := model.License{}
	if result := global.DB.Where("number=?", number).First(&l); result.RowsAffected == 0 {
		return 0, errors.New("车牌不存在")
	}
	if result := global.DB.Where("license_id=? AND status=?", l.ID, "WAIT_BUYER_PAY").First(&o); result.RowsAffected == 0 {
		return 0, errors.New("订单不存在")
	}
	o.PayTime = &end
	o.Status = "PAYING"
	dur := time.Since(*o.StartTime)
	ch, err := GetCharge()
	if err != nil {
		return 0, err
	}
	var money float32
	if dur <= time.Hour {
		money = float32(ch.A)
	} else if dur.Hours() > 24 {
		days := int(math.Floor(dur.Hours() / 24))
		remaining := dur - time.Duration(days)*time.Hour*24
		if remaining <= time.Hour {
			money = float32(days*ch.C + ch.A)
		} else {
			m := int(math.Ceil(float64((remaining-time.Hour)/time.Hour)))*ch.B + ch.A
			if m >= ch.C {
				money = float32(ch.C * (days + 1))
			} else {
				money = float32(m + ch.C*days)
			}
		}
	} else {
		m := int(math.Ceil(float64((dur-time.Hour)/time.Hour)))*ch.B + ch.A
		if m >= ch.C {
			money = float32(ch.C)
		} else {
			money = float32(m)
		}
	}
	o.OrderMount = money
	res := global.DB.Save(&o)
	return float64(money), res.Error
}

func SetCharge(a, b, c int) error {
	ch := model.Charge{}
	ch.ID = 1
	if result := global.DB.First(&ch); result.RowsAffected == 0 {
		ch.A = int32(a)
		ch.B = int32(b)
		ch.C = int32(c)
		res := global.DB.Create(&ch)
		return res.Error
	}
	ch.A = int32(a)
	ch.B = int32(b)
	ch.C = int32(c)
	res := global.DB.Save(&ch)
	return res.Error
}

func GetCharge() (model.ChargeResp, error) {
	ch := model.Charge{}
	ch.ID = 1
	if result := global.DB.First(&ch); result.RowsAffected == 0 {
		return model.ChargeResp{}, errors.New("数据不存在")
	}
	return model.ChargeResp{
		A: int(ch.A),
		B: int(ch.B),
		C: int(ch.C),
	}, nil
}
