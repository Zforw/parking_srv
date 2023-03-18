package global

import (
	ut "github.com/go-playground/universal-translator"
	"gorm.io/gorm"
	"parking/config"
)

var (
	DB           *gorm.DB
	Trans        ut.Translator
	ServerConfig *config.ServerConfig
	NacosConfig  *config.NacosConfig
)
