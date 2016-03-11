package restapi
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"github.com/astaxie/beego/orm"
	"regexp"
	"github.com/yeldars/crm/utils"
)

func ImportAdvancedReferenceRestApi(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	type referenceElement struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	type referenceImportRequest struct {
		Entity string `json:"entity"`
		ReferenceElements []referenceElement `json:"referenceElements"`
	}

	type referenceImportResponse struct {
		UpdateCount int `json:"updateCount"`
		InsertCount int `json:"insertCount"`
		DeleteCount int `json:"deleteCount"`
		SkipCount int `json:"skipCount"`
		ErrorCount int `json:"errorCount"`
		ErrorTexts string `json:"errorTexts"`
	}

	contents, err := ioutil.ReadAll(req.Body)
	log.Println(string(contents))

	defer req.Body.Close()

	var t referenceImportRequest
	err = json.Unmarshal(contents,&t)

	if err != nil {
		RestCheckPanic(err,res)
		return
	}


	var resP referenceImportResponse

	o := orm.NewOrm()
	o.Using("default")

	if !utils.CheckEntity(t.Entity){
		resP.ErrorCount ++;
		resP.ErrorTexts += "Entity Not Found"
	}else {
		for _, element := range t.ReferenceElements {



			sqlCnt := "select count(1) cnt from " + regexp.QuoteMeta(t.Entity) + " where code=?";
			cnt := 0
			err := o.Raw(sqlCnt, element.Code).QueryRow(&cnt)

			if cnt > 0 {
				if t.Entity == "accounts" {
					sql := "update " + regexp.QuoteMeta(t.Entity) + " set title=? where code=?";
					_, err = o.Raw(sql, element.Name, element.Code).Exec()
				}
				if err == nil {
					resP.UpdateCount ++
				}
			}else {
				if t.Entity == "accounts" {
					_, err = o.Raw("insert into " + regexp.QuoteMeta(t.Entity) + " (code,name) values (?,?)", element.Code, element.Name).Exec()
				}

				if err == nil {
					resP.InsertCount ++
				}
			}

			if err != nil {
				resP.ErrorCount ++
				resP.ErrorTexts += err.Error()
			}

		}
	}

	resP.DeleteCount = 0
	resP.SkipCount = 0
	j,_ := json.Marshal(resP)
	fmt.Fprint(res,string(j))

	}

