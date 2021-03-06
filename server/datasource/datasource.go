package datasource

import (
	"brown/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	masterEngine *xorm.Engine
)

func InstanceMaster() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}

	engine, err := xorm.NewEngine("mysql", "root:960919abc@/finego?charset=utf8")
	if err != nil {
		log.Fatal("dbhelper.DbInstanceMaster,", err)
		return nil
	}
	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(true)


	masterEngine = engine
	return engine
}

func init() {
	var err error
	db, err := xorm.NewEngine("mysql", "root:960919abc@/finego?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err := db.Sync2(new(model.WheelPrize), new(model.PrizeContent)); err != nil { //同步struct到数据库表
		log.Fatal("数据表同步失败:", err)
	}
	masterEngine = db
}
