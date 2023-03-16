package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"parking/form"
	"parking/handler"
)

func CreateUser(ctx *gin.Context) {
	u := form.CreateUserForm{}
	err := handler.CreateUser(u.OpenId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

func GetUserList(ctx *gin.Context) {

}

func PasswordLogin(ctx *gin.Context) {

}

func Register(ctx *gin.Context) {

}

func GetUserByMobile(ctx *gin.Context) {

}

func GetUserById(ctx *gin.Context) {

}

func UpdateUser(ctx *gin.Context) {

}
