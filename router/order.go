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
		OrderRouter.GET("money", middlewares.JWTAuth(), api.GetMoney)
	}
}
