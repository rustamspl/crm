package models

import (
	"github.com/astaxie/beego/orm"
)

type Files struct {
	Id          int64 `orm:"auto" json:"id"`
	FileName        string `json:"filename"`
	Data string `orm:type(blob)"`
}



func init() {
	// Need to register model in init
	orm.RegisterModel(new(Files))
}
