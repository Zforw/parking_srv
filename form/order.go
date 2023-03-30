package form

type CreateOrderForm struct {
	Number string `json:"number" binding:"required"`
}

type UpdateOrderForm struct {
	Number  string `json:"number" binding:"required"`
	PayType string `json:"spotNo" binding:"required"` //alipay(支付宝)， wechat(微信)，cash(现金)
}

type SetMoneyForm struct {
	A int `json:"a" binding:"required"`
	B int `json:"b" binding:"required"`
	C int `json:"c" binding:"required"`
	D int `json:"d" binding:"required"`
}
