package service

import (
	"brown/dao"
	"brown/model"
)

type PrizeListService interface {//接口
	PrizeList(prize model.PrizeContent) (result model.Result)//函数名(参数名 类型)(返回值 类型)
}

type prizeListServices struct {//结构体
	dao dao.PrizeListDao //属性名 类型
}

func NewPrizeListServices() PrizeListService {
	return &prizeListServices{
		dao: dao.NewPrizeListDao(),
	}
}

func (i prizeListServices) PrizeList(prizeList model.PrizeContent) (result model.Result) {
	//添加
	result.Data = i.dao.PrizeList(prizeList)
	result.Msg = "Success"
	result.Code = 0
	return
}
