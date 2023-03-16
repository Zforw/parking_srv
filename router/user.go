package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"parking/api"
)

func InitUserRouter(group *gin.RouterGroup) {
	UserRouter := group.Group("u")
	zap.S().Debug("配置用户相关的URL")
	{
		//UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("add", api.CreateUser)
		UserRouter.GET("list", api.GetUserList)
		UserRouter.POST("update", api.UpdateUser)
	}
}
