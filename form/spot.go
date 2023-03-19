package form

type CreateSpotForm struct {
	SpotNo string  `json:"number" binding:"required"`
	X      float64 `json:"x" binding:"required"`
	Y      float64 `json:"y" binding:"required"`
}

type UpdateSpotForm struct {
	SpotNo string  `json:"number" binding:"required"`
	Status string  `json:"status" binding:"required,oneof=TU NTU"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
}
