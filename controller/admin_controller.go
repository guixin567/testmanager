package controller

import (
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"test1/service"
)

type AdminController struct {
	//iris框架自动绑定Context
	Context iris.Context

	//admin功能逻辑
	Service service.AdminService

	//Session对象
	Session *sessions.Session
}

const (
	Admin = "Admin"
)

type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (controller *AdminController) PostLogin() mvc.Result {
	var adminLogin AdminLogin
	controller.Context.ReadJSON(&adminLogin)
	if adminLogin.UserName == "" || adminLogin.Password == "" {
		return mvc.Response{
			Object: iris.Map{
				"message": "用户名或者密码为空",
				"status":  "0",
				"success": "登陆失败",
			},
		}
	}

	admin, isSuccess := controller.Service.AdminLogin(adminLogin.UserName, adminLogin.Password)
	if !isSuccess {
		return mvc.Response{
			Object: iris.Map{
				"message": "登陆失败",
				"status":  "0",
				"success": "登陆失败",
			},
		}
	}
	adminByte, _ := json.Marshal(admin)
	controller.Session.Set(Admin, adminByte)
	return mvc.Response{
		Object: iris.Map{
			"message": "登陆成功",
			"status":  "1",
			"success": "登陆成功",
		},
	}

}
