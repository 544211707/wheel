package dao
//data access object 数据库访问对象，夹在业务逻辑与数据库资源中间。连接操纵数据库的。
import (
	"brown/datasource"
	"brown/model"
	"fmt"
	"github.com/go-xorm/xorm"
)

type PrizeListDao interface {
	PrizeList(prizeList model.PrizeContent) model.PrizeContent
}

type prizeListDao struct {
	engine *xorm.Engine
}

func NewPrizeListDao() PrizeListDao {
	return &prizeListDao{
		engine: datasource.InstanceMaster(),
	}
}

//抽奖写入
func (i prizeListDao) PrizeList(prizeList model.PrizeContent) model.PrizeContent {


	fmt.Println(prizeList,"dao")
	i.engine.Insert(prizeList)
	fmt.Println("执行dao写入了")
	return prizeList
}
