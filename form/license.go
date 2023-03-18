package form

type CreateLicenseForm struct {
	Number string `json:"number" binding:"required"`
	OpenId string `json:"id" binding:"required"`
}

type UpdateLicenseForm struct {
	Number string `json:"number" binding:"required"`
	Status string `json:"status" binding:"required,oneof=IN OUT"`
}
