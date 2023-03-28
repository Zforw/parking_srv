package model

type Block struct {
	BaseModel
	BlockNo string  `gorm:"type:varchar(20)"` //停车区编号：A、B、C...
	Lat     float64 `gorm:"not null"`         //经度
	Lgt     float64 `gorm:"not null"`         //纬度
}

// Spot 停车位
type Spot struct {
	BaseModel
	BlockID int32
	Block   Block  `gorm:"foreignKey:BlockID"`
	SpotNo  string `gorm:"type:varchar(20)"` //停车位编号：A01...
}

type SpotResp struct {
	BlockNo string  `json:"blockNo"`
	SpotNo  string  `json:"spotNo"`
	Lat     float64 `json:"lat"`
	Lgt     float64 `json:"lgt"`
}

type BLockResp struct {
	BlockNo string  `json:"blockNo"`
	Lat     float64 `json:"lat"`
	Lgt     float64 `json:"lgt"`
}

func (Block) TableName() string {
	return "block"
}

func (Spot) TableName() string {
	return "spot"
}
