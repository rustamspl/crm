package restapi
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/yeldars/crm/bpms"
	"encoding/json"
	"fmt"
)


func BPMCreateInstance(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	type createInstanceRequest struct {
		ProcessId int64 `json:"processId"`
	}

	type createInstanceResponse struct {
		Ok bool `json:"ok"`
		Instance int64 `json:"instance"`
		ErrorText string `json:"errorText"`
	}
	var request createInstanceRequest
	var response createInstanceResponse

	////Warning need uncomment///
	//if RestCheckAuth(res, req) {
	//	return
	//}
	////Warning need uncomment///


	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&request)
	if err!=nil {
		response.Ok = false
		response.ErrorText = err.Error()
		//return
	}else {
		response.Instance, err = bpms.CreateInstance(request.ProcessId)
		if err!=nil {
			response.Ok = false
			response.ErrorText = err.Error()
			//return
		}
	}
	resP,_ := json.Marshal(response)
	fmt.Fprint(res,string(resP))


}