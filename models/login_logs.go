package models

import (
	"github.com/astaxie/beego/orm"
)

type LoginLogs struct {
	Id          int64 `orm:"auto" json:"id"`
	UserId        int64 `json:"user_id"`
	LoginType int64 `json:"login_type"`
}




func init() {
	// Need to register model in init
	orm.RegisterModel(new(LoginLogs))
}
