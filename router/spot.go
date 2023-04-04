package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"parking/api"
	"parking/middlewares"
)

func InitSpotRouter(group *gin.RouterGroup) {
	SpotRouter := group.Group("s")
	zap.S().Debug("配置停车位相关的URL")
	{
		SpotRouter.POST("add", middlewares.JWTAuth(), api.CreateSpot)
		SpotRouter.POST("addb", middlewares.JWTAuth(), api.CreateBlock)
		SpotRouter.POST("update_spot", middlewares.JWTAuth(), api.UpdateSpot)
		SpotRouter.POST("update_block", middlewares.JWTAuth(), api.UpdateBlock)
		SpotRouter.GET("list", middlewares.JWTAuth(), api.GetSpotList)
		SpotRouter.GET("blist", middlewares.JWTAuth(), api.GetBlockList)
		SpotRouter.GET("find_spot", middlewares.JWTAuth(), api.FindSpot)
		SpotRouter.GET("find_block", middlewares.JWTAuth(), api.FindBlock)
		SpotRouter.GET("remaining", middlewares.JWTAuth(), api.GetRemaining)
	}
}
