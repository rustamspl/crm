package restapi
import (
	"time"
	"github.com/astaxie/beego/orm"
	"log"
)


func getOperIdByCode(code string) int64{
	o := orm.NewOrm()
	o.Using("default")
	result := int64(0)
	err := o.Raw("select id from stage_opers where oper_code=?",code).QueryRow(&result)
	if err!=nil {
		panic(err)
	}
	return result
}
func run( t * DealOpRun){

log.Println("opercode="+t.Op.OperCode)
if t.Op.OperCode == "departure"{
	run_departure( t )
}else if t.Op.OperCode == "prolong"{
	run_prolong( t )
}else if t.Op.OperCode == "refuse"{
	run_refuse( t )
}else if t.Op.OperCode == "notanswered"{
	run_notAnswered( t )
}

}

func run_departure( t * DealOpRun){

	//log.Println(t.Op.Attrs["date_receipt"])
	date_departure, err := time.Parse(
		time.RFC3339,
		t.Op.Attrs["date_departure"])
	CheckPanic(err)

	time_departure, err := time.Parse(
		time.RFC3339,
		t.Op.Attrs["time_departure"])
	CheckPanic(err)

	if err==nil{
		date_departure = time.Date(date_departure.Year(), date_departure.Month(), date_departure.Day(), time_departure.Hour(),
			time_departure.Minute(), time_departure.Second(), time_departure.Nanosecond(), date_departure.Location())
	}
	//log.Println(t1.Hour())

	o := orm.NewOrm()
	o.Using("default")
	//log.Print("deal to stage:")
	//log.Println(t.Op.ToStageId)
	_,err = o.Raw("update deals set x_departure_datetime=?,x_departure_comment=? where id=?",date_departure,t.Op.Attrs["comment"],t.DealId).Exec();
	if err != nil {
		log.Println("error on update deals")
		CheckPanic(err)
	}

	AddJrn(t.JrnId, "datetime_departure",date_departure.Format(time.RFC3339) )
	AddJrn(t.JrnId, "comment",t.Op.Attrs["comment"] )

}


func run_refuse( t * DealOpRun){

	AddJrn(t.JrnId, "comment",t.Op.Attrs["comment"] )
}

func run_prolong( t * DealOpRun){

	//log.Println(t.Op.Attrs["date_receipt"])
	date_prolong, err := time.Parse(
		time.RFC3339,
		t.Op.Attrs["date_prolong"])
	CheckPanic(err)

	time_prolong, err := time.Parse(
		time.RFC3339,
		t.Op.Attrs["time_prolong"])
	CheckPanic(err)

	if err==nil{
		date_prolong = time.Date(date_prolong.Year(), date_prolong.Month(), date_prolong.Day(), time_prolong.Hour(),
			time_prolong.Minute(), time_prolong.Second(), time_prolong.Nanosecond(), date_prolong.Location())
	}
	//log.Println(t1.Hour())

	o := orm.NewOrm()
	o.Using("default")
	_,err = o.Raw("update deals set x_prolong_datetime=?,x_prolong_comment=? where id=?",date_prolong,t.Op.Attrs["comment"],t.DealId).Exec();
	CheckPanic(err)


	AddJrn(t.JrnId, "datetime_prolong",date_prolong.Format(time.RFC3339) )
	AddJrn(t.JrnId, "comment",t.Op.Attrs["comment"] )
}


func run_notAnswered( t * DealOpRun){

	o := orm.NewOrm()
	o.Using("default")
	_,err := o.Raw("update deals set x_prolong_datetime=NOW() + INTERVAL 1 DAY,x_prolong_comment='NOT ANSWERED' where id=?",t.DealId).Exec();
	CheckPanic(err)


	//AddJrn(t.JrnId, "datetime_prolong",date_prolong.Format(time.RFC3339) )
	AddJrn(t.JrnId, "comment","NOT ANSWERED" )

}

func AddJrn(jrnid int64, code,value string ){
	o := orm.NewOrm()
	o.Using("default")

	_,err := o.Raw("insert into stage_op_jrn_dtls (jrn_id,code,value) values( ?,?,?)",jrnid,code, value).Exec()
	CheckPanic(err)
}