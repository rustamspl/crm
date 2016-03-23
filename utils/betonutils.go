package utils


import (
"log"
"encoding/xml"
"time"
"bytes"
"net/http"
"io/ioutil"
"github.com/astaxie/beego/orm"
"strings"
)

func BetonCancelSend(reqId string) error {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("update bi_beton_reqs set status_id=1 where id=?", reqId).Exec()
	return err
}
func BetonReqSendTo1C(reqId string) error {

	type CreateBRequest struct {
		XMLName               xml.Name `xml:"crm:create"`
		Address               string   `xml:"crm:adres"`
		ArrivalTime           time.Time `xml:"crm:arrival_time"`
		Central               int `xml:"crm:central"`
		ClientReceive         string `xml:"crm:client_receive"`
		Construction          string `xml:"crm:construction"`
		Contacts              string `xml:"crm:contacts"`
		Contragent            string `xml:"crm:Contragent"`
		Date                  time.Time `xml:"crm:date"`
		Deal                  string `xml:"crm:deal"`
		DeliveryType          string `xml:"crm:delivery_type"`
		FrequencyOfDeliveries time.Time `xml:"crm:frequency_of_deliveries"`
		Individuals           string `xml:"crm:individuals"`
		Item                  string `xml:"crm:item"`
		KebRequestNumber      string `xml:"crm:KebRequestNumber"`
		Mobility              string `xml:"crm:mobility"`
		OsadkaKonusa          string `xml:"crm:osadka_konusa"`
		Period                time.Time `xml:"crm:period"`
		OnCall                int `xml:"crm:on_call"`
		PlanOn                int `xml:"crm:plan_on"`
		Priority              int `xml:"crm:priority"`
		Quantity              float64 `xml:"crm:quantity"`
		ReqSpecVehicles       int `xml:"crm:req_spec_vehicles"`
		RequestStatus         string `xml:"crm:request_status"`
		Responsible           string `xml:"crm:responsible"`
		ShippingPickup        time.Time `xml:"crm:shipping_pickup"`
		SpecVehicles          int `xml:"crm:spec_vehicles"`
		TimeApply             time.Time `xml:"crm:time_apply"`
		VHodke                string `xml:"crm:v_hodke"`
		VehiclesType          string `xml:"crm:vehicles_type"`
		Comment               string `xml:"crm:comment"`
	}
	v := &CreateBRequest{}
	o := orm.NewOrm()
	o.Using("default")
	err := o.Raw(`
	select (select a.code from bi_addresses a where a.id=r.address_id) address,
			Now() arrival_time,
			r.is_central central,
			(select a.code from accounts a where a.id=r.consignee_id) client_receive,
			(select c.code from bi_constructions c where c.id=r.construction_id) construction,
			(select c.code from contacts c where c.id=r.contact_id) contacts,
			(select a.code from accounts a where a.id=r.account_id) contragent,
			created_at date,
			(select b.code from bi_deals b where b.id=r.deal_id) deal,
			Now() frequency_of_deliveries,
			(select b.title from bi_individuals b where b.id=r.contact_object_id) individuals,
			(select b.code from bi_nomens b where b.id=r.nomen_id) item,
			13 keb_request_number,
			(select b.code from bi_mobilities b where b.id=r.mobility_id) mobility,
			r.by_call on_call,
			(select b.code from bi_mobilities b where b.id=r.mobility_id) osadka_konusa,
			r.period,
			1 plan_on,
			1 priority,
			r.quantity,
			1 req_spec_vehicles,
			(select b.title from bi_beton_req_statuses b where b.id=r.status_id) request_status,
			'' responsible,
			now() ShippingPickup,
			1 spec_vehicles,
			now() time_apply,
			r.v_hodke,
			(select v.code from bi_vehicle_vids v limit 1) vehicles_type,
			r.title comment
			from bi_beton_reqs r where r.id=?`, reqId).QueryRow(&v);
	if err != nil {
		return err
	}

	output, err := xml.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}
	log.Println(string(output))


	buf := bytes.NewBufferString(
		`<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:crm="crmnewrequest">
   <soap:Header/>
   <soap:Body>
   ` + string(output) +
		`   </soap:Body>
</soap:Envelope>`)
	resp, err := http.Post("http://ws_user:123@185.46.152.129:8080/test_keb_1c/ws/ws2.1cws", "application/soap+xml;charset=UTF-8;action=\"crmnewrequest#crmnewrequest:create\"", buf)
	if err != nil {
		return err
	}
	response, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(response))
	respTxt := string(response)
	respTxt = strings.Replace(respTxt, "m:", "", -1)
	respTxt = strings.Replace(respTxt, "soap:", "", -1)
	log.Println(respTxt)

	type Result struct {
		XMLName xml.Name `xml:"Envelope"`
		Return  string `xml:"Body>createResponse>return"`
	}

	var respStr Result

	err = xml.Unmarshal([]byte(respTxt), &respStr)
	log.Println(respStr)
	arrRes := strings.Split(respStr.Return, "|")
	log.Println(arrRes[0])
	log.Println(arrRes[1])

	_, err = o.Raw("update bi_beton_reqs set status_id=2, num=?,code=? where id=?", arrRes[0], arrRes[1], reqId).Exec()
	if err != nil {
		return err
	}

	return err

}



