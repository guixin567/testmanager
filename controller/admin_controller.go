package controller

import (
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"test1/service"
	"test1/util"
)

type AdminController struct {
	//iris框架自动绑定Context
	Context iris.Context

	//admin功能逻辑
	Service service.AdminService

	//Session对象
	Session *sessions.Session
}

//管理员登陆Session的Key
const (
	Admin = "Admin"
)

//管理员登陆请求实体
type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

//管理员登陆
//1，将请求的application/json参数放到AdminLogin里面
//2,如果用记名或者密码为空，返回错误信息（这一步可能前端已经过滤）
//3，调用service查询数据库里面的管理员用户名和密码，并返回Admin实体和是否查询成功
//4,查询失败，返回错误信息
//5,将Admin解析成json.Marshal，并赋值给session
//6,返回正常信息
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
	admin, isSuccess := controller.Service.PostLogin(adminLogin.UserName, adminLogin.Password)
	if !isSuccess {
		return mvc.Response{
			Object: iris.Map{
				"message": "登陆失败",
				"status":  "0",
				"success": "登陆失败",
			},
		}
	}
	//adminByte, _ := json.Marshal(admin)
	controller.Session.Set(Admin, admin.AdminId)
	return mvc.Response{
		Object: iris.Map{
			"message": "登陆成功",
			"status":  "1",
			"success": "登陆成功",
		},
	}

}

//管理员用户信息查询
//调用service里面的GetInfo查询数据库，并返回信息
func (controller *AdminController) GetInfo() mvc.Result {
	//adminByte := controller.Session.Get(Admin)
	//if adminByte == nil {
	//	return mvc.Response{
	//		Object: iris.Map{
	//			"status":  util.UnLogin,
	//			"type":    util.ErrorUnLogin,
	//			"message": util.GetCodeMessage(util.ErrorUnLogin),
	//		},
	//	}
	//}
	//var admin model.Admin
	//err := json.Unmarshal(adminByte.([]byte), &admin)
	adminId, err := controller.Session.GetInt64(Admin)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"status":  util.UnLogin,
				"type":    util.ErrorUnLogin,
				"message": util.GetCodeMessage(util.ErrorUnLogin),
			},
		}
	}
	admin, err := controller.Service.GetAdminById(adminId)
	adminJson, err := json.Marshal(admin)
	var adminMap map[string]interface{}
	err = json.Unmarshal(adminJson, &adminMap)
	//var adminMap map[string]interface{}
	//err = json.Unmarshal(adminByte.([]byte), &adminMap)
	return mvc.Response{
		Object: iris.Map{
			"status": util.OK,
			"data":   adminMap,
		},
	}
}

//管理员用户信息查询
//调用service里面的GetInfo查询数据库，并返回信息
func (controller *AdminController) GetCount() mvc.Result {
	count, err := controller.Service.GetCount()
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"status":  util.FAIL,
				"message": "",
				"count":   "0",
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

//管理员退出
//删除Session里面对应Key的信息
func (controller *AdminController) GetSingout() mvc.Result {
	controller.Session.Delete(Admin)
	return mvc.Response{
		Object: iris.Map{
			"status":  util.OK,
			"success": util.GetCodeMessage(util.SuccessLogout),
		},
	}
}
