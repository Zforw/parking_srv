package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"parking/form"
	"parking/global"
	"parking/model"
	"strings"
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
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(),
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

func PayTime2Chn(p *time.Time) string {
	if p == nil {
		return ""
	} else {
		return p.Format("2006-01-02 15:04:05")
	}
}

func CreateOrder(number, blockNo string, start time.Time) error {
	b := model.Block{}
	if result := global.DB.Where("block_no=?", blockNo).First(&b); result.RowsAffected == 0 {
		return errors.New("停车区不存在")
	}
	l := model.License{}
	if result := global.DB.Where("number=?", number).First(&l); result.RowsAffected == 0 {
		//return errors.New("车牌不存在")
		CreateLicense(number, "anonymous")
		global.DB.Where("number=?", number).First(&l)
	}
	o0 := model.OrderInfo{}
	if result := global.DB.Where("license_id=? AND status <> ?", l.ID, "TRADE_SUCCESS").First(&o0); result.RowsAffected != 0 {
		if o0.Status != "TRADE_SUCCESS" {
			return errors.New("订单已存在")
		}
	}
	o := model.OrderInfo{
		UserID:    l.UserID,
		User:      l.User, // l.User == nil
		BlockID:   b.ID,
		Block:     b,
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

func GetOrderList(pn, psize, year, month, day int) ([]model.OrderResp, int, error) {
	var oo []model.OrderInfo
	var data []model.OrderResp
	if year == 1000 {
		var ordersCount int64
		global.DB.Model(&model.OrderInfo{}).Count(&ordersCount)
		result := global.DB.Preload("License").Preload("Block").Scopes(Paginate(pn, psize)).Find(&oo)
		for _, v := range oo {
			data = append(data, model.OrderResp{
				OrderSn:       v.OrderSn,
				PayType:       PayType2Chn(v.PayType),
				Status:        Status2Chn(v.Status),
				BlockNo:       v.Block.BlockNo,
				OrderMount:    v.OrderMount,
				StartTime:     PayTime2Chn(v.StartTime),
				PayTime:       PayTime2Chn(v.PayTime),
				LicenseNumber: v.License.Number,
			})
		}
		count := ordersCount
		return data, int(count), result.Error
	} else {
		result := global.DB.Preload("License").Preload("Block").Find(&oo)
		date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
		nextDate := date.Add(time.Hour * 24)
		for _, v := range oo {
			if v.StartTime.After(date) && v.StartTime.Before(nextDate) {
				data = append(data, model.OrderResp{
					OrderSn:       v.OrderSn,
					PayType:       PayType2Chn(v.PayType),
					Status:        Status2Chn(v.Status),
					BlockNo:       v.Block.BlockNo,
					OrderMount:    v.OrderMount,
					StartTime:     PayTime2Chn(v.StartTime),
					PayTime:       PayTime2Chn(v.PayTime),
					LicenseNumber: v.License.Number,
				})
			}
		}
		count := result.RowsAffected
		if len(data) <= psize {
			return data, int(count), result.Error
		} else {
			st := pn*psize - 1
			ed := st + psize
			if ed >= len(data) {
				ed = len(data) - 1
			}
			return data[st : ed+1], int(count), result.Error
		}
	}

}

func GetUserOrderList(pn, psize int, id string) ([]model.OrderResp, int, error) {
	var oo []model.OrderInfo
	u := model.User{}
	var ordersCount int64
	result := global.DB.Where("open_id=?", id).First(&u)
	if result.RowsAffected == 0 {
		return nil, 0, errors.New("用户不存在")
	}
	ordersCount = result.RowsAffected
	var data []model.OrderResp
	localDB := global.DB
	localDB = localDB.Where("user_id=?", u.ID)
	result = localDB.Preload("License").Preload("Block").Scopes(Paginate(pn, psize)).Find(&oo)
	for _, v := range oo {
		data = append(data, model.OrderResp{
			OrderSn:       v.OrderSn,
			PayType:       PayType2Chn(v.PayType),
			Status:        Status2Chn(v.Status),
			BlockNo:       v.Block.BlockNo,
			OrderMount:    v.OrderMount,
			StartTime:     PayTime2Chn(v.StartTime),
			PayTime:       PayTime2Chn(v.PayTime),
			LicenseNumber: v.License.Number,
		})
	}
	count := int(ordersCount)
	return data, count, result.Error
}

func GetLicenseOrderList(pn, psize int, number string) ([]model.OrderResp, int, error) {
	var oo []model.OrderInfo
	l := model.License{}
	var ordersCount int64
	global.DB.Model(&model.OrderInfo{}).Count(&ordersCount)
	if result := global.DB.Where("number=?", number).First(&l); result.RowsAffected == 0 {
		return nil, 0, errors.New("车牌不存在")
	}
	var data []model.OrderResp
	localDB := global.DB
	localDB = localDB.Where("license_id=?", l.ID)
	result := localDB.Preload("License").Preload("Block").Scopes(Paginate(pn, psize)).Find(&oo)
	for _, v := range oo {
		data = append(data, model.OrderResp{
			OrderSn:       v.OrderSn,
			PayType:       PayType2Chn(v.PayType),
			Status:        Status2Chn(v.Status),
			BlockNo:       v.Block.BlockNo,
			OrderMount:    v.OrderMount,
			StartTime:     PayTime2Chn(v.StartTime),
			PayTime:       PayTime2Chn(v.PayTime),
			LicenseNumber: v.License.Number,
		})
	}
	count := int(ordersCount)
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
	dur := end.Sub(*o.StartTime)
	//dur := time.Since(*o.StartTime)
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

const API_KEY = "jS5yvCqC2BzcBIScnxQid6jA"
const SECRET_KEY = "hxczoXkqp2Plx0kInLiTewyxWn3Ssorv"

var TOKEN string

func GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", API_KEY, SECRET_KEY)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	accessTokenObj := map[string]string{}
	json.Unmarshal([]byte(body), &accessTokenObj)
	return accessTokenObj["access_token"]
}

func Recognize(image string) (model.NumberResp, error) {
	URL := "https://aip.baidubce.com/rest/2.0/ocr/v1/license_plate?access_token=" + GetAccessToken()
	payload := strings.NewReader("image=" + url.QueryEscape(image))
	client := &http.Client{}
	req, err := http.NewRequest("POST", URL, payload)

	if err != nil {
		zap.S().Error(err)
		return model.NumberResp{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		zap.S().Error(err)
		return model.NumberResp{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		zap.S().Error(err)
		return model.NumberResp{}, err
	}
	var number form.NumberResp
	err = json.Unmarshal(body, &number)
	if number.WordsResult.Number == "" {
		return model.NumberResp{}, errors.New("图片目标识别错误")
	}
	return model.NumberResp{
		Number: number.WordsResult.Number,
	}, nil
}