func BetonInvoiceSendClaimTo1C(invoiceId,text string) (string,error) {

	type CreateBRequest struct {
		XMLName               xml.Name `xml:"crm:newclaim"`
		Ttnguid               string   `xml:"crm:ttnguid"`
		Text           time.Time `xml:"crm:text"`
		Claimid               string `xml:"crm:claimid"`

	}
	v := &CreateBRequest{}
	o := orm.NewOrm()
	o.Using("default")
	err := o.Raw(`
	select  r.code as ttnguid,? as text,'' as claimid
	from bi_beton_invoices r where r.id=?`, text,invoiceId).QueryRow(&v);
	if err != nil {
		return "",err
	}

	output, err := xml.MarshalIndent(v, "", "\t")
	if err != nil {
		return "",err
	}
	log.Println(string(output))


	buf := bytes.NewBufferString(
		`<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:crm="crmnewrequest">
   <soap:Header/>
   <soap:Body>
   ` + string(output) +
		`   </soap:Body>
</soap:Envelope>`)
	resp, err := http.Post("http://ws_user:123@185.46.152.129:8080/test_keb_1c/ws/ws2.1cws", "application/soap+xml;charset=UTF-8;action=\"crmnewrequest#crmnewrequest:newclaim\"", buf)
	if err != nil {
		return "",err
	}
	response, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(response))
	respTxt := string(response)
	respTxt = strings.Replace(respTxt, "m:", "", -1)
	respTxt = strings.Replace(respTxt, "soap:", "", -1)
	log.Println(respTxt)

	type Result struct {
		XMLName xml.Name `xml:"Envelope"`
		Return  string `xml:"Body>newclaimResponse>return"`
	}

	var respStr Result

	err = xml.Unmarshal([]byte(respTxt), &respStr)
	log.Println(respStr)
	log.Println(respStr.Return)

	return respStr.Return,err

}


func BetonInvoicesSendCloseTo1C(invoiceId string) (bool,error) {

	type CreateBRequest struct {
		XMLName               xml.Name `xml:"crm:newclaim"`
		Ttnguid               string   `xml:"crm:ttnguid"`
		Answer           time.Time `xml:"crm:answer"`
	}
	v := &CreateBRequest{}
	o := orm.NewOrm()
	o.Using("default")
	err := o.Raw(`
	select  r.code as ttnguid,'' as answer
	from bi_beton_invoices r where r.id=?`, invoiceId).QueryRow(&v);
	if err != nil {
		return false,err
	}

	output, err := xml.MarshalIndent(v, "", "\t")
	if err != nil {
		return false,err
	}
	log.Println(string(output))


	buf := bytes.NewBufferString(
		`<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:crm="crmnewrequest">
   <soap:Header/>
   <soap:Body>
   ` + string(output) +
		`   </soap:Body>
</soap:Envelope>`)
	resp, err := http.Post("http://ws_user:123@185.46.152.129:8080/test_keb_1c/ws/ws2.1cws", "application/soap+xml;charset=UTF-8;action=\"crmnewrequest#crmnewrequest:ttnclosed\"", buf)
	if err != nil {
		return false,err
	}
	response, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(response))
	respTxt := string(response)
	respTxt = strings.Replace(respTxt, "m:", "", -1)
	respTxt = strings.Replace(respTxt, "soap:", "", -1)
	log.Println(respTxt)

	type Result struct {
		XMLName xml.Name `xml:"Envelope"`
		Return  string `xml:"Body>ttnclosedResponse>return"`
	}

	var respStr Result

	err = xml.Unmarshal([]byte(respTxt), &respStr)
	log.Println(respStr)
	log.Println(respStr.Return)

	return respStr.Return=="Ok",err

}