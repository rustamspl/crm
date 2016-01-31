package restapi

import "net/http"
import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"main/auth"
)


type SubMenu struct {
	Id          	int64 `json:"id"`
	IdHi          	int64 `json:"id_hi"`
	Title        	string `json:"title"`
	TitleEn        	string `json:"title_en"`
	TitleRu        	string `json:"title_ru"`
	TitleKk        	string `json:"title_kk"`
	Url 			string `json:"url"`
	Icon 			string `json:"icon"`
	Code	 		string `json:"code"`
}



type Menu struct {
	Id          	int64 `json:"id"`
	IdHi          	int64 `json:"id_hi"`
	Title        	string `json:"title"`
	TitleEn        	string `json:"title_en"`
	TitleRu        	string `json:"title_ru"`
	TitleKk        	string `json:"title_kk"`
	Url 			string `json:"url"`
	Icon 			string `json:"icon"`
	Code	 		string `json:"code"`
	CntChild       int64   `json:"cnt_child"`
	Items [] SubMenu `json:"items"`
}



func MenuRestApiGetTree(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	if RestCheckAuth(res,req){
		return
	}

	o := orm.NewOrm()
	o.Using("default")
	var arr [] Menu
	var subArr [] SubMenu
	_, err := o.Raw("SELECT m.*"+
	",(select en from translates t where t.code=m.title limit 1) title_en "+
	",(select ru from translates t where t.code=m.title limit 1) title_ru "+
	",(select kk from translates t where t.code=m.title limit 1) title_kk "+
	",(select count(1) from menus mmm where mmm.id_hi=m.id) cnt_child "+
	" FROM menus m where m.id_hi is null and m.active=? "+
	"and exists (select 1 from role_menus rm,user_roles ur where rm.role_id=ur.role_id and ur.user_id=? and m.id=rm.menu_id) order by m.position",1,auth.UserId(req)).QueryRows(&arr)

	for index,element := range arr {
		o.Raw("SELECT m.*"+
		",(select en from translates t where t.code=m.title limit 1) title_en "+
		",(select ru from translates t where t.code=m.title limit 1) title_ru "+
		",(select kk from translates t where t.code=m.title limit 1) title_kk "+
		" FROM menus m where m.id_hi=? and m.active=? "+
		"and exists (select 1 from role_menus rm,user_roles ur where rm.role_id=ur.role_id and ur.user_id=? and m.id=rm.menu_id) order by m.position",element.Id,1,auth.UserId(req)).QueryRows(&subArr)
		arr[index].Items = subArr
	}

	jsonData, err := json.Marshal(arr)
	CheckPanic(err)
	fmt.Fprint(res,string(jsonData))

}



