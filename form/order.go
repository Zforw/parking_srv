package form

type CreateOrderForm struct {
	Number string `json:"number" binding:"required"`
}

type UpdateOrderForm struct {
	Number  string `json:"number" binding:"required"`
	PayType string `json:"payType" binding:"required"` //alipay(支付宝)， wechat(微信)，cash(现金)
}

type NumberImageForm struct {
	Base64 string `json:"base64" binding:"required"`
}

type SetMoneyForm struct {
	A int `json:"a" binding:"required"`
	B int `json:"b" binding:"required"`
	C int `json:"c" binding:"required"`
}

type VertexesLocation struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type WordsResult struct {
	Number           string             `json:"number"`
	VertexesLocation []VertexesLocation `json:"vertexes_location"`
	Color            string             `json:"color"`
	Probability      []float64          `json:"probability"`
}

type NumberResp struct {
	WordsResult WordsResult `json:"words_result"`
	LogId       int64       `json:"log_id"`
}
