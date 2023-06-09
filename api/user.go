package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"parking/form"
	"parking/global"
	"parking/handler"
	"parking/middlewares"
	"parking/utils"
	"strconv"
	"time"
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
	pn, _ := strconv.Atoi(ctx.DefaultQuery("pn", "1"))
	pSize, _ := strconv.Atoi(ctx.DefaultQuery("psize", "10"))
	zap.S().Info("【获取用户列表】pn=", pn, ", psize=", pSize)
	data, count, err := handler.GetUserList(pn, pSize)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  1,
			"count": 0,
			"data":  nil,
			"msg":   "获取失败，" + err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  0,
			"count": count,
			"data":  data,
			"msg":   "",
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
			"code": 1,
			"msg":  "注册失败，" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "注册成功",
	})
}

func Login(ctx *gin.Context) {
	ad := form.CreateAdminForm{}
	if err := ctx.ShouldBind(&ad); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	zap.S().Info("【管理员登录】 ", ad.OpenId)
	err := handler.Login(ad.OpenId, ad.Pass)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"token": "",
		})
		return
	}
	j := middlewares.NewJWT()
	claims := form.CustomClaims{
		ID:          ad.OpenId,
		AuthorityID: 1,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*global.ServerConfig.JWTInfo.ExpireHour, //24小时过期
			Issuer:    "ZHP",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		zap.S().Error("验证成功, 但生成token失败", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "生成token失败",
			"token": "",
		})
		return
	}
	zap.S().Info("验证成功, 生成token", token)
	ctx.JSON(http.StatusOK, gin.H{
		"error": "",
		"token": token,
	})
	zap.S().Info(ad.OpenId, "成功登录")
}
