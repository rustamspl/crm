package bpms
import (
	"github.com/astaxie/beego/orm"
	"log"
)

//Создание процесса
func CreateInstance(process_id int64) (int64,error){

	o := orm.NewOrm()
	o.Using("default")
	rs,err :=  o.Raw("insert into bp_instances (process_id,is_terminated) values (?,?)",process_id,0).Exec()
	if err!=nil{
		o.Rollback()
		return int64(0),err
	}
	instanceId, err := rs.LastInsertId()
	_,_,err = InitRun(instanceId)
	if err!=nil {
		o.Rollback()
		return int64(0), err
	}
	o.Commit()
	return instanceId,err
}

//Запуск таска
func RunTask(instanceId int64 ) (int64, bool, error) {
	return int64(0),false,nil
}



func findFirstPoint(instanceId int64) (int64,error) {
	o := orm.NewOrm()
	o.Using("default")
	pointId := int64(0)

	err := o.Raw(`select id from bp_points
			where process_id=(select process_id from bp_instances where id=?)
			 and type_id=(select id from bp_point_types where code='startevent')
			 `,instanceId).QueryRow(&pointId)
	if err!=nil{
		log.Println("findFirstPoint. No next task found "+err.Error())
		return pointId,err
	}else{
		_,err := o.Raw("insert into bp_instance_points (instance_id,point_id,is_finished) values (?,?,0)",instanceId,pointId).Exec()
		return pointId,err
	}



}

func terminateInstance(instanceId int64 ) error {
	o := orm.NewOrm()
	o.Using("default")
	_,err := o.Raw("update bp_instance_tasks set is_terminated = 1 where instance_id=?",instanceId).Exec()
	_,err = o.Raw("update bp_instances set is_terminated = 1 where id=?",instanceId).Exec()
	return err
}
func InitRun(instanceId int64 ) (int64, bool, error) {

	o := orm.NewOrm()
	o.Using("default")

	pointId, err := findFirstPoint(instanceId)
	if err!=nil{
		err2 := terminateInstance(instanceId)
		if err2!=nil{
			log.Println("InitRun. Error on terminate task "+err.Error())
			return pointId, false, err2
		}
		return pointId, true, err
	}else {
		return pointId, false, err
	}
}
