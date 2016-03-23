package restapi
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
"github.com/yeldars/crm/utils"
	"log"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
)



func Beton1CNewClaimRestApi(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	type typeRequest struct {
		InvoiceId string `json:"invoiceId"`
		Text string `json:"text"`
	}
	type typeResponse struct {
		ClaimId string `json:"claimId"`
	}

	contents, err := ioutil.ReadAll(req.Body)
	if RestCheckPanic(err, res) {
		return;
	}

	defer req.Body.Close()

	var request typeRequest
	err = json.Unmarshal(contents, &request)

	//log.Println(request)
	log.Println("############")
	log.Println(request.InvoiceId)

	var response = typeResponse{}

	response.ClaimId, err = utils.BetonInvoiceSendClaimTo1C(request.InvoiceId,request.Text)

	o := orm.NewOrm()
	o.Using("default")
	_,err = o.Raw("insert into bi_beton_invoice_claims (code,invoice_id,title,dscr) values (?,?,?,?)",response.ClaimId, request.InvoiceId,request.Text,request.Text).Exec()

	if RestCheckDBPanic(err, res,o) {
		return;
	}
	j,_ := json.Marshal(response)
	fmt.Fprint(res,string(j))

}




func Beton1CCloseInvoiceRestApi(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	type typeRequest struct {
		InvoiceId string `json:"invoiceId"`
		Answer string `json:"answer"`
	}
	type typeResponse struct {
		Answer bool `json:"answer"`
	}

	contents, err := ioutil.ReadAll(req.Body)
	if RestCheckPanic(err, res) {
		return;
	}

	defer req.Body.Close()

	var request typeRequest
	err = json.Unmarshal(contents, &request)

	//log.Println(request)
	log.Println("############")
	log.Println(request.InvoiceId)

	var response = typeResponse{}

	response.Answer, err = utils.BetonInvoicesSendCloseTo1C(request.InvoiceId)

	o := orm.NewOrm()
	o.Using("default")
	_,err = o.Raw("update bi_invoices set status_id=(select id from bi_invoice_statuses where code='closed') where id=?",request.InvoiceId).Exec()

	if RestCheckDBPanic(err, res,o) {
		return;
	}
	j,_ := json.Marshal(response)
	fmt.Fprint(res,string(j))

}






