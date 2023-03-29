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

func CreateBlock(ctx *gin.Context) {
	s := form.CreateBlockForm{}
	if err := ctx.ShouldBind(&s); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	zap.S().Info("【创建停车区】 ", s)
	err := handler.CreateBlock(s.BlockNo, s.Lat, s.Lgt)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "创建失败，" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "创建成功",
	})
}

func CreateSpot(ctx *gin.Context) {
	s := form.CreateSpotForm{}
	if err := ctx.ShouldBind(&s); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	zap.S().Info("【创建停车位】 ", s)
	err := handler.CreateSpot(s.BlockNo, s.SpotNo)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "创建失败，" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "创建停车位成功",
	})
}

func UpdateSpot(ctx *gin.Context) {
	s := form.UpdateSpotForm{}
	if err := ctx.ShouldBind(&s); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	zap.S().Info("【更新停车位信息】 ", s)
	err := handler.UpdateSpot(s.SpotNo, s.BlockNo, s.NewSpotNo, s.NewBlockNo)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "更新失败，" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "更新成功",
	})
}

func UpdateBlock(ctx *gin.Context) {
	b := form.UpdateBlockForm{}
	if err := ctx.ShouldBind(&b); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	zap.S().Info("【更新停车区信息】 ", b)
	err := handler.UpdateBlock(b.BlockNo, b.NewBlockNo, b.Lat, b.Lgt)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "更新失败，" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "更新成功",
	})
}

func FindSpot(ctx *gin.Context) {
	no := ctx.Query("spot_no")
	zap.S().Info("【查找停车位】 spot_no=", no)
	data, err := handler.FindSpot(no)
	if err != nil {
		zap.S().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"data": nil,
			"msg":  "查找失败，" + err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": data,
			"msg":  "查找成功",
		})
	}
}

func GetBlockList(ctx *gin.Context) {
	pn, _ := strconv.Atoi(ctx.DefaultQuery("pn", "0"))
	pSize, _ := strconv.Atoi(ctx.DefaultQuery("psize", "90"))
	zap.S().Info("【获取停车区列表】 pn=", pn, ", psize=", pSize)
	data, count, err := handler.GetBlockList(pn, pSize)
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
			"msg":   "获取成功",
		})
	}
}

func GetSpotList(ctx *gin.Context) {
	pn, _ := strconv.Atoi(ctx.DefaultQuery("pn", "0"))
	pn -= 1
	pSize, _ := strconv.Atoi(ctx.DefaultQuery("psize", "10"))
	spotNo := ctx.DefaultQuery("spoNo", "0")
	zap.S().Info("【获取停车位列表】 pn=", pn, ", psize=", pSize)
	data, count, err := handler.GetSpotList(pn, pSize, spotNo)
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
			"msg":   "获取成功",
		})
	}
}
