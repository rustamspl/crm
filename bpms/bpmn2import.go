package bpms

import 	_ "github.com/go-sql-driver/mysql"

import (
	"encoding/xml"
	"log"
	"github.com/astaxie/beego/orm"
	"os"
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


func ImportBPMN2(xmlStr string )error{



	type TypeStartEvent struct {
		Id   string `xml:"id,attr"`
		Outgoing [] string
	}

	type TypeStopEvent struct {
		Id   string `xml:"id,attr"`
		Incoming [] string
	}

	type TypeUserTask struct {
		Id   string `xml:"id,attr"`
		Outgoing [] string `xml:"incoming"`
		Incoming [] string `xml:"outgoing"`
	}

	type TypeProcess struct {
		Id   string `xml:"id,attr"`
		UserTask []TypeUserTask `xml:"userTask"`
		StartEvent []TypeStartEvent `xml:"startEvent"`
		EndEvent []TypeStopEvent `xml:"endEvent"`
	}

	type TypeBPMN2 struct {
		XMLName   xml.Name `xml:"definitions"`
		Process   TypeProcess `xml:"process"`
	}

	v := &TypeBPMN2{}

	err := xml.Unmarshal([]byte(xmlStr), &v)
	if err!=nil{
		panic(err)
	}

	//log.Println(v.Process.UserTask[0].Incoming[0])
	//log.Println(v.Process.StartEvent[0].Id)
	//log.Println(v.Process.Id)
	//return nil



	o := orm.NewOrm()
	o.Using("default")
	cnt := 0
	o.Raw("select count(1) cnt from bp_processes where code=?",v.Process.Id).QueryRow(&cnt)
	if cnt == 0{
		o.Raw("insert into bp_processes (code,title) values (?,?)",v.Process.Id,v.Process.Id).Exec()
	}

	for _, element := range v.Process.StartEvent {

		err := o.Raw("select count(1) cnt from bp_events where code=?",element.Id).QueryRow(&cnt)
		if err!=nil{
			panic(err)
		}
		if cnt == 0{
			_,err := o.Raw("insert into bp_events (code,event_type_id,process_id) values (?,(select id from bp_event_types where code='start'),(select id from bp_processes where code=?))",element.Id,v.Process.Id).Exec()
			if err!=nil{
				panic(err)
			}
		}
	}

	for _, element := range v.Process.EndEvent {
		err := o.Raw("select count(1) cnt from bp_events where code=?",element.Id).QueryRow(&cnt)
		if err!=nil{
			panic(err)
		}
		if cnt == 0{
			_,err := o.Raw("insert into bp_events (code,event_type_id,process_id) values (?,(select id from bp_event_types where code='end'),(select id from bp_processes where code=?))",element.Id,v.Process.Id).Exec()
			if err!=nil{
				panic(err)
			}
		}
	}



	return nil


}