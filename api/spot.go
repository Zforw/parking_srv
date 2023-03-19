package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"parking/form"
	"parking/handler"
	"parking/utils"
	"strconv"
)

func CreateSpot(ctx *gin.Context) {
	s := form.CreateSpotForm{}
	if err := ctx.ShouldBind(&s); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	zap.S().Info("创建停车位 ", s)
	err := handler.CreateSpot(s.SpotNo, s.X, s.Y)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}

func UpdateSpot(ctx *gin.Context) {
	l := form.UpdateLicenseForm{}
	if err := ctx.ShouldBind(&l); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	zap.S().Info("更新停车位 ", l)
	err := handler.UpdateLicense(l.Number, l.Status)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}

func GetSpotList(ctx *gin.Context) {
	pn, _ := strconv.Atoi(ctx.DefaultQuery("pn", "0"))
	pSize, _ := strconv.Atoi(ctx.DefaultQuery("psize", "90"))
	zap.S().Info("获取停车位列表, pn=", pn, "psize=", pSize)
	data, count, err := handler.GetLicenseList(pn, pSize)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"count": 0,
			"data":  nil,
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"count": count,
			"data":  data,
			"error": nil,
		})
	}
}
