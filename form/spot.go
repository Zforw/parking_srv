package form

type CreateSpotForm struct {
	SpotNo string  `json:"number"`
	Status string  `json:"status" binding:"required,oneof=TU NTU"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
}

type UpdateSpotForm struct {
	SpotNo string  `json:"number"`
	Status string  `json:"status" binding:"required,oneof=TU NTU"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
}
