package models

import (
	"github.com/astaxie/beego/orm"
)

type Ngorods struct {
	Id          int `orm:"auto"`
	Name        string
	Address  string
	Phone string
	Email string
	Site string
	Category string
	Dscr string
	Ngid int
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Ngorods))
}
