package dao

import (
	"brown/datasource"
	"brown/model"
	"fmt"
	"github.com/go-xorm/xorm"
)

type WheelDao interface {
	Wheel(prize model.WheelPrize) model.WheelPrize
}

type wheelDao struct {
	engine *xorm.Engine
}

func NewWheelDao() WheelDao {
	return &wheelDao{
		engine: datasource.InstanceMaster(),
	}
}

//抽奖写入
func (l wheelDao) Wheel(wheelPrize model.WheelPrize) model.WheelPrize {


	fmt.Println(wheelPrize,"dao")
	l.engine.Insert(wheelPrize)

	sql := "update prize_content set prize_amount = prize_amount-1 where prize_name =?"
	l.engine.Exec(sql,wheelPrize.PrizeName)




	return wheelPrize
}
