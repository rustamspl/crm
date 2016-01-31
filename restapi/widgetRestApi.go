package restapi
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
"github.com/astaxie/beego/orm"
	"fmt"
	"encoding/json"
	"strings"
	"strconv"
	"main/auth"
)

type WidgetSql struct {
	Sqltext string `json:"sqltext"`
	Code string `json:"code"`
}


func WidgetRestApiGet(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	if RestCheckAuth(res,req){
		return
	}

	req.ParseForm();

	o := orm.NewOrm()
	o.Using("default")


	var ws = [] WidgetSql{}
	_,err := o.Raw("select ws.* from widgets w, widget_sqls ws where ws.widget_id=w.id and w.code=?",req.Form.Get("code")).QueryRows(&ws)
	if RestCheckDBPanic(err ,res ,o ) {
		return
	}
	//fmt.Fprint(res,&ws)



	i := 0
	//fmt.Print(len(ws))
	fmt.Fprintln(res,"{")
	for _,element := range ws {

		i++
		var maps []orm.Params
		for formIndex,formValue := range req.Form {

			formValue[0] = strings.Replace(formValue[0],"\"","",-1)
			formValue[0] = strings.Replace(formValue[0],"'","",-1)
			formValue[0] = strings.Replace(formValue[0],"/*","",-1)
			formValue[0] = strings.Replace(formValue[0],"--","",-1)
			element.Sqltext = strings.Replace(element.Sqltext,":"+formIndex,formValue[0],-1)
		}


		element.Sqltext = strings.Replace(element.Sqltext,":user_id",strconv.Itoa(int(auth.UserId(req))),-1)
		_, err :=   o.Raw(element.Sqltext).Values(&maps)
		//checkErr(err)

		fmt.Fprintln(res,"\""+element.Code+"\":")
		j,_ := json.Marshal(maps)
		fmt.Fprint(res,string(j))


		if len(ws)>i {
			fmt.Fprintln(res, ",")
		}

		if RestCheckDBPanic(err ,res ,o ) {
			return
		}
	}
	fmt.Fprintln(res,"}")









}



func WidgetRestApiGetWidgetTemplate(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	if RestCheckAuth(res,req){
		return
	}

	req.ParseForm()
	o := orm.NewOrm()
	o.Using("default")
	s := "";
	err := o.Raw("SELECT p.template FROM widgets p where p.id=?",req.Form.Get("id")).QueryRow(&s)
	RestCheckDBPanic(err,res,o)
	fmt.Fprint(res,s)

}




