package service

import (
	"github.com/go-xorm/xorm"
	"test1/model"
)

type AdminService interface {
	//管理员登陆
	PostLogin(userName, password string) (model.Admin, bool)
	//管理员数量
	GetCount() (int64, error)
}

func GetAdminService(engine *xorm.Engine) AdminService {
	return &adminServiceWrap{
		engine: engine,
	}
}

type adminServiceWrap struct {
	engine *xorm.Engine
}

func (Service adminServiceWrap) PostLogin(userName, password string) (model.Admin, bool) {
	var admin model.Admin
	Service.engine.Where("admin_name = ? and pwd = ?", userName, password).Get(&admin)
	return admin, admin.AdminId != 0
}

func (Service adminServiceWrap) GetCount() (int64, error) {
	count, err := Service.engine.Count(new(model.Admin))
	return count, err
}
