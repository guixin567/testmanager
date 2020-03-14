package service

import (
	"github.com/go-xorm/xorm"
	"test1/model"
)

type OrderService interface {
	GetCount() (int64, error)
}

func GetOrderService(Engine *xorm.Engine) OrderService {
	return &OrderServiceWrap{
		Engine: Engine,
	}
}

type OrderServiceWrap struct {
	Engine *xorm.Engine
}

func (service *OrderServiceWrap) GetCount() (int64, error) {
	count, err := service.Engine.Count(new(model.UserOrder))
	return count, err
}
