package main

import 	_ "github.com/go-sql-driver/mysql"

import (
	"log"
"encoding/xml"
	"time"
	"bytes"
	"net/http"
	"io/ioutil"

	"github.com/astaxie/beego/orm"
	"os"
	"strings"
)

func init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err!=nil{
		panic(err)
	}

	err = orm.RegisterDataBase("default", "mysql", os.Getenv("OPENSHIFT_MYSQL_DB_USERNAME")+":"+os.Getenv("OPENSHIFT_MYSQL_DB_PASSWORD")+"@tcp("+os.Getenv("OPENSHIFT_MYSQL_DB_HOST")+":"+os.Getenv("OPENSHIFT_MYSQL_DB_PORT")+")/"+os.Getenv("OPENSHIFT_APP_NAME")+"?charset=utf8")
	if err!=nil{
		panic(err)
	}else{
		log.Println("ok... openshift_port="+os.Getenv("OPENSHIFT_GO_PORT"))
	}

}


func main(){



	type CreateBRequest struct {
		XMLName   xml.Name `xml:"crm:create"`
		Address string   `xml:"crm:adres"`
		ArrivalTime time.Time `xml:"crm:arrival_time"`
		Central int `xml:"crm:central"`
		ClientReceive string `xml:"crm:client_receive"`
		Construction string `xml:"crm:construction"`
		Contacts string `xml:"crm:contacts"`
		Contragent string `xml:"crm:Contragent"`
		Date time.Time `xml:"crm:date"`
		Deal string `xml:"crm:deal"`
		DeliveryType string `xml:"crm:delivery_type"`
		FrequencyOfDeliveries time.Time `xml:"crm:frequency_of_deliveries"`
		Individuals string `xml:"crm:individuals"`
		Item string `xml:"crm:item"`
		KebRequestNumber string `xml:"crm:KebRequestNumber"`
		Mobility string `xml:"crm:mobility"`
		OsadkaKonusa string `xml:"crm:osadka_konusa"`
		Period time.Time `xml:"crm:period"`
		OnCall int `xml:"crm:on_call"`
		PlanOn int `xml:"crm:plan_on"`
		Priority int `xml:"crm:priority"`
		Quantity float64 `xml:"crm:quantity"`
		ReqSpecVehicles int `xml:"crm:req_spec_vehicles"`
		RequestStatus string `xml:"crm:request_status"`
		Responsible string `xml:"crm:responsible"`
		ShippingPickup time.Time `xml:"crm:shipping_pickup"`
		SpecVehicles int `xml:"crm:spec_vehicles"`
		TimeApply time.Time `xml:"crm:time_apply"`
		VHodke float64 `xml:"crm:v_hodke"`
		VehiclesType string `xml:"crm:vehicles_type"`
		Comment string `xml:"crm:comment"`
	}

	v := &CreateBRequest{}

	o := orm.NewOrm()
	o.Using("default")
	//var vals orm.Params


//	var vals map[string]string
//	vals = make(map[string]string)

	//o.Raw("select 'e2f14b8d-cedc-11e5-b841-000c29d408f3' address ").QueryRow(&vals)
	reqId := 1
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
			(select b.code from bi_individuals b where b.id=r.contact_object_id) individuals,
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
			7.7 v_hodke,
			(select v.code from bi_vehicle_vids v limit 1) vehicles_type,
			r.title comment






			from bi_beton_reqs r where r.id=?`,reqId).QueryRow(&v);
	if err!=nil{
		panic(err)
	}
	log.Println(v.Address)
	log.Println(v.ArrivalTime)
	log.Println(v.Central)
	log.Println(v.ClientReceive)
	log.Println(v.Construction)
	log.Println(v.Contacts)
	log.Println(v.Contragent)
	log.Println(v.Date)
	log.Println(v.Deal)
	log.Println(v.FrequencyOfDeliveries)
	log.Println(v.Individuals)
	log.Println(v.Item)
	log.Println(v.KebRequestNumber)
	log.Println(v.Mobility)
	log.Println(v.OnCall)
	log.Println(v.OsadkaKonusa)
	log.Println(v.Period)
	log.Println(v.PlanOn)
	log.Println(v.Priority)
	log.Println(v.Quantity)
	log.Println(v.ReqSpecVehicles)
	log.Println(v.RequestStatus)
	log.Println(v.Responsible)
	log.Println(v.ShippingPickup)
	log.Println(v.SpecVehicles)
	log.Println(v.TimeApply)
	log.Println(v.VHodke)
	log.Println(v.VehiclesType)
	log.Println(v.Comment)

	//return


	//v.Address = "e2f14b8d-cedc-11e5-b841-000c29d408f3"//vals["address"].(string)
	//v.ArrivalTime = time.Now()
	//v.Central = 1
	//v.ClientReceive = "e71b65ef-e8ac-11e4-8140-2c41387d88d0"
	//v.Construction = "fe8903bb-257e-11e5-a135-000c29272e31"
	//v.Contacts = "54abb0ae-e8e7-11e5-a682-000c29c99fbb"
	//v.Contragent = "347cf5da-2e61-11e4-b550-2c41387d88d0"
	//v.Date = time.Now()
	//v.Deal = "0bf4aff5-ef28-11e4-8140-2c41387d88d0"
	//v.DeliveryType = ""
	//v.FrequencyOfDeliveries = time.Now()
	//v.Individuals = "e486c728-e87e-11e5-a682-000c29c99fbb"
	//v.Item = "e519d46d-843c-11e5-b4fd-000c29d0ccb0"
	//v.KebRequestNumber = "13"
	//v.Mobility = "8cb516af-6577-4528-9cdf-bf6a630300d5"
	//v.OnCall = 1
	//v.OsadkaKonusa = "aea0a7cb-4a2c-11e5-ba04-000c29272e31"

	//v.Period = time.Now()
	//v.PlanOn = 1
	//v.Priority = 1
	//v.Quantity = 20
	//v.ReqSpecVehicles = 1
	//v.RequestStatus = "Создан"
	//v.Responsible = ""
	//v.ShippingPickup = time.Now()
	//v.SpecVehicles = 1
	//v.TimeApply = time.Now()
	//v.VHodke = 5.5
	//v.VehiclesType = "d78d320e-05f2-11e5-b7fe-001f5b7cd426"
	//v.Comment = "Hello from Beton CRM"



	output, err := xml.MarshalIndent(v, "", "\t")
	if err!=nil{
		panic(err)
	}
	log.Println(string(output))


buf := bytes.NewBufferString(
`<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:crm="crmnewrequest">
   <soap:Header/>
   <soap:Body>
   ` + string(output)+
`   </soap:Body>
</soap:Envelope>`	)
	resp, err := http.Post("http://ws_user:123@185.46.152.129:8080/test_keb_1c/ws/ws2.1cws", "application/soap+xml;charset=UTF-8;action=\"crmnewrequest#crmnewrequest:create\"", buf)
	if err!=nil{
		panic(err)
	}
	response, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(response))

	respTxt  := string(response)

	respTxt   = strings.Replace(respTxt,"m:","",-1)
	respTxt   = strings.Replace(respTxt,"soap:","",-1)



	log.Println(respTxt)

	type Result struct {
		XMLName xml.Name `xml:"Envelope"`
		Return string `xml:"Body>createResponse>return"`

		}
	var respStr Result

	err = xml.Unmarshal([]byte(respTxt), &respStr)
	log.Println(respStr)
	arrRes := strings.Split(respStr.Return,"|")
	log.Println(arrRes[0])
	log.Println(arrRes[1])

	_,err = o.Raw("update bi_beton_reqs set num=?,code=? where id=?",arrRes[0],arrRes[1],reqId).Exec()
	if err!=nil{
		panic(err)
	}


}
