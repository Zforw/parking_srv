package tests

import (
	"go.uber.org/zap"
	"parking/handler"
	"parking/model"
	"testing"
)

func TestCreateUser(t *testing.T) {
	user := &model.User{
		OpenId: "ow6HF5Y8rSlm61hl8igL1k9nOlAI",
	}
	err := handler.CreateUser(user)
	if err != nil {
		zap.S().Fatal(err)
		return
	}
}
