package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"parking/form"
	"parking/handler"
	"strconv"
)

func CreateUser(ctx *gin.Context) {
	u := form.CreateUserForm{}
	if err := ctx.ShouldBind(&u); err != nil {
		zap.S().Error(err.Error())
		return
	}
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
	pn, _ := strconv.Atoi(ctx.DefaultQuery("pn", "0"))
	pSize, _ := strconv.Atoi(ctx.DefaultQuery("psize", "90"))
	data, count := handler.GetUserList(pn, pSize)
	ctx.JSON(http.StatusOK, gin.H{
		"count": count,
		"data":  data,
	})
}

func UpdateUser(ctx *gin.Context) {

}
