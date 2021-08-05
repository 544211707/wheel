package service

import (
	"brown/dao"
	"brown/model"
)

type WheelService interface {//接口
	Wheel(prize model.WheelPrize) (result model.Result)//函数名(参数名 类型)(返回值 类型)
}

type wheelServices struct {//结构体
	dao dao.WheelDao //属性名 类型
}

func NewWheelServices() WheelService {
	return &wheelServices{
		dao: dao.NewWheelDao(),
	}
}

func (p wheelServices) Wheel(wheelPrize model.WheelPrize) (result model.Result) {
	//添加
	result.Data = p.dao.Wheel(wheelPrize)
	result.Msg = "Success"
	result.Code = 0
	return
}
