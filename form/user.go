package form

type CreateUserForm struct {
	OpenId string `json:"id" binding:"required"`
}
