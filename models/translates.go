package models

import (
	"github.com/astaxie/beego/orm"
)

type Translates struct {
	Id          int64 `orm:"auto"`
	Code    	string
	En      	string
	Ru 			string
	Kk			string
	Description string
	AddInfo 	string
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Translates))
}
