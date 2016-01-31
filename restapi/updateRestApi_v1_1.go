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
	"github.com/yeldars/crm/auth"
	"strconv"
)

type UpdateRequest_v_1_1_Item struct {
	Values [] orm.Params `json:"values"`
	TableName string `json:"table_name"`
	Action string `json:"action"`

}

type UpdateRequest_v_1_1 struct {
	Items [] UpdateRequest_v_1_1_Item `json:"items"`

}

type UpdateResponse_v_1_1_Item struct {
	TableName string `json:"table_name"`
	LastInsertId int64 `json:"last_insert_id"`

}
type UpdateResponse_v_1_1 struct {
	Items [] UpdateResponse_v_1_1_Item `json:"items"`
	Error int64 `json:"error"`
	ErrorText string `json:"error_text"`

}

func UpdateRestApi_v_1_1(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	if RestCheckAuth(res,req){
		return
	}

	decoder := json.NewDecoder(req.Body)
	var t UpdateRequest_v_1_1
	var resP UpdateResponse_v_1_1
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	o := orm.NewOrm()
	o.Using("default")
	o.Begin()

	for items := range t.Items {
		for _,value := range t.Items[items].Values {
			var arr [] interface{}
			updateColumns := ""
			insertColumns := ""
			insertValues := ""



			for fieldName,fieldValue := range value {

				fieldName=regexp.QuoteMeta(fieldName)

				if !strings.HasPrefix(fieldName,"_") {
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

			lastInsertId := int64(0)
			//log.Println("id ="+ element["id"].(string) )
			if t.Items[items].Action =="insert" {
				sql := "insert into " + t.Items[items].TableName + " ( " + insertColumns + " ) values ( "+ insertValues +")"
				log.Println("insert sql="+sql)
				i, err := o.Raw(sql, arr).Exec()
				if RestCheckDBPanic(err ,res ,o ) {
					return
				}
				lastInsertId,err = i.LastInsertId()
				log.Println("sql="+sql)
				if RestCheckDBPanic(err ,res ,o ) {
					return
				}
				resP.Items = append(resP.Items,UpdateResponse_v_1_1_Item{LastInsertId:lastInsertId,TableName:t.Items[items].TableName})

			}else if t.Items[items].Action =="update"{
				sql := "update " + t.Items[items].TableName + " set " + updateColumns + " where id=?"
				log.Println("update sql="+sql)
				_, err := o.Raw(sql, arr, value["id"]).Exec()
				if RestCheckDBPanic(err ,res ,o ) {
					return
				}
			}else if t.Items[items].Action =="delete"{
				sql := "delete from " + t.Items[items].TableName +" where id=?"
				log.Println("delete sql="+sql)
				_, err := o.Raw(sql, value["id"]).Exec()
				if RestCheckDBPanic(err ,res ,o ) {
					return
				}
			}


			///LOGGING

			lid := ""
			if t.Items[items].Action == "insert" {
				lid = strconv.Itoa(int(lastInsertId))
			}else{
				lid = value["id"].(string)
			}
			id,err := o.Raw("insert into table_logs (user_id,table_name,action,pk) values (?,?,?,?)",auth.UserId(req),t.Items[items].TableName,t.Items[items].Action,lid).Exec()

			if RestCheckDBPanic(err ,res ,o ) {
				return
			}
			iid,_ := id.LastInsertId()

			for fieldName,fieldValue := range value {
				fieldName = regexp.QuoteMeta(fieldName)
				if !strings.HasPrefix(fieldName, "_") {
					if fieldValue == nil {
						_, err := o.Raw("insert into table_log_dtls (col,val,log_id) values (?,NULL,?)", fieldName, iid).Exec()
						if RestCheckDBPanic(err, res, o) {
							return
						}
					}  else {
						if fieldName != "id" {
							_, err := o.Raw("insert into table_log_dtls (col,val,log_id) values (?,?,?)", fieldName, fieldValue, iid).Exec()
							if RestCheckDBPanic(err, res, o) {
								return
							}
						}
					}
				}
			}

			///LOGGING
		}
	}
	o.Commit()


	resP.Error = 0
	resP.ErrorText = "OK"
	jsonData, err := json.Marshal(resP)
	fmt.Fprint(res,string(jsonData))

}
