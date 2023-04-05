package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"parking/form"
	"parking/handler"
	"parking/model"
	"parking/utils"
	"strconv"
	"strings"
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
	err := handler.CreateOrder(o.Number, o.BlockNo, st)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "订单创建失败" + err.Error(),
		})
		return
	}
	zap.S().Info("修改车牌 ", o.Number, " 状态为IN")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "订单创建成功",
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
			"code": 1,
			"msg":  "订单状态更新失败，" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "订单状态更新成功",
	})
}

func GetOrderList(ctx *gin.Context) {
	pn, _ := strconv.Atoi(ctx.DefaultQuery("pn", "1"))
	pSize, _ := strconv.Atoi(ctx.DefaultQuery("psize", "10"))
	dates := ctx.DefaultQuery("date", "1000-01-01")
	date := strings.Split(dates, "-")
	year, _ := strconv.Atoi(date[0])
	month, _ := strconv.Atoi(date[1])
	day, _ := strconv.Atoi(date[2])
	zap.S().Info("【获取订单列表】pn=", pn, ", psize=", pSize, ", year=", year, ", month=", month, ", day=", day)
	data, count, err := handler.GetOrderList(pn, pSize, year, month, day)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  1,
			"count": 0,
			"data":  nil,
			"msg":   "获取订单列表失败，" + err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  0,
			"count": count,
			"data":  data,
			"msg":   "获取订单列表成功",
		})
	}
}

func GetUserOrderList(ctx *gin.Context) {
	id := ctx.Query("id")
	pn, _ := strconv.Atoi(ctx.DefaultQuery("pn", "1"))
	pSize, _ := strconv.Atoi(ctx.DefaultQuery("psize", "10"))
	zap.S().Info("【获取用户订单列表】 open_id=", id)
	data, count, err := handler.GetUserOrderList(pn, pSize, id)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  1,
			"count": count,
			"data":  data,
			"msg":   "获取用户订单列表失败，" + err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  0,
			"count": count,
			"data":  data,
			"msg":   "获取用户订单列表成功",
		})
	}
}

func GetLicenseOrderList(ctx *gin.Context) {
	number := ctx.Query("number")
	pn, _ := strconv.Atoi(ctx.DefaultQuery("pn", "1"))
	pSize, _ := strconv.Atoi(ctx.DefaultQuery("psize", "10"))
	zap.S().Info("【获取车牌订单列表】 number=", number)
	data, count, err := handler.GetLicenseOrderList(pn, pSize, number)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  1,
			"count": count,
			"data":  data,
			"msg":   "获取订单列表失败" + err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  0,
			"count": count,
			"data":  data,
			"msg":   "获取车牌订单列表成功",
		})
	}
}

func CalcMoney(ctx *gin.Context) {
	end := time.Now()
	number := ctx.Query("number")
	zap.S().Info("【计算金额】 ", number, end.Format("2006-01-02-15:04:05"))
	money, err := handler.CalcMoney(number, end)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "计算金额失败，" + err.Error(),
			"data": model.MoneyResp{
				Money: 0,
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "计算金额成功",
		"data": model.MoneyResp{
			Money: money,
		},
	})
}

func SetCharge(ctx *gin.Context) {
	m := form.SetMoneyForm{}
	if err := ctx.ShouldBind(&m); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	zap.S().Info("【设置收费标准】 ", m)
	err := handler.SetCharge(m.A, m.B, m.C)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "设置收费标准失败，" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "设置收费标准成功",
	})
}

func GetCharge(ctx *gin.Context) {
	zap.S().Info("【查看收费标准】")
	money, err := handler.GetCharge()
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "获取收费标准失败，" + err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取收费标准成功",
		"data": money,
	})
}

func Recognize(ctx *gin.Context) {
	m := form.NumberImageForm{}
	if err := ctx.ShouldBind(&m); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	zap.S().Info("【识别车牌】")
	number, err := handler.Recognize(m.Base64)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "车牌识别失败，" + err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "车牌识别成功",
		"data": number,
	})
}
