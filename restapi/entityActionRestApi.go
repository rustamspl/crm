package restapi
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"encoding/json"
	"github.com/yeldars/crm/utils"
)



func DoEntityAction(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	type entityActionRequest struct {
		EntityId int `json:"entity_id"`
		ActionId int `json:"action_id"`
	}
	type entityActionResponse struct {
		Status  string `json:"Status"`
	}

	contents, err := ioutil.ReadAll(req.Body)
	if RestCheckPanic(err,res){
		return;
	}

	defer req.Body.Close()



	var request entityActionRequest
	err = json.Unmarshal(contents, &request)



	//log.Println(request)
	//log.Println(request.EntityId)



	if request.ActionId == 1 {
		err = utils.BetonReqSendTo1C(request.EntityId)
	}else{
		err = utils.BetonCancelSend(request.EntityId)
	}
	if RestCheckPanic(err,res){
		return;
	}










}