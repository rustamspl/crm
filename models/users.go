package models

import (
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id          int64 `orm:"auto"`
	Name        string
	CompanyId int
	DeptId int
	Email string
}



func init() {
	// Need to register model in init
	orm.RegisterModel(new(Users))
}
