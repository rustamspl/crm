package restapi
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/astaxie/beego/orm"
	"fmt"
	"strings"
	"strconv"
	"main/auth"
	"encoding/json"
	"log"
)

type DetailSql struct {
	SqlText string `json:"sql_text"`
	SqlConditionBuildText string `json:"sql_condition_build_text"`
	Code string `json:"code"`

}

func DetailRestApi(res http.ResponseWriter, req *http.Request, _ httprouter.Params){

	if RestCheckAuth(res,req){
		return
	}

	req.ParseForm();

	o := orm.NewOrm()
	o.Using("default")


	id,err := strconv.Atoi(req.Form.Get("id"))
	code := req.Form.Get("code");
	if	RestCheckDBPanic(err,res,o){
		return
	}
	var ws = [] DetailSql{}
	_,err = o.Raw("select ws.* from details w, detail_queries ws where ws.detail_id=w.id and w.code=?",code).QueryRows(&ws)
	CheckPanic(err)
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
	element.SqlText = strings.Replace(element.SqlText,":"+formIndex,formValue[0],-1)
	}


	element.SqlText = strings.Replace(element.SqlText,":user_id",strconv.Itoa(int(auth.UserId(req))),-1)



	if strings.Contains(element.SqlText,"#grant#"){

		_,err := o.Raw("SET group_concat_max_len = 1000000;").Exec()
		if err!=nil{
			panic(err)
		}
			buildCond := ""
			err = o.Raw(element.SqlConditionBuildText).QueryRow(&buildCond)
			if err != nil {
				panic(err)
			}

			element.SqlText = strings.Replace(element.SqlText, "#grant#", buildCond, -1)
		}



	_, err =   o.Raw(element.SqlText,id).Values(&maps)
	//checkErr(err)
		if err!=nil{
			log.Println("sql"+element.SqlText)
			panic(err)
		}else{
			log.Println("sql ok"+element.SqlText)
		}


	fmt.Fprintln(res,"\""+element.Code+"\":")
	j,_ := json.Marshal(maps)
	fmt.Fprint(res,string(j))


	if len(ws)>i {
	fmt.Fprintln(res, ",")
	}

	CheckPanic(err)
	}
	fmt.Fprintln(res,"}")









	}
