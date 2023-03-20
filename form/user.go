package form

import "github.com/dgrijalva/jwt-go"

type CreateUserForm struct {
	OpenId string `json:"id" binding:"required"`
}

type CreateAdminForm struct {
	OpenId string `json:"id" binding:"required"`
	Pass   string `json:"pass" binding:"required"`
}

type CustomClaims struct {
	ID          string
	AuthorityID uint
	jwt.StandardClaims
}
