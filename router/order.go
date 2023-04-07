package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"parking/api"
	"parking/middlewares"
)

func InitOrderRouter(group *gin.RouterGroup) {
	OrderRouter := group.Group("o")
	zap.S().Debug("配置订单相关的URL")
	{
		OrderRouter.POST("add", middlewares.JWTAuth(), api.CreateOrder)
		OrderRouter.POST("update", middlewares.JWTAuth(), api.UpdateOrder)
		OrderRouter.GET("list", middlewares.JWTAuth(), api.GetOrderList)
		OrderRouter.GET("ulist", middlewares.JWTAuth(), api.GetUserOrderList)
		OrderRouter.GET("llist", middlewares.JWTAuth(), api.GetLicenseOrderList)
		OrderRouter.GET("calc_money", middlewares.JWTAuth(), api.CalcMoney)
		OrderRouter.POST("set_charge", middlewares.JWTAuth(), middlewares.AdminAuth(), api.SetCharge)
		OrderRouter.GET("get_charge", middlewares.JWTAuth(), api.GetCharge)
		OrderRouter.POST("recognize", middlewares.JWTAuth(), api.Recognize)
	}
}
