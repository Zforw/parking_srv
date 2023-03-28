package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"parking/form"
	"parking/handler"
	"parking/utils"
	"strconv"
	"time"
)

func CreateOrder(ctx *gin.Context) {
	st := time.Now()
	o := form.CreateOrderForm{}
	if err := ctx.ShouldBind(&o); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	zap.S().Info("【创建订单】 ", o, st.Format("2006-01-02-15:04:05"))
	err := handler.CreateOrder(o.Number, st)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	zap.S().Info("修改车牌 ", o.Number, " 状态为IN")
	ctx.JSON(http.StatusOK, gin.H{
		"error": "",
	})
}

func UpdateOrder(ctx *gin.Context) {
	s := form.UpdateOrderForm{}
	if err := ctx.ShouldBind(&s); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	zap.S().Info("【更新订单状态】 ", s)
	err := handler.UpdateOrder(s.Number, s.PayType)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": "",
	})
}

func GetOrderList(ctx *gin.Context) {
	pn, _ := strconv.Atoi(ctx.DefaultQuery("pn", "0"))
	pSize, _ := strconv.Atoi(ctx.DefaultQuery("psize", "90"))
	zap.S().Info("【获取订单列表】pn=", pn, ", psize=", pSize)
	data, count, err := handler.GetOrderList(pn, pSize)
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
			"error": "",
		})
	}
}

func GetUserOrderList(ctx *gin.Context) {
	id := ctx.Query("id")
	zap.S().Info("【获取用户订单列表】 open_id=", id)
	data, count, err := handler.GetLicenseOrderList(id)
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
			"error": "",
		})
	}
}

func GetLicenseOrderList(ctx *gin.Context) {
	number := ctx.Query("number")
	zap.S().Info("【获取车牌订单列表】 number=", number)
	data, count, err := handler.GetLicenseOrderList(number)
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
			"error": "",
		})
	}
}

func GetMoney(ctx *gin.Context) {
	end := time.Now()
	number := ctx.Query("number")
	zap.S().Info("【计算金额】 ", number, end.Format("2006-01-02-15:04:05"))
	money, err := handler.GetMoney(number, end)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"money": 0,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": "",
		"money": money,
	})
}
