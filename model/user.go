package model

import "time"

//用户信息
//用户ID，用户名称，用户注册时间，用户移动手机号，用户是否激活，用户的账户余额，用户的头像，用户的账户密码，是否被删除的标志字段，软删除，用户所在的城市名称，城市
type User struct {
	Id           int64     `xorm:"pk autoincr" json:"id"`        //用户ID，主键，自增
	UserName     string    `xorm:"varchar(12)" json:"user_name"` //用户名
	RegisterTime time.Time `json:"register_time"`                //注册时间
	Mobile       string    `xorm:"varchar(12)" json:"mobile"`    //手机号
	IsActive     int64     `json:"is_active"`                    //是否激活
	Balance      int64     `json:"balance"`                      //账户余额
	Avatar       string    `xorm:"varchar(255)" json:"avatar"`   //头像地址
	Pwd          string    `json:"pwd"`                          //密码
	DelFlag      string    `json:"del_flag"`                     //软删除，是否被删除
	CityName     string    `xorm:"varchar(12)" json:"city_name"` //用户所在城市
	City         *City     `xorm:"- <- ->"`
}
