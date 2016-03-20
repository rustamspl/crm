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



func findFirstTask(instanceId int64) (int64,error) {
	o := orm.NewOrm()
	o.Using("default")
	taskId := int64(0)

	err := o.Raw(`select ts.task_id from bp_task_sfs ts where  ts.is_incoming=1 and
	ts.sf_id=(select e.sf_id from bp_events e  where e.process_id=(select process_id from bp_instances where id=?) and
	e.event_type_id=(select id from bp_event_types et where et.code='start'))`,instanceId).QueryRow(&taskId)
	if err!=nil{
		log.Println("findFirstTask. No next task found "+err.Error())
		return taskId,err
	}else{
		_,err := o.Raw("insert into bp_instance_tasks (instance_id,task_id) values (?,?)",instanceId,taskId).Exec()
		return taskId,err
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

	taskId, err := findFirstTask(instanceId)
	if err!=nil{
		err := terminateInstance(instanceId)
		if err!=nil{
			log.Println("InitRun. Error on terminate task "+err.Error())
			return taskId, false, err
		}
		return taskId, true, err
	}else {
		return taskId, false, err
	}
}
