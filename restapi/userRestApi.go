package restapi

import "net/http"
import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/yeldars/crm/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"strconv"
)

type userGetResponse struct {
PageCount   int
Items [] UsersList`json:"items"`
}


type UsersList struct {
	Id    int64 `orm:"auto" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Roles string `json:"roles"`
}

type userUpdateRequest struct {
	Items [] models.Users
}

type userSaveRequest struct {
	Items [] models.Users
}

type userDeleteRequest struct {
	Items [] models.Users
	RowsAffected int
}

func UserRestApiGet(res http.ResponseWriter, req *http.Request, _ httprouter.Params){

	if RestCheckAuth(res,req){
		return
	}

	o := orm.NewOrm()
	o.Using("default")
	var arr [] UsersList

	req.ParseForm()
	limitFrom := "1"
	limitTo := "5"

	iPerPage,_ := strconv.ParseInt(req.Form.Get("perpage"),10,32);
	iPage,_ := strconv.ParseInt(req.Form.Get("page"),10,32);

	limitFrom = strconv.Itoa(int(iPerPage*(iPage-1)))
	//if limitFrom==0{

	//}
	fmt.Println(limitFrom)
	limitTo = req.Form.Get("perpage");

	pageCount := 10;
	o.Raw("SELECT ceil(count(1)/?) FROM users",limitTo).QueryRow(&pageCount)



	_, err := o.Raw("SELECT u.*" +
	",(select GROUP_CONCAT(r.title SEPARATOR ';') from user_roles ur, roles r where r.id=ur.role_id and ur.user_id=u.id) roles"+
	" FROM users u limit ?,?",limitFrom,limitTo).QueryRows(&arr)
	respO := userGetResponse{}
	respO.Items = arr
	respO.PageCount = pageCount
	jsonData, err := json.Marshal(respO)
	if RestCheckDBPanic(err ,res ,o ) {
		return
	}
	fmt.Fprint(res,string(jsonData))

}

func UserRestApiInsert(res http.ResponseWriter, req *http.Request, _ httprouter.Params){

	if RestCheckAuth(res,req){
		return
	}

	decoder := json.NewDecoder(req.Body)
	var t userUpdateRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	o := orm.NewOrm()
	o.Using("default")
	o.Begin()


	qs := o.QueryTable("users")
	i, _ := qs.PrepareInsert()
	for index, element := range  t.Items {
		id,err := i.Insert(&element)
		element.Id=id

		t.Items[index].Id=id

		fmt.Printf("%d",id)
		if RestCheckDBPanic(err ,res ,o ) {
			return
		}
	}
	i.Close()
	o.Commit()

	j,_ := json.Marshal(t)
	fmt.Fprint(res,string(j))
}


func UserRestApiUpdate(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(req.Body)
	var t userUpdateRequest
	err := decoder.Decode(&t)
	if RestCheckPanic(err ,res ) {
		return
	}

	o := orm.NewOrm()
	o.Using("default")
	o.Begin()

	for _,element := range t.Items {
		fmt.Fprint(res,element.Name);

		o.QueryTable("users").Filter("id", element.Id).Update(orm.Params{
			"name":element.Name,
			"email": element.Email,
		})
	}

	o.Commit()

}


func UserRestApiDelete(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	decoder := json.NewDecoder(req.Body)
	var t userDeleteRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	o := orm.NewOrm()
	o.Using("default")
	o.Begin()
	RowAffected  := 0;

	for _,element := range t.Items {
		ra,_ := o.QueryTable("users").Filter("id", element.Id).Delete()
		RowAffected = int(ra) + RowAffected
	}

	o.Commit()

	t.RowsAffected = RowAffected
	j,_ := json.Marshal(t)
	fmt.Fprint(res,string(j))

}


