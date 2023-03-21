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
	SpotNo     string `json:"spotNo" binding:"required"`
	NewSpotNo  string `json:"newSpotNo" binding:"required"`
	BlockNo    string `json:"blockNo" binding:"required"`
	NewBlockNo string `json:"newBlockNo" binding:"required"`
}

type UpdateBlockForm struct {
	BlockNo    string `json:"blockNo" binding:"required"`
	NewBlockNo string `json:"newBlockNo" binding:"required"`

	Lat float64 `json:"lat" binding:"required"`
	Lgt float64 `json:"lgt" binding:"required"`
}
