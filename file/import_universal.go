package file
import (
	"net/http"
	"fmt"
	"os"
	"github.com/astaxie/beego/orm"
	"log"
	"encoding/json"
	"github.com/tealeg/xlsx"
	"github.com/yeldars/crm/restapi"

	"strconv"
	"strings"
	"github.com/yeldars/crm/macros"
)


const QueryString = "select ea.code attr_code,ea.list_id,dt.code data_type_code,tm.col_num from im_type_maps tm,entity_attrs ea,data_types dt"+
" where ea.id=tm.attr_id and dt.id=ea.data_type_id and tm.type_id=? and ea.entity_link_id is null order by tm.col_num"


func BuildQuery(importType int) (string){

	type tmyMap struct {
		AttrCode string `json:"attr_code"`
		ListId int `json:"list_id"`
		DataTypeCode string `json:"data_type_code"`
		ColNum int `json:"col_num"`
	}
	var myMap = []tmyMap{}
	o := orm.NewOrm()
	o.Using("default")

	TableName := ""
	err := o.Raw("select code from entities e,im_types it where it.entity_id=e.id and it.id=?",importType).QueryRow(&TableName)

	if err!=nil{
		panic(err)
	}
	_,err = o.Raw(QueryString,importType).QueryRows(&myMap)

	if err!=nil{
		panic(err)
	}

	var attrs []string
	var values []string
	for _, row := range myMap {
		log.Println("attrCode="+row.AttrCode)
		attrs = append(attrs,row.AttrCode)
		if row.ListId ==0 {
		values = append(values,"?")
		}else{
			values = append(values,"(select id from list_values lv where lv.value=? and lv.list_id="+strconv.Itoa(row.ListId)+")")
		}
	}

	return "insert into "+TableName+" ("+strings.Join(attrs,",")+") values ("+strings.Join(values,",")+")"
}

func BuildValues(importType int,rows *xlsx.Row) ([]string){

	type tmyMap struct {
		AttrCode string `json:"attr_code"`
		ListId int64 `json:"list_id"`
		DataTypeCode string `json:"data_type_code"`
		ColNum int `json:"col_num"`
	}
	var myMap = []tmyMap{}
	o := orm.NewOrm()
	o.Using("default")
	_,err := o.Raw(QueryString,importType).QueryRows(&myMap)

	var values []string

	if err!=nil{
		log.Println("vata")
		log.Println(err)
	}
	for _, row := range myMap {
		str,_ := rows.Cells[row.ColNum].String()
		values= append(values,str)
	}

	return values
}

func ImportUniversal(w http.ResponseWriter, r *http.Request, fileName string) (error) {

   //fmt.Fprint(w,"{\"result\": \"importok\"}")

	type resultT struct {
		Result string `json:"result"`
		OkCnt   int `json:"ok_cnt"`
		ErrCnt  int `json:"err_cnt"`
		SkipCnt int	`json:"skip_cnt"`
	}

	var result resultT
	file, err := os.Open(fileName)
	if err != nil {
		return  err
	}
	defer file.Close()

	o := orm.NewOrm()
	o.Using("default")



	xlFile, err := xlsx.OpenFile(fileName)
	restapi.CheckPanic(err)
	template,err := strconv.Atoi(r.Form.Get("template"))

	skipRows := 0

	err = o.Raw("select skip_rows from im_types where id=?",template).QueryRow(&skipRows)

	restapi.CheckPanic(err)



	sql := BuildQuery(template)
	log.Println("sql = " + sql)
	for _, sheet := range xlFile.Sheets {
		ri := 0
		for _, row := range sheet.Rows {



			ri ++

			if ri <= skipRows{
				continue
			}


			rSet,err := o.Raw(sql,BuildValues(template,row)).Exec()


			if err==nil{
				result.OkCnt ++
				lid,_ := rSet.LastInsertId()
				macros.RunMacro(template,lid,row)
			}else{
				result.ErrCnt++
				log.Println(err)
			}

		}
	}
	result.Result="ok"
	jsonData, err := json.Marshal(result)
	fmt.Fprint(w,string(jsonData))

	return  err
	//return lines, scanner.Err()
}
