package form

type CreateBlockForm struct {
	BlockNo string  `json:"blockNo" binding:"required"`
	Lat     float64 `json:"lat" binding:"required"`
	Lgt     float64 `json:"lgt" binding:"required"`
}

type CreateSpotForm struct {
	SpotNo  string `json:"spotNo" binding:"required"`
	BlockNo string `json:"blockNo" binding:"required"`
}

type UpdateSpotForm struct {
	SpotNo string `json:"number" binding:"required"`
	Status string `json:"status" binding:"required,oneof=TU NTU"`
}
