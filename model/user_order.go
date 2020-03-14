package model

import "time"

type UserOrder struct {
	Id            int64        `xorm:"pk autoincr" json:"id"`      //订单ID
	SumMoney      int64        `xorm:"default 0" json:"sum_money"` //订单总金额
	Time          time.Time    `xorm:"DateTime" json:"time"`
	OrderTime     uint64       `json:"order_time"`
	OrderStatusId int64        `xorm:"index" json:"order_status_id"`
	OrderStatus   *OrderStatus `xorm:"-"`
	UserId        int64        `xorm:"index" json:"user_id"`
	User          *User        `xorm:"-"`
	ShopId        int64        `xorm:"index" json:"shop_id"`
	Shop          *Shop        `xorm:"-"`
	AddressId     int64        `xorm:"index" json:"address_id"`
	Address       *Address     `xorm:"-"`
	DelFlag       int64        `xorm:"defult 0" json:"del_flag"`
}
