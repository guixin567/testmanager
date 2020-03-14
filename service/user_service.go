package service

import (
	"github.com/go-xorm/xorm"
	"test1/model"
)

type UserService interface {
	GetCount() (int64, error)
}

func GetUserService(Engine *xorm.Engine) UserService {
	return &UserServiceWrap{
		Engine: Engine,
	}
}

type UserServiceWrap struct {
	Engine *xorm.Engine
}

func (service *UserServiceWrap) GetCount() (int64, error) {
	count, err := service.Engine.Where("del_flag = ?", 0).Count(new(model.User))
	return count, err
}
