package restapi

import "net/http"
import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/yeldars/crm/models"
	"encoding/json"
	"strconv"
	"github.com/julienschmidt/httprouter"
	"github.com/yeldars/crm/auth"
)


type AccountsList struct {
	Id          int64 `orm:"auto" json:"id"`
	Name        string `json:"name"`
	CompanyId int `json:"company_id"`
	Website string `json:"website"`
	Address string `json:"address"`
	Phones  string `json:"phones"`
	Emails string `json:"emails"`
	OwnerName string `json:"owner_name"`
	OwnerID int64 `json:"owner_id"`
}

type AccountsContact struct{
	Id int64 `orm:"auto" json:"id"`
	Cont string `json:"cont"`
	ContType string `json:"cont_type"`
	ContTypeId int64 `json:"cont_type_id"`
	ContGroup string `json:"cont_group"`
	Deleted bool `json:"deleted"`
}

type AccountsDetails struct {
	Id          int64 `orm:"auto" json:"id"`
	Name        string `json:"name"`
	CompanyId int `json:"company_id"`
	Website string `json:"website"`
	Address string `json:"address"`
	Contacts  [] AccountsContact `json:"contacts"`

}


type accountGetResponse struct {
PageCount   int `json:"pageCount"`
SessionInfo auth.SessionInfo `json:"session_info"`
Items [] AccountsList `json:"items"`
}

type accountUpdateRequest struct {
	Items [] AccountsDetails`json:"items"`
}

type accountInsertRequest struct {
	Items [] models.Accounts `json:"items"`
}

type accountSaveRequest struct {
	Items [] models.Accounts`json:"items"`
}

type accountDeleteRequest struct {
	Items [] models.Accounts
	RowsAffected int
}


func AccountRestApiDetail(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	if RestCheckAuth(res,req){
		return
	}
	req.ParseForm()
	item  := AccountsDetails{}
	o := orm.NewOrm()
	o.Using("default")
	//err := o.QueryTable("accounts").Filter("id",req.Form.Get("id")).One(&item)

	err := o.Raw("SELECT * from accounts a where a.id=?",req.Form.Get("id")).QueryRow(&item)
	if RestCheckDBPanic(err ,res ,o ) {
		return
	}

    //var conts [] AccountsContact
	_, err = o.Raw("SELECT ac.*,ct.title cont_type,cg.title cont_group from accountconts ac,cont_types ct,cont_groups cg where cg.id=ct.group_id and ct.id=cont_type_id and ac.account_id=?",item.Id).QueryRows(&item.Contacts)

	jsonData, err := json.Marshal(item)
	if RestCheckDBPanic(err ,res ,o ) {
		return
	}
	fmt.Fprint(res,string(jsonData))

}
func AccountRestApiGet(res http.ResponseWriter, req *http.Request, _ httprouter.Params){

	if RestCheckAuth(res,req){
		return
	}

	o := orm.NewOrm()
	o.Using("default")
	var arr [] AccountsList

	req.ParseForm()
	limitFrom := "1"
	limitTo := "5"

	iPerPage,_ := strconv.ParseInt(req.Form.Get("perpage"),10,32);
	iPage,_ := strconv.ParseInt(req.Form.Get("page"),10,32);

	limitFrom = strconv.Itoa(int(iPerPage*(iPage-1)))
	limitTo = req.Form.Get("perpage");

	pageCount := 10;
	o.Raw("SELECT ceil(count(1)/?) FROM accounts",limitTo).QueryRow(&pageCount)



	_, err := o.Raw("SELECT a.*" +
	",(select GROUP_CONCAT(cont SEPARATOR ';') from accountconts ac where ac.account_id=a.id and ac.cont_type_id in (select id from cont_types where group_id=1 )) phones"+
	",(select GROUP_CONCAT(cont SEPARATOR ';') from accountconts ac where ac.account_id=a.id and ac.cont_type_id in (select id from cont_types where group_id=3 )) emails"+
	" FROM accounts a limit ?,?",limitFrom,limitTo).QueryRows(&arr)
	respO := accountGetResponse{}
	respO.Items = arr
	respO.PageCount = pageCount
	respO.SessionInfo.UserId = auth.UserId(req)
	jsonData, err := json.Marshal(respO)
	if RestCheckDBPanic(err ,res ,o ) {
		return
	}
	fmt.Fprint(res,string(jsonData))

}

func AccountRestApiInsert(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	if RestCheckAuth(res,req){
		return
	}

	decoder := json.NewDecoder(req.Body)
	var t accountInsertRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	o := orm.NewOrm()
	o.Using("default")
	o.Begin()


	qs := o.QueryTable("accounts")
	i, _ := qs.PrepareInsert()
	for index, element := range  t.Items {
		element.Id = 0;
		id,err := i.Insert(&element)
		element.Id=id

		t.Items[index].Id=id

		fmt.Printf("%d",id)
		if err != nil{
			o.Rollback()
			panic(err)
		}
	}
	i.Close()
	o.Commit()

	j,_ := json.Marshal(t)
	fmt.Fprint(res,string(j))
}


func AccountRestApiUpdate(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if RestCheckAuth(res,req){
		return
	}
	decoder := json.NewDecoder(req.Body)
	var t accountUpdateRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	o := orm.NewOrm()
	o.Using("default")
	o.Begin()

	for _,element := range t.Items {
		fmt.Fprint(res,element.Name);

		o.QueryTable("accounts").Filter("id", element.Id).Update(orm.Params{
			"name":element.Name,
			"address": element.Address,
			"website": element.Website,
		})

		for _,contact := range element.Contacts {
			//fmt.Fprint(res,contact.Cont)
			if contact.Deleted{
				o.QueryTable("accountconts").Filter("id", contact.Id).Delete()
			} else if contact.Id==0{
				var newContact = models.Accountconts{}
				newContact.Cont = contact.Cont
				newContact.AccountId = element.Id
				newContact.ContTypeId = contact.ContTypeId
				i, _ := o.QueryTable("accountconts").PrepareInsert()
				_,err := i.Insert(&newContact)
				if RestCheckDBPanic(err ,res ,o ) {
					return
				}
			}else {
				o.QueryTable("accountconts").Filter("id", contact.Id).Update(orm.Params{
					"cont":contact.Cont,
				})
			}



		}
	}

	o.Commit()

}


func AccountRestApiDelete(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	if RestCheckAuth(res,req){
		return
	}
	decoder := json.NewDecoder(req.Body)
	var t accountDeleteRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	o := orm.NewOrm()
	o.Using("default")
	o.Begin()
	RowAffected  := 0;

	for _,element := range t.Items {
		ra,_ := o.QueryTable("accounts").Filter("id", element.Id).Delete()
		RowAffected = int(ra) + RowAffected
	}

	o.Commit()

	t.RowsAffected = RowAffected
	j,_ := json.Marshal(t)
	fmt.Fprint(res,string(j))

}

func AccountRestApiRemoveAll(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if RestCheckAuth(res,req){
		return
	}
	o := orm.NewOrm()
	o.Using("default")
	o.Raw("delete from accounts").Exec()
}
