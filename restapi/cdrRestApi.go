package restapi


import "net/http"
import (
	"github.com/julienschmidt/httprouter"
	"log"

	"encoding/json"
"github.com/astaxie/beego/orm"

)


type CDRQueryGetResponse struct {
	RowCount   int `json:"rowCount"`
	Error      string `json:"error"`
	Items [] orm.Params`json:"items"`
}


func Cdr(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//fmt.Fprint(w, "Welcome!\n")

	o := orm.NewOrm()
	o.Using("default")
	log.Print("URA GOI! 2222 TEMA!!!")

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	var t CDRQueryGetResponse
	err := decoder.Decode(&t)
	//log.Println("BEBEBE 1")



	log.Println(t.RowCount)

	for _,value := range t.Items {

		_,errrr := o.Raw("insert into col_calls "+
		"(ext_id,local_ip_v4,caller_id_name,caller_id_number,"+

		"destination_number,context,start_stamp,answer_stamp,end_stamp,duration, "+

		"billsec,hangup_cause,uuid,bleg_uuid, "+
		"accountcode,read_codec,write_codec,sip_hangup_disposition,ani) "+
		"values(?,?,?,?, ?,?,FROM_UNIXTIME(?),FROM_UNIXTIME(?),FROM_UNIXTIME(?),?, ?,?,?,?, ?,?,?,?,?)",
			value["id"],
			value["local_ip_v4"],
			value["caller_id_name"],
			value["caller_id_number"],

			value["destination_number"],
			value["context"],
			//value["start_stamp"],
			//time.Unix(int64(start_stamp),0),
			value["start_stamp"],
			value["answer_stamp"],
			value["end_stamp"],
			value["duration"],
			//time.Now(),

			value["billsec"],
			value["hangup_cause"],
			value["uuid"],
			value["bleg_uuid"],

			value["accountcode"],
			value["read_codec"],
			value["write_codec"],
			value["sip_hangup_disposition"],
			value["ani"],
		).Exec()
		if errrr != nil{
			panic(errrr)
		}


//		ext_id                        serial primary key,
//		local_ip_v4               inet not null,
//		caller_id_name            varchar,
//		caller_id_number          varchar,
//		destination_number        varchar not null,
//		context                   varchar not null,
//		start_stamp               timestamp with time zone not null,
//		answer_stamp              timestamp with time zone,
//		end_stamp                 timestamp with time zone not null,
//		duration                  int not null,
//		billsec                   int not null,
//		hangup_cause              varchar not null,
//		uuid                      uuid not null,
//		bleg_uuid                 uuid,
//		accountcode               varchar,
//		read_codec                varchar,
//		write_codec               varchar,
//		sip_hangup_disposition    varchar,
//		ani                       varchar
//		);

	}
	if err!=nil{
		panic(err)
	}


	o.Raw("update col_calls c set c.caller_account_id="+
	"(select max(ac.account_id) from accountconts ac where ac.cont=concat('7',substr(c.caller_id_number,2,255)) )"+
	"").Exec()

	o.Raw("update col_calls c set c.dest_account_id="+
	"(select max(ac.account_id) from accountconts ac where ac.cont=concat('7',substr(c.destination_number,2,255)) )"+
	"").Exec()

	//log.Println("BEBEBE 2")




}
