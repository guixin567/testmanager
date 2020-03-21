package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"strings"
	"test1/service"
	"test1/util"
)

//controller
//定义一个struct的controller，声明上下文,service,session
//定义一个请求的方法

const (
	UserType  = "UserType"
	OrderType = "OrderType"
	AdminType = "AdminType"
)

//数据统计
//新增用户，新增订单，新增管理员
//上下文，Service,Session
type StatisticController struct {
	Context iris.Context
	Service service.StatisticService
	Session *sessions.Session
}

//
func (controller *StatisticController) GetCount() mvc.Result {
	//	/statis/user/2019-03-10/count
	path := controller.Context.Path()
	splitPath := strings.Split(path, "/")

	//如果不是5个，代表不符合请求格式
	if len(splitPath) != 5 {
		return mvc.Response{
			Object: iris.Map{
				"status": util.FAIL,
				"count":  0,
			},
		}
	}

	requestType := splitPath[2]
	date := splitPath[3]

	var result int64

	switch requestType {
	case "user":
		userCount := controller.Session.Get(UserType + date)
		if userCount != nil {
			result = int64(userCount.(float64))
		} else {
			result = controller.Service.GetUserCount(date)
			controller.Session.Set(UserType+date, result)
		}

	case "order":
		orderCount := controller.Session.Get(OrderType + date)
		if orderCount != nil {
			result = int64(orderCount.(float64))
		} else {
			result = controller.Service.GetOrderCount(date)
			controller.Session.Set(OrderType+date, result)
		}

	case "admin":
		adminCount := controller.Session.Get(AdminType + date)
		if adminCount != nil {
			result = int64(adminCount.(float64))
		} else {
			result = controller.Service.GetAdminCount(date)
			controller.Session.Set(AdminType+date, result)
		}
	}

	return mvc.Response{
		Object: iris.Map{
			"status": util.OK,
			"count":  result,
		},
	}
}
