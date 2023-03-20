package form

type CreateUserForm struct {
	OpenId string `json:"id" binding:"required"`
}

type CreateAdminForm struct {
	OpenId string `json:"id" binding:"required"`
	Pass   string `json:"pass" binding:"required"`
}
