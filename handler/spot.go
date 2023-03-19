package handler

import (
	"errors"
	"parking/global"
	"parking/model"
)

func CreateSpot(number string, lat, lgt float64) error {
	spot := model.Spot{
		SpotNo: number,
		Status: "NTU",
		Lat:    lat,
		Lgt:    lgt,
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

func GetSpotList(pn, psize int) ([]model.SpotResp, int, error) {
	var spots []model.Spot
	result := Paginate(pn, psize)(global.DB).Find(&spots)
	var data []model.SpotResp
	for _, v := range spots {
		data = append(data, model.SpotResp{
			SpotNo: v.SpotNo,
			Status: v.Status,
			Lat:    v.Lat,
			Lgt:    v.Lgt,
		})
	}
	count := int(result.RowsAffected)
	return data, count, result.Error
}
