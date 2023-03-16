package global

import (
	"gorm.io/gorm"
	"parking/config"
)

var (
	DB           *gorm.DB
	ServerConfig *config.ServerConfig
	NacosConfig  *config.NacosConfig
)
