package models

import (
	"github.com/astaxie/beego/orm"
)

type Pages struct {
	Id          	int64 `orm:"auto"`
	Title        	string
	Url 			string
	TemplateUrl 	string
	Controller 		string
}



func init() {
	// Need to register model in init
	orm.RegisterModel(new(Pages))
}
