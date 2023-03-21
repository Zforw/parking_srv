package handler

import (
	"errors"
	"parking/global"
	"parking/model"
)

func CreateBlock(blockNo string, lat, lgt float64) error {
	block := model.Block{}
	if result := global.DB.Where("block_no=?", blockNo).First(&block); result.RowsAffected != 0 {
		return errors.New("停车区已存在")
	}
	block.Lat = lat
	block.Lgt = lgt
	res := global.DB.Create(&block)
	return res.Error
}

func CreateSpot(blockNo, number string) error {
	block := model.Block{}
	if result := global.DB.Where("block_no=?", blockNo).First(&block); result.RowsAffected == 0 {
		return errors.New("停车区不存在")
	}
	spot := model.Spot{
		BlockID: block.ID,
		Block:   block,
		SpotNo:  number,
		Status:  "NTU",
	}
	if result := global.DB.Where("spot_no=?", number).First(&spot); result.RowsAffected != 0 {
		return errors.New("停车位已存在")
	}
	res := global.DB.Create(&spot)
	return res.Error
}

func UpdateSpot(number, status string) error {
	spot := model.Spot{
		SpotNo: number,
	}
	if result := global.DB.Where("spot_no=?", number).First(&spot); result.RowsAffected == 0 {
		return errors.New("停车位不存在")
	}
	spot.Status = status
	res := global.DB.Save(&spot)
	return res.Error
}

func GetBlockList(pn, psize int) ([]model.BLockResp, int, error) {
	var blocks []model.Block
	result := global.DB.Scopes(Paginate(pn, psize)).Find(&blocks)
	var data []model.BLockResp
	for _, v := range blocks {
		data = append(data, model.BLockResp{
			BlockNo: v.BlockNo,
			Lat:     v.Lat,
			Lgt:     v.Lgt,
		})
	}
	count := int(result.RowsAffected)
	return data, count, result.Error
}

func GetSpotList(pn, psize int) ([]model.SpotResp, int, error) {
	var spots []model.Spot
	result := global.DB.Preload("Block").Scopes(Paginate(pn, psize)).Find(&spots)
	var data []model.SpotResp
	for _, v := range spots {
		data = append(data, model.SpotResp{
			SpotNo: v.SpotNo,
			Status: v.Status,
			Lat:    v.Block.Lat,
			Lgt:    v.Block.Lgt,
		})
	}
	count := int(result.RowsAffected)
	return data, count, result.Error
}
