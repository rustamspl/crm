package models

import (
	"github.com/astaxie/beego/orm"
)

type Deals struct {
	Id          int64 `orm:"auto" json:"id"`
	Title        string `json:"title"`
	Amount float64 `json:"amount"`
	AccountId int64 `json:"account_id"`
	DealStageId int `json:"deal_stage_id"`
	XCity int `json:"x_city"`
	XRegion int `json:"x_region"`
	XDealType int `json:"x_deal_type"`
	XObjectType int `json:"x_object_type"`
	XDscr string `json:"x_dscr"`
	OwnerId int64 `json:"owner_id"`
}



func init() {
	// Need to register model in init
	orm.RegisterModel(new(Deals))
}
