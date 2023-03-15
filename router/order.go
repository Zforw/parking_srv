package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"parking/api"
)

func InitOrderRouter(group *gin.RouterGroup) {
	UserRouter := group.Group("user")
	zap.S().Debug("配置用户相关的URL")
	{
		//UserRouter.GET("list", middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.GET("list", api.GetOrderList)
	}
}
