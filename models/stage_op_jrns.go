package models

import (
	"github.com/astaxie/beego/orm"
)

type StageOpJrns struct {
	Id          int64 `orm:"auto" json:"id"`
	OpId        int64 `json:"op_id"`
	//ExTime orm.DateTimeField `json:"ex_time"`
	UserId int64 `json:"user_id"`
	DealId int64 `json:"deal_id"`
}



func init() {
	// Need to register model in init
	orm.RegisterModel(new(StageOpJrns))
}
