package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"parking/form"
	"parking/handler"
	"parking/utils"
	"strconv"
)

func CreateUser(ctx *gin.Context) {
	u := form.CreateUserForm{}
	if err := ctx.ShouldBind(&u); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	zap.S().Info("【创建用户】 ", u)
	err := handler.CreateUser(0, u.OpenId, "")
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": "",
	})
}

func GetUserList(ctx *gin.Context) {
	pn, _ := strconv.Atoi(ctx.DefaultQuery("pn", "0"))
	pSize, _ := strconv.Atoi(ctx.DefaultQuery("psize", "90"))
	zap.S().Info("【获取用户列表】pn=", pn, ", psize=", pSize)
	data, count, err := handler.GetUserList(pn, pSize)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"count": 0,
			"data":  nil,
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"count": count,
			"data":  data,
			"error": "",
		})
	}
}

func Register(ctx *gin.Context) {
	ad := form.CreateAdminForm{}
	if err := ctx.ShouldBind(&ad); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	ad.Pass = utils.EncryptPass(ad.Pass)
	zap.S().Info("【创建管理员】 ", ad)
	err := handler.CreateUser(1, ad.OpenId, ad.Pass)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": "",
	})
}

func Login(ctx *gin.Context) {
	ad := form.CreateAdminForm{}
	if err := ctx.ShouldBind(&ad); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	ad.Pass = utils.EncryptPass(ad.Pass)
	zap.S().Info("【管理员登录】 ", ad)
	err := handler.Login(ad.OpenId, ad.Pass)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": "",
		"token": "",
	})
}
