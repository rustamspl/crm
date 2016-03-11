package restapi
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
)

func InvoiceCreateRestApi(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	type InvoiceCreateRestApiRequestStruct struct {
		RequestNumber string `json:"requestNumber"`
		InvoiceNumber string `json:"invoiceNumber"`
	}

	type InvoiceCreateRestApiResponseStruct struct {
		ResponseOk bool `json:"responseOk"`
		ResponseStatusCode string `json:"responseStatusCode"`
		ResponseText string `json:"responseText"`
		ResponseRefer string `json:"responseRefer"`
	}


	contents, err := ioutil.ReadAll(req.Body)
	log.Println(string(contents))

	defer req.Body.Close()


	var t InvoiceCreateRestApiRequestStruct
	err = json.Unmarshal(contents,&t)

	if err != nil {
		RestCheckPanic(err,res)
		return
	}

	var resP InvoiceCreateRestApiResponseStruct

	resP.ResponseOk = true
	resP.ResponseStatusCode = "OK"
	resP.ResponseText = "Накладная принята" + t.InvoiceNumber
	resP.ResponseRefer = "48454FF45454AA"
	j,_ := json.Marshal(resP)
	fmt.Fprint(res,string(j))

}