package initialize

import (
	"github.com/gin-gonic/gin"
	router2 "parking/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	//配置跨域
	//Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/")
	router2.InitUserRouter(ApiGroup)
	return Router
}
