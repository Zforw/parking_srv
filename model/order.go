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

	User    User   //用户对象
	OrderSn string `gorm:"type:varchar(30);index"` //订单号，自己生成的订单号
	PayType string `gorm:"type:varchar(20) comment 'alipay(支付宝)， wechat(微信)，cash(现金)'"`

	Status     string     `gorm:"type:varchar(20)  comment 'PAYING(待支付), TRADE_SUCCESS(成功)，WAIT_BUYER_PAY(交易创建), TRADE_FINISHED(交易结束)'"`
	TradeNo    string     `gorm:"type:varchar(100) comment '交易号'"` //支付宝的订单号 查账
	OrderMount float32    //金额
	StartTime  *time.Time `gorm:"type:datetime"` //开始时间
	PayTime    *time.Time `gorm:"type:datetime"` //结束时间

	License License //车牌对象
}

func (OrderInfo) TableName() string {
	return "order"
}
