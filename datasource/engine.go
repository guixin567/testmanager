package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"test1/model"
)

// Database init
func New() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:12345678@/test1?charset=utf8")
	err = engine.Sync2(
		new(model.Admin),
		new(model.UserOrder),
		new(model.User),
	)
	//模拟数据库数据
	//for i := 10; i < 17; i++ {
	//	date1, err := time.Parse("2006-01-02", "2020-03-09")
	//	if err!=nil {
	//		panic(err)
	//	}
	//
	//	dateNow := date1.AddDate(0,0,17-i)
	//
	//
	//	engine.Insert(model.Admin{AdminId: i, AdminName: "root"+strconv.Itoa(i), Pwd: "root"+strconv.Itoa(i),CreateTime: dateNow})
	//	engine.Insert(model.UserOrder{SumMoney: 355,OrderStatusId: 33333+int64(i),Time: dateNow})
	//	engine.Insert(model.User{RegisterTime: dateNow,UserName: "zhengxin"+strconv.Itoa(i),Mobile:"1888888888"+strconv.Itoa(i) })
	//}
	if err != nil {
		panic(err.Error())
	}

	engine.ShowSQL(true)
	engine.SetMaxOpenConns(18)
	return engine
}
