package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"test1/service"
	"test1/util"
)

type UserController struct {
	Context iris.Context
	Service service.UserService
	Session sessions.Session
}

func (controller *UserController) GetCount() mvc.Result {
	count, err := controller.Service.GetCount()
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"status": util.FAIL,
				"count":  "0",
			},
		}
	}
	return mvc.Response{
		Object: iris.Map{
			"status": util.OK,
			"count":  count,
		},
	}
}
