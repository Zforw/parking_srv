package global

import (
	"github.com/anaskhan96/go-password-encoder"
	ut "github.com/go-playground/universal-translator"
	"gorm.io/gorm"
	"parking/config"
)

var (
	OP           *password.Options
	DB           *gorm.DB
	Trans        ut.Translator
	ServerConfig *config.ServerConfig
	NacosConfig  *config.NacosConfig
)
