package models

import (
	"github.com/astaxie/beego/orm"
)

type Accounts struct {
	Id          int64 `orm:"auto" json:"id"`
	Name        string `json:"name"`
	CompanyId int `json:"company_id"`
	Website string `json:"website"`
	Address string `json:"address"`
}


type Accountconts struct {
	Id          int64 `orm:"auto" json:"id"`
	AccountId        int64 `json:"account_id"`
	ContTypeId int64 `json:"cont_type_id"`
	Cont string `json:"cont"`

}



func init() {
	// Need to register model in init
	orm.RegisterModel(new(Accounts))
	orm.RegisterModel(new(Accountconts))
}
