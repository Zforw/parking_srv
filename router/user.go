package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"parking/api"
)

func InitUserRouter(group *gin.RouterGroup) {
	UserRouter := group.Group("user")
	zap.S().Debug("配置用户相关的URL")
	{
		//UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.GET("list", api.GetUserList)
		UserRouter.POST("login", api.PasswordLogin)
		UserRouter.POST("register", api.Register)
		UserRouter.GET("find_mobile", api.GetUserByMobile)
		UserRouter.GET("find_id", api.GetUserById)
		UserRouter.POST("update", api.UpdateUser)
	}
}
