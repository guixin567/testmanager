package service

import (
	"github.com/go-xorm/xorm"
	"test1/model"
	"time"
)

//1,定义Service,并声明方法
type StatisticService interface {
	GetUserCount(date string) int64
	GetOrderCount(date string) int64
	GetAdminCount(date string) int64
}

//4,MVC配置时，将数据库转给Service，并将Service赋值给Controller
func GetStatisticService(Engine *xorm.Engine) StatisticService {
	return &StatisticServiceWrap{Engine: Engine}
}

//2,定义一个Service的包装类，并持有数据库Engine
type StatisticServiceWrap struct {
	Engine *xorm.Engine
}

//3,在包装类中重写Service中声明的方法
//获取当天用户数量
func (service *StatisticServiceWrap) GetUserCount(date string) int64 {
	//如果是当天，获取当前时间
	if date == "NaN-NaN-NaN" {
		date = time.Now().Format("2006-01-02")
	}
	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}
	//结束时间通过加1天得到
	endDate := startDate.AddDate(0, 0, 1)
	count, err := service.Engine.Where("register_time between ? and ? and del_flag = 0", startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05")).Count(new(model.User))
	if err != nil {
		return 0
	}
	return count
}

//获取当天管理员数量
func (service *StatisticServiceWrap) GetAdminCount(date string) int64 {
	//如果是当天，获取当前时间
	if date == "NaN-NaN-NaN" {
		date = time.Now().Format("2006-01-02")
	}
	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}
	//结束时间通过加1天得到
	endDate := startDate.AddDate(0, 0, 1)
	count, err := service.Engine.Where("create_time between ? and ?", startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05")).Count(new(model.Admin))
	if err != nil {
		return 0
	}
	return count
}

//获取当天订单数量
func (service *StatisticServiceWrap) GetOrderCount(date string) int64 {
	//如果是当天，获取当前时间
	if date == "NaN-NaN-NaN" {
		date = time.Now().Format("2006-01-02")
	}
	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}
	//结束时间通过加1天得到
	endDate := startDate.AddDate(0, 0, 1)
	count, err := service.Engine.Where("time between ? and ?", startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05")).Count(new(model.UserOrder))
	if err != nil {
		return 0
	}
	return count
}
