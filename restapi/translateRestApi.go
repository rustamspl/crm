package restapi

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"main/models"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"main/auth"
)


type translateGetResponse struct {
	ru  (map[string]string)
	kk  (map[string]string)
	en  (map[string]string)
}




func TranslateRestApiGet(res http.ResponseWriter, req *http.Request, _ httprouter.Params){

	o := orm.NewOrm()
	o.Using("default")
	var arr [] models.Translates
	_, err := o.Raw("SELECT * FROM translates").QueryRows(&arr)
	if RestCheckDBPanic(err ,res ,o ) {
		return
	}
	/*respO := translateGetResponse{}
	respO.Items = arr
	respO.PageCount = 10
	jsonData, err := json.Marshal(respO)

	fmt.Fprint(res,string(jsonData))*/

	ru_dict := make (map[string]string)
	kk_dict := make (map[string]string)
	en_dict := make (map[string]string)


	for _,element := range arr {
		// element is the element from someSlice for where we are
		ru_dict[element.Code]=element.Ru;
		kk_dict[element.Code]=element.Kk;
		en_dict[element.Code]=element.En;
	}

	//out := translateGetResponse{ru:ru_dict}

	fmt.Fprint(res,"{\"ru\":")
	jsonData, err := json.Marshal(ru_dict)
	fmt.Fprint(res,string(jsonData))

	fmt.Fprint(res,",\"en\":")
	jsonData, err = json.Marshal(en_dict)
	fmt.Fprint(res,string(jsonData))

	fmt.Fprint(res,",\"kk\":")
	jsonData, err = json.Marshal(kk_dict)
	fmt.Fprint(res,string(jsonData))
	fmt.Fprint(res,", \"lang\":\""+auth.GetLanguage2(req)+"\"}");




}

