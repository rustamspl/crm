package restapi
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/yeldars/crm/bpms"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
)

func BPMPublish(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	type bpmPublishRequest struct {
		ProcessId int64 `json:"processId"`
	}

	type bpmPublishResponse struct {
		Ok bool `json:"ok"`
		ErrorText string `json:"errorText"`
	}

	var request bpmPublishRequest
	var response bpmPublishResponse

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&request)
	if err!=nil {
		response.Ok = false
		response.ErrorText = err.Error()
		//return
	}else {
		o := orm.NewOrm()
		o.Using("default")
		diagram := ""
		err := o.Raw("select diagram from bp_processes where id=?",request.ProcessId).QueryRow(&diagram)
		//log.Println("diagram="+diagram)
		if err==nil {
			err = bpms.ImportBPMN2(diagram,request.ProcessId)
			if err != nil {
				response.Ok = false
				response.ErrorText = err.Error()
				//return
			}else {
				response.Ok = true
			}
		}
	}
	resP,_ := json.Marshal(response)
	fmt.Fprint(res,string(resP))

}

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
		}else{
			response.Ok = true
		}
	}
	resP,_ := json.Marshal(response)
	fmt.Fprint(res,string(resP))


}

func BPMManualExecInstance(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	type execInstanceRequest struct {
		InstanceId int64 `json:"instanceId"`
	}

	type execInstanceResponse struct {
		Ok bool `json:"ok"`
		ErrorText string `json:"errorText"`
	}
	var request execInstanceRequest
	var response execInstanceResponse

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
		err = bpms.ManualExecInstance(request.InstanceId)
		if err!=nil {
			response.Ok = false
			response.ErrorText = err.Error()
			//return
		}else{
			response.Ok = true
		}
	}
	resP,_ := json.Marshal(response)
	fmt.Fprint(res,string(resP))


}