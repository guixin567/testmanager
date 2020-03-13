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
	)
	engine.Insert(model.Admin{AdminId: 1, AdminName: "root", Pwd: "root"})
	if err != nil {
		panic(err.Error())
	}

	engine.ShowSQL(true)
	engine.SetMaxOpenConns(18)
	return engine
}
