package model

// Spot 停车位
type Spot struct {
	BaseModel
	SpotNo string  `gorm:"type:varchar(20);not null"` //停车位编号：A01...
	Status string  `gorm:"type:varchar(20)  comment 'NTU(未占用), TU(占用)'"`
	X      float64 `gorm:"not null"` //X坐标
	Y      float64 `gorm:"not null"` //Y坐标
}

func (Spot) TableName() string {
	return "spot"
}
