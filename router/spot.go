package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"parking/api"
)

func InitSpotRouter(group *gin.RouterGroup) {
	SpotRouter := group.Group("s")
	zap.S().Debug("配置停车位相关的URL")
	{
		SpotRouter.POST("add", api.CreateSpot)
		SpotRouter.POST("update", api.UpdateSpot)
		SpotRouter.GET("list", api.GetSpotList)
	}
}
