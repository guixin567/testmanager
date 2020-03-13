package service

import (
	"github.com/go-xorm/xorm"
	"test1/model"
)

type AdminService interface {
	AdminLogin(userName, password string) (model.Admin, bool)
}

func GetAdminService(engine *xorm.Engine) AdminService {
	return &adminServiceWrap{
		engine: engine,
	}
}

type adminServiceWrap struct {
	engine *xorm.Engine
}

func (AdminService adminServiceWrap) AdminLogin(userName, password string) (model.Admin, bool) {
	var admin model.Admin
	AdminService.engine.Where("admin_name = ? and pwd = ?", userName, password).Get(&admin)
	return admin, admin.AdminId != 0
}
