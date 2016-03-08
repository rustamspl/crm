package restapi
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/yeldars/crm/utils"
	"encoding/json"
	"fmt"
)

func GenerateDDL(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	type TDDLGenerateRequest struct {
		EntityCode string `json:"entityCode"`
	}

	type TDDLGenerateResponse struct {
		Ok bool `json:"ok"`
		ErrorText string `json:"errorText"`
	}
	var request TDDLGenerateRequest
	var response TDDLGenerateResponse
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&request)
	if err!=nil {
		response.Ok = false
	}	else {
		err = utils.GenerateDDL(request.EntityCode)
		response.Ok = err==nil
		if err!=nil {
			response.ErrorText = err.Error()
		}

	}

	resP,_ := json.Marshal(response)
	fmt.Fprint(res,string(resP))
}
