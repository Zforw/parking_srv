package tests

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"parking/global"
	"parking/model"
	"testing"
	"time"
)

var DB *gorm.DB

func init() {
	c := global.ServerConfig.MySqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Pass, c.Host, c.Port, c.DB)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // 禁用彩色打印
		},
	)

	// 全局模式
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}

func TestCreateUser(t *testing.T) {
	user := &model.User{
		OpenId: "ow6HF5Y8rSlm61hl8igL1k9nOlAI",
	}
	err := DB.Create(&user)
	if err != nil {
		zap.S().Fatal(err)
		return
	}
}

func TestCreateLicense(t *testing.T) {
	var user model.User
	user.OpenId = "ow6HF5Y8rSlm61hl8igL1k9nOlAI"
	if result := DB.First(&user); result.RowsAffected == 0 {
		zap.S().Debug("用户不存在")
		return
	}
	license := &model.License{
		Number: "鄂A-1XC25",
		UserID: 1,
		User:   user,
		Status: "OUT",
	}
	err := DB.Create(&license)
	if err != nil {
		zap.S().Fatal(err)
		return
	}
}
