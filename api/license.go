package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"parking/form"
	"parking/handler"
	"strconv"
)

func CreateLicense(ctx *gin.Context) {
	l := form.CreateLicenseForm{}
	if err := ctx.ShouldBind(&l); err != nil {
		zap.S().Error(err.Error())
		return
	}
	zap.S().Info(l)
	err := handler.CreateLicense(l.Number, l.OpenId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

func GetLicenseList(ctx *gin.Context) {
	pn, _ := strconv.Atoi(ctx.DefaultQuery("pn", "0"))
	pSize, _ := strconv.Atoi(ctx.DefaultQuery("psize", "90"))
	zap.S().Info(pn, pSize)
	data, count := handler.GetLicenseList(pn, pSize)
	ctx.JSON(http.StatusOK, gin.H{
		"count": count,
		"data":  data,
	})
}
