package restapi
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/astaxie/beego/orm"
	"strconv"
	"fmt"
	"encoding/json"
	"strings"
	"github.com/yeldars/crm/auth"
	"log"
)

type queryGetResponse struct {
	PageCount   int `json:"pageCount"`
	Error      string `json:"error"`
	Items [] orm.Params`json:"items"`
}

func FilterBuild(sql string, filterArray []string,req *http.Request,){

}
func QueryRestApiGet(res http.ResponseWriter, req *http.Request, _ httprouter.Params){

	if RestCheckAuth(res,req){
		return
	}
	o := orm.NewOrm()
	o.Using("default")

	var arr [] orm.Params

	req.ParseForm()
	limitFrom := "1"
	limitTo := "5"

	iPerPage,_ := strconv.ParseInt(req.Form.Get("perpage"),10,32);
	iPage,_ := strconv.ParseInt(req.Form.Get("page"),10,32);
	code := req.Form.Get("code")

	limitFrom = strconv.Itoa(int(iPerPage*(iPage-1)))
	//if limitFrom==0{

	//}
	//fmt.Println(limitFrom)
	limitTo = req.Form.Get("perpage");

	pageCount := 0

	sql := ""
	o.Raw("select sql_text from queries where code=?",code).QueryRow(&sql)
	//log.Println("sql="+sql)
	sql = strings.Replace(sql,":user_id",strconv.Itoa(int(auth.UserId(req))),-1)

	var formArray [] string
	var filterArray [] string

	if strings.Contains(sql,"%filter%") {
		sql_filter := " where 1 = 1";


		for formName,formValue := range req.Form {
			if strings.HasPrefix(formName,"flt$") {
				//formArray = append(formArray,formValue[0])
				filterArray = append(filterArray, "%"+formValue[0]+"%")
				sql_filter += " and "+strings.Replace(formName,"flt$","",-1)+" like ? "
			}

		}

		sql = strings.Replace(sql,"%filter%",sql_filter,-1)

	}


	for formName,formValue := range req.Form {
		if strings.HasPrefix(formName,"param") {
			formArray = append(formArray,formValue[0])
		}
	}


	if limitTo==""{ //No pagination

		_, err := o.Raw(sql,formArray,filterArray).Values(&arr)
		if RestCheckDBPanic(err ,res ,o ) {
			log.Println(sql)
			return
		}

	}else {
		err := o.Raw("SELECT ceil(count(1)/?) FROM (" + sql + ") alldata", limitTo,formArray,filterArray).QueryRow(&pageCount)
		if RestCheckDBPanic(err ,res ,o ) {
			log.Println(sql)
			return
		}

		_, err = o.Raw(sql + " limit ?,?", formArray, filterArray, limitFrom, limitTo).Values(&arr)
	}
	respO := queryGetResponse{}
	respO.Items = arr
	respO.Error = "0"
	respO.PageCount = pageCount
	jsonData, err := json.Marshal(respO)
	if RestCheckDBPanic(err ,res ,o ) {
		return
	}
	fmt.Fprint(res,string(jsonData))
	}
