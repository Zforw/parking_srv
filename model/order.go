package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int32          `gorm:"primarykey;type:int" json:"id"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"-"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	IsDeleted bool           `json:"-"`
}

// OrderInfo 订单信息
type OrderInfo struct {
	BaseModel

	UserID  int32
	User    User   `gorm:"foreignKey:UserID"`
	OrderSn string `gorm:"type:varchar(30);index"` //订单号，自己生成的订单号
	PayType string `gorm:"type:varchar(20) comment 'alipay(支付宝)， wechat(微信)，cash(现金)'"`

	Status     string     `gorm:"type:varchar(20)  comment 'PAYING(待支付), TRADE_SUCCESS(成功)，WAIT_BUYER_PAY(交易创建)'"`
	TradeNo    string     `gorm:"type:varchar(100) comment '交易号'"` //支付宝的订单号 查账
	OrderMount float32    //金额
	StartTime  *time.Time `gorm:"type:datetime"` //开始时间
	PayTime    *time.Time `gorm:"type:datetime"` //结束时间

	BlockID int32
	Block   Block `gorm:"foreignKey:BlockID"`

	LicenseID int32
	License   License `gorm:"foreignKey:LicenseID"`

	EmpID int32
	Emp   User `gorm:"foreignKey:EmpID"`
}

type OrderResp struct {
	OrderSn       string  //订单号，自己生成的订单号
	PayType       string  //alipay(支付宝)， wechat(微信)，cash(现金)
	Status        string  //PAYING(待支付), TRADE_SUCCESS(成功)，WAIT_BUYER_PAY(交易创建)
	OrderMount    float32 //金额
	StartTime     string  //开始时间
	PayTime       string  //支付时间
	LicenseNumber string
	BlockNo       string
	Emp           string
}

func (OrderInfo) TableName() string {
	return "order_info"
}

type Charge struct {
	BaseModel
	A int32
	B int32
	C int32
}

type ChargeResp struct {
	A int `json:"a"`
	B int `json:"b"`
	C int `json:"c"`
}

type MoneyResp struct {
	Money float64 `json:"money"`
}

type NumberResp struct {
	Number string `json:"number"`
}

func (Charge) TableName() string {
	return "charge"
}
