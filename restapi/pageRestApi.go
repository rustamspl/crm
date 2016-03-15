package restapi

import "net/http"
import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"strconv"
	"github.com/julienschmidt/httprouter"
	"github.com/yeldars/crm/auth"
	"log"
)


type View struct {
		Name string `json:"name"`
		Templateurl string `json:"templateurl"`
}



type State struct {
	Id          	int64 `json:"id"`
	Title        	string `json:"title"`
	Url 			string `json:"url"`
	Templateurl 	string `json:"templateurl"`
	DbTemplate      int `json:"db_template"`
	Controller 		string `json:"controller"`
	Name        	string `json:"name"`
	Views [] View `json:"views"`
	Files [] string `json:"files"`

}



func PageRestApiGet(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	res.Header().Set("Access-Control-Allow-Origin", "*")

	if RestCheckAuth(res,req){
		return
	}

	o := orm.NewOrm()
	o.Using("default")
	var arr [] State
	_, err := o.Raw("SELECT p.*,pt.controller FROM pages p, page_types pt where p.page_type_id=pt.id").QueryRows(&arr)


	for index,_ := range arr {
		// element is the element from someSlice for where we are
		var z[]View
//		log.Println("DB Template")
//		log.Println(arr[index].DbTemplate)
//		log.Println(arr[index].Id)
		if arr[index].DbTemplate==1{
			arr[index].Templateurl = "../restapi/pagetemplate?id="+strconv.Itoa(int(arr[index].Id))

		}
		z = append(z,View{Templateurl:arr[index].Templateurl, Name:"state"+strconv.Itoa(int(arr[index].Id)) })
		arr[index].Name="state"+strconv.Itoa(int(arr[index].Id))
		arr[index].Views = z

		var arrFiles[] string
		//o.Raw("SELECT j.url FROM `j_s_plugins` j,`page_j_s_plugins` pp where pp.page_id=? and pp.js_id=j.id",arr[index].Id).QueryRows(&arrFiles)
		o.Raw("SELECT j.url FROM `j_s_plugins` j,`page_types` pt,`pages` p,`page_type_js` ptj where p.id=? and p.page_type_id=pt.id and ptj.js_id=j.id and ptj.page_type_id=pt.id",arr[index].Id).QueryRows(&arrFiles)

		arr[index].Files = arrFiles



	}


	jsonData, err := json.Marshal(arr)
	if RestCheckDBPanic(err ,res ,o ) {
		return
	}
	fmt.Fprint(res,string(jsonData))

}

func CheckGrantToPage(userId int64,pageId string) bool{

	o := orm.NewOrm()
	o.Using("default")

	 cnt := 0
	err := o.Raw("select count(1) cnt from pages p,role_pages rp,user_roles ur where ur.user_id=? and ur.role_id=rp.role_id and p.id=rp.page_id and p.id=?",userId,pageId).QueryRow(&cnt)
	CheckPanic(err)
	log.Println(cnt)
	return cnt>0
}

func CheckGrantToPageCode(userId int64,pageCode string) bool{
	o := orm.NewOrm()
	o.Using("default")

	cnt := 0
	err := o.Raw("select count(1) cnt from pages p,role_pages rp,user_roles ur where ur.user_id=? and ur.role_id=rp.role_id and p.id=rp.page_id and p.code=?",userId,pageCode).QueryRow(&cnt)
	CheckPanic(err)
	log.Println(cnt)
	return cnt>0
}

func PageRestApiGetPageTemplate(res http.ResponseWriter, req *http.Request, _ httprouter.Params){

	/*
	if RestCheckAuth(res,req){
		return
	}
	*/

	req.ParseForm()
	o := orm.NewOrm()
	o.Using("default")
	s := "";

	if req.Form.Get("id")!="" {
		err1 := o.Raw("SELECT p.template FROM pages p where p.id=?", req.Form.Get("id")).QueryRow(&s)
		if err1 != nil {
			err1 := o.Raw("SELECT p.template FROM pages p where p.code=?","404").QueryRow(&s)
			RestCheckDBPanic(err1,res,o)
		}

		if !CheckGrantToPage(auth.UserId(req),req.Form.Get("id")){
			err1 := o.Raw("SELECT p.template FROM pages p where p.code=?","403").QueryRow(&s)
			RestCheckDBPanic(err1,res,o)
		}

		fmt.Fprint(res,s)
	}	else if req.Form.Get("code")!="" {
		err2 := o.Raw("SELECT p.template FROM pages p where p.code=?", req.Form.Get("code")).QueryRow(&s)
		if err2 != nil {
			err2 := o.Raw("SELECT p.template FROM pages p where p.code=?","404").QueryRow(&s)
			RestCheckDBPanic(err2,res,o)
		}
		if !CheckGrantToPageCode(auth.UserId(req),req.Form.Get("code")){
			err1 := o.Raw("SELECT p.template FROM pages p where p.code=?","403").QueryRow(&s)
			RestCheckDBPanic(err1,res,o)
		}
		fmt.Fprint(res,s)
	}





}

