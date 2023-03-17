package form

type CreateLicenseForm struct {
	Number string `json:"number"`
	OpenId string `json:"id"`
}

type UpdateLicenseForm struct {
	Number string `json:"number"`
	Status string `json:"status" binding:"required,oneof=IN OUT"`
}
