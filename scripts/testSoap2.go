package main
import (
	"log"
"encoding/xml"
	"time"
	"bytes"
	"net/http"
	"io/ioutil"
)

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
	v.Address = "e2f14b8d-cedc-11e5-b841-000c29d408f3"
	v.ArrivalTime = time.Now()
	v.Central = 1
	v.ClientReceive = "e71b65ef-e8ac-11e4-8140-2c41387d88d0"
	v.Construction = "fe8903bb-257e-11e5-a135-000c29272e31"
	v.Contacts = "54abb0ae-e8e7-11e5-a682-000c29c99fbb"
	v.Contragent = "347cf5da-2e61-11e4-b550-2c41387d88d0"
	v.Date = time.Now()
	v.Deal = "0bf4aff5-ef28-11e4-8140-2c41387d88d0"
	v.DeliveryType = ""
	v.FrequencyOfDeliveries = time.Now()
	v.Individuals = "e486c728-e87e-11e5-a682-000c29c99fbb"
	v.Item = "e519d46d-843c-11e5-b4fd-000c29d0ccb0"
	v.KebRequestNumber = "13"
	v.Mobility = "8cb516af-6577-4528-9cdf-bf6a630300d5"
	v.OnCall = 1
	v.OsadkaKonusa = "aea0a7cb-4a2c-11e5-ba04-000c29272e31"
	v.Period = time.Now()
	v.PlanOn = 1
	v.Priority = 1
	v.Quantity = 20
	v.ReqSpecVehicles = 1
	v.RequestStatus = "Создан"
	v.Responsible = ""
	v.ShippingPickup = time.Now()
	v.SpecVehicles = 1
	v.TimeApply = time.Now()
	v.VHodke = 5.5
	v.VehiclesType = "d78d320e-05f2-11e5-b7fe-001f5b7cd426"
	v.Comment = "Hello from Beton CRM"



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

}
