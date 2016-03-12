package restapi
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"github.com/astaxie/beego/orm"
	"github.com/yeldars/crm/utils"
)

func ImportAdvancedReferenceRestApi(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {




	type referenceImportEntitiesRequest struct {
		Entities []orm.Params `json:"entities"`
	}
	type referenceImportResponse struct {
		UpdateCount int `json:"updateCount"`
		InsertCount int `json:"insertCount"`
		DeleteCount int `json:"deleteCount"`
		SkipCount   int `json:"skipCount"`
		ErrorCount  int `json:"errorCount"`
		ErrorTexts  string `json:"errorTexts"`
	}

	contents, err := ioutil.ReadAll(req.Body)
	log.Println(string(contents))

	defer req.Body.Close()

	var t referenceImportEntitiesRequest
	err = json.Unmarshal(contents, &t)

	if err != nil {
		RestCheckPanic(err, res)
		return
	}


	var resP referenceImportResponse

	o := orm.NewOrm()
	o.Using("default")

	for _, element := range t.Entities {
			entity := element["entity"].(string)
			if utils.CheckTableRegexp(entity) {
				sqlCnt := "select count(1) cnt from " + entity + " where code=?";
				cnt := 0
				err := o.Raw(sqlCnt,element["code"]).QueryRow(&cnt)
				if cnt > 0 {
					_, err = AdvancedImportCaseUpdate(entity, o, element)
					if err == nil {
						resP.UpdateCount ++
					}
				}else {
					_, err = AdvancedImportCaseInsert(entity, o, element)
					if err == nil {
						resP.InsertCount ++
					}
				}
				if err != nil {
					resP.ErrorCount ++
					resP.ErrorTexts += err.Error()+" in "+entity+ "\n"
				}
			}else{
				resP.ErrorCount ++
				resP.ErrorTexts += "Invalid tablename \""+entity+"\"\n"
			}
	}
	resP.DeleteCount = 0
	resP.SkipCount = 0
	j, _ := json.Marshal(resP)
	fmt.Fprint(res, string(j))
	log.Println(string(j))

}

