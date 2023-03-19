package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"parking/global"
	"strings"
)

func Struct2String(fields map[string]string) string {
	errs := ""
	for _, err := range fields {
		errs += err + ","
	}
	return errs
}

func RemoveTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleValidatorError(ctx *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		zap.S().Error(err.Error())
		return
	}
	er := Struct2String(RemoveTopStruct(errs.Translate(global.Trans)))
	zap.S().Error(er)
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": er,
	})
}
