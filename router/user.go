package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"parking/api"
	"parking/middlewares"
)

func InitUserRouter(group *gin.RouterGroup) {
	UserRouter := group.Group("u")
	zap.S().Debug("配置用户相关的URL")
	{
		//UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("add", api.CreateUser)
		UserRouter.GET("list", middlewares.JWTAuth(), api.GetUserList)
		UserRouter.POST("register", middlewares.JWTAuth(), middlewares.AdminAuth(), api.Register)
		UserRouter.POST("login", api.Login)
	}
}
