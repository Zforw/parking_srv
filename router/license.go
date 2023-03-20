package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"parking/api"
	"parking/middlewares"
)

func InitLicenseRouter(group *gin.RouterGroup) {
	LicenseRouter := group.Group("l")
	zap.S().Debug("配置车牌相关的URL")
	{
		LicenseRouter.POST("add", middlewares.JWTAuth(), api.CreateLicense)
		LicenseRouter.POST("update", api.UpdateLicense)
		LicenseRouter.GET("list", api.GetLicenseList)
		LicenseRouter.GET("ulist", api.GetUserLicenseList)
	}
}
