package handler

import (
	"errors"
	"parking/global"
	"parking/model"
)

func CreateBlock(blockNo string, lat, lgt float64) error {
	block := model.Block{
		BlockNo: blockNo,
		Lat:     lat,
		Lgt:     lgt,
	}
	if result := global.DB.Where("block_no=?", blockNo).First(&block); result.RowsAffected != 0 {
		return errors.New("停车区已存在")
	}
	res := global.DB.Create(&block)
	return res.Error
}

func CreateSpot(blockNo, number string) error {
	block := model.Block{
		BlockNo: blockNo,
	}
	if result := global.DB.Where("block_no=?", blockNo).First(&block); result.RowsAffected == 0 {
		return errors.New("停车区不存在")
	}
	spot := model.Spot{
		BlockID: block.ID,
		Block:   block,
		SpotNo:  number,
	}
	if result := global.DB.Where("spot_no=?", number).First(&spot); result.RowsAffected != 0 {
		return errors.New("停车位已存在")
	}
	res := global.DB.Create(&spot)
	return res.Error
}

func FindSpot(spotNo string) (model.SpotResp, error) {
	spot := model.Spot{
		SpotNo: spotNo,
	}
	if result := global.DB.Where("spot_no=?", spotNo).First(&spot); result.RowsAffected == 0 {
		return model.SpotResp{}, errors.New("停车位不存在")
	}
	return model.SpotResp{SpotNo: spot.SpotNo, BlockNo: spot.Block.BlockNo, Lat: spot.Block.Lat, Lgt: spot.Block.Lgt}, nil
}

func UpdateSpot(spotNo, blockNo, newSpotNo, newBlockNo string) error {
	block := model.Block{
		BlockNo: blockNo,
	}
	if result := global.DB.Where("block_no=?", blockNo).First(&block); result.RowsAffected == 0 {
		return errors.New("停车区不存在")
	}
	spot := model.Spot{
		SpotNo: spotNo,
	}
	if result := global.DB.Where("spot_no=?", spotNo).First(&spot); result.RowsAffected == 0 {
		return errors.New("停车位不存在")
	}
	if newSpotNo != spotNo {
		ns := model.Spot{
			SpotNo: newSpotNo,
		}
		if result := global.DB.Where("spot_no=?", spotNo).First(&ns); result.RowsAffected != 0 {
			return errors.New("新的停车位编号已被使用")
		}
		spot.SpotNo = newSpotNo
	}
	if newBlockNo != blockNo {
		nb := model.Block{
			BlockNo: newBlockNo,
		}
		if result := global.DB.Where("block_no=?", blockNo).First(&nb); result.RowsAffected == 0 {
			return errors.New("新的停车区编号不存在")
		}
		spot.BlockID = nb.ID
		spot.Block = nb
	}
	res := global.DB.Save(&spot)
	return res.Error
}

func UpdateBlock(blockNo, newBlockNo string, lat, lgt float64) error {
	block := model.Block{
		BlockNo: blockNo,
	}
	if result := global.DB.Where("block_no=?", blockNo).First(&block); result.RowsAffected == 0 {
		return errors.New("停车位不存在")
	}
	if newBlockNo != blockNo {
		nb := model.Block{
			BlockNo: newBlockNo,
		}
		if result := global.DB.Where("block_no=?", newBlockNo).First(&nb); result.RowsAffected != 0 {
			return errors.New("新的编号已被使用")
		}
		block.BlockNo = newBlockNo
	}
	if lat != 400 {
		block.Lat = lat
	}
	if lgt != 400 {
		block.Lgt = lgt
	}
	res := global.DB.Save(&block)
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
			BlockNo: v.Block.BlockNo,
			SpotNo:  v.SpotNo,
			Lat:     v.Block.Lat,
			Lgt:     v.Block.Lgt,
		})
	}
	count := int(result.RowsAffected)
	return data, count, result.Error
}
