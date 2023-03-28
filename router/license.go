package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"parking/api"
)

func InitLicenseRouter(group *gin.RouterGroup) {
	LicenseRouter := group.Group("l")
	zap.S().Debug("配置车牌相关的URL")
	{
		LicenseRouter.POST("add", api.CreateLicense)
		LicenseRouter.POST("delete", api.DeleteLicense)
		//LicenseRouter.POST("update", api.UpdateLicense)
		LicenseRouter.GET("list", api.GetLicenseList)
		LicenseRouter.GET("ulist", api.GetUserLicenseList)
	}
}
