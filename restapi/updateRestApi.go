package restapi
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"strings"
	"regexp"
	"log"
	"fmt"
	"strconv"
)

type UpdateRequest struct {
	Items [] orm.Params `json:"items"`
	TableName string `json:"table_name"`
}

func UpdateRestApi(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	if RestCheckAuth(res,req){
		return
	}

	decoder := json.NewDecoder(req.Body)
	var t UpdateRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	o := orm.NewOrm()
	o.Using("default")
	o.Begin()


	for _,element := range t.Items {
		var arr [] interface{}
		updateColumns := ""
		insertColumns := ""
		insertValues := ""



		for fieldName,fieldValue := range element {

			fieldName=regexp.QuoteMeta(fieldName)

			if fieldName=="_table_name_"{
				t.TableName=regexp.QuoteMeta(fieldValue.(string))
			}else if !strings.HasPrefix(fieldName,"_") {
				if fieldValue==nil{
					updateColumns = updateColumns + " `" + fieldName + "`=NULL,"
					insertColumns = insertColumns + fieldName + ","
					insertValues = insertValues + "NULL,"
				}  else {
					if fieldName!="id" {
						updateColumns = updateColumns + " `" + fieldName + "`=?,"
						insertColumns = insertColumns + fieldName + ","
						insertValues = insertValues + "?,"
						arr = append(arr, fieldValue)
					}


				}
			}
		}

		updateColumns = strings.TrimRight(updateColumns,",")
		insertColumns = strings.TrimRight(insertColumns,",")
		insertValues = strings.TrimRight(insertValues,",")
		//log.Println("id ="+ element["id"].(string) )
		if element["id"]==nil || element["id"]=="0" {
			sql := "insert into " + t.TableName + " ( " + insertColumns + " ) values ( "+ insertValues +")"
			log.Println("insert sql="+sql)
			i, err := o.Raw(sql, arr).Exec()
			if RestCheckDBPanic(err ,res ,o ) {
				return
			}
			lastInsertId,err := i.LastInsertId()
			log.Println("sql="+sql)
			if RestCheckDBPanic(err ,res ,o ) {
				return
			}
			if err!=nil{
				fmt.Fprint(res, "{\"error\":\"1\"}")
				log.Println(err)
			}else {
				fmt.Fprint(res, "{\"error\":\"0\", \"items\":[{ \"id\":" + strconv.Itoa(int(lastInsertId)) + "}] }")
			}
		}else{
			sql := "update " + t.TableName + " set " + updateColumns + " where id=?"
			log.Println("update sql="+sql)
			_, err := o.Raw(sql, arr, element["id"]).Exec()
			if err!=nil {
				fmt.Fprint(res, "{\"error\":\"1\"}")
				log.Println(err)
			}else {
				fmt.Fprint(res, "{\"error\":\"0\"}")
			}
		}

		//o.QueryTable(t.TableName).Filter("id", element["id"]).Update(element)
	}

	o.Commit()

}
