package model

// Spot 停车位
type Spot struct {
	BaseModel
	SpotNo string  `gorm:"type:varchar(20);not null"` //停车位编号：A01...
	Status string  `gorm:"type:varchar(20)  comment 'NTU(未占用), TU(占用)'"`
	Lat    float64 `gorm:"not null"` //经度
	Lgt    float64 `gorm:"not null"` //纬度
}

type SpotResp struct {
	SpotNo string  `json:"number"`
	Status string  `json:"status"`
	Lat    float64 `json:"lat"`
	Lgt    float64 `json:"lgt"`
}

func (Spot) TableName() string {
	return "spot"
}
