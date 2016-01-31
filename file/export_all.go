package file
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/astaxie/beego/orm"
	"github.com/tealeg/xlsx"
"strings"
"main/auth"
	"strconv"
)

func ExportAll(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	req.ParseForm()
	o := orm.NewOrm()
	o.Using("default")
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, err := file.AddSheet(req.Form.Get("code"))




	sql := ""
	o.Raw("select sql_text from queries where code=?",req.Form.Get("code")).QueryRow(&sql)
	var p  [] orm.ParamsList

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


	var formArray [] string
	for formName,formValue := range req.Form {
		if strings.HasPrefix(formName,"param") {
			formArray = append(formArray,formValue[0])
		}
	}

	sql = strings.Replace(sql,":user_id",strconv.Itoa(int(auth.UserId(req))),-1)
	_,err = o.Raw(sql,&formArray,&filterArray).ValuesList(&p)
	if err!=nil{
		panic(err)
	}
	for _,element := range p {
		row = sheet.AddRow()
		for _,element2 := range element {
			cell = row.AddCell()
			if element2!=nil {
				cell.Value = element2.(string)
			}
		}
	}
	err = file.Write(res)






}
