package bpms
import (
	"github.com/astaxie/beego/orm"
	"log"
	"errors"
	"strings"
)

//Создание процесса
func CreateInstance(process_id int64) (int64,error){

	o := orm.NewOrm()
	o.Using("default")
	pointId, err := findFirstPoint(process_id)
	rs,err :=  o.Raw("insert into bp_instances (is_finished,process_id,is_terminated,point_id) values (0,?,?,?)",process_id,0,pointId).Exec()
	if err!=nil{
		return int64(0),err
	}
	instanceId, err := rs.LastInsertId()
	err = gotoNextPoint(instanceId)
	if err!=nil{
		return int64(0),err
	}
	err = execInstance(instanceId)
	return instanceId,err
}

func getPointByInstance(instanceId int64) (int64,error) {
	o := orm.NewOrm()
	o.Using("default")
	pointId:=int64(0)
	err := o.Raw("select point_id from bp_instances where id=?",instanceId).QueryRow(&pointId)
	return pointId,err
}

func getNextPointFromExclusiveGateway(instanceId int64 ) (int64,error){
	pt,err := getPointTypeByInstance(instanceId)
	if err!=nil{
		return int64(0),err
	}
	if pt!="exclusivegateway"{
		return int64(0),errors.New("only exclusivegateway can calculate next point")
	}

	type typeCond struct {
		Cond string `json:"cond"`
		Point int64 `json:"point"`
	}

	var conds = []typeCond{}

	o := orm.NewOrm()
	_,err = o.Raw(`select cond,ps2.point_id point from
			  bp_sequence_flows sf,
			  bp_point_sfs ps1,
			  bp_point_sfs ps2,
			  bp_instances i
where ps1.sf_id=sf.id and
      ps1.point_id=i.point_id and
      i.id=? and ps1.is_incoming=0 and
      ps2.is_incoming=1 and
      ps2.sf_id=ps1.sf_id`,instanceId).QueryRows(&conds)

	if err!=nil{
		return int64(0),err
	}

	for _, element := range conds {
		res,err := checkCondString(element.Cond)
		if err!=nil{
			return int64(0),nil
		}
		if res {
			return element.Point,nil
		}
	}

	return 0,errors.New("No Conditions")
}

func checkCondString(cond string) (bool,error){
	if strings.TrimSpace(cond) == ""  {
		return true, nil
	}

	o := orm.NewOrm()
	o.Using("default")
	cnt := 0
	err := o.Raw("select count(1) cnt from dual where "+cond).QueryRow(&cnt)
	return cnt >0, err
}
func getPointNameByInstance(instanceId int64) (string,error) {
	o := orm.NewOrm()
	o.Using("default")
	name:=""
	err := o.Raw("select p.title from bp_instances i,bp_points p where i.point_id=p.id and i.id=?",instanceId).QueryRow(&name)
	return name,err
}


func isLoopPointByInstance(instanceId int64) (bool,error) {
	o := orm.NewOrm()
	o.Using("default")
	isLoop:=int64(0)
	err := o.Raw("select p.is_loop from bp_instances i,bp_points p where i.point_id=p.id and i.id=?",instanceId).QueryRow(&isLoop)
	return (isLoop==1),err
}

func runServiceTask(instanceId int64) error {
	log.Println("SERVICE TASK DONE")
	return nil
}

func ManualExecInstance(instanceId int64) error {
	pt,err := getPointTypeByInstance(instanceId)
	if err!=nil{
		return err
	}

	loop,err := isLoopPointByInstance(instanceId)
	if err!=nil{
		return err
	}

	if (!loop  && pt == "servicetask" ) {
		if (pt != "usertask" && pt != "manualtask") {
			return errors.New("manual exec only for manualTask and userTask or loop serviceTask. break")
		}
	}

	err = gotoNextPoint(instanceId)
	if err!=nil{
		return err
	}

	execInstance(instanceId)
	if err!=nil{
		return err
	}

return nil
}
//Выполняем текущий шаг
func execInstance(instanceId int64) error {
	pt,err := getPointTypeByInstance(instanceId)
	if err!=nil{
		return err
	}

	loop,err := isLoopPointByInstance(instanceId)
	if err!=nil{
		return err
	}

	if loop && pt == "servicetask" {
		log.Println("detect loop service task. break")
		return nil
	}

	if pt == "usertask" || pt == "manualtask" {
		log.Println("detect "+pt+" break")
		return nil
	}

	name,err := getPointNameByInstance (instanceId)
	if err!=nil{
		return err
	}
	log.Println("executing "+pt+ ".... "+name)



	if pt == "servicetask"{
		runServiceTask(instanceId)
	}


		err = gotoNextPoint(instanceId)
		if err!=nil{
			return err
		}
		execInstance(instanceId)
		if err!=nil{
			return err
		}
	return nil
}

func getPointTypeByInstance(instanceId int64) (string,error) {
	o := orm.NewOrm()
	o.Using("default")
	pointType:=""
	error := o.Raw(`select pt.code from bp_instances i,
										 bp_points p,
										 bp_point_types pt
				  where i.id=? and p.id=i.point_id and p.type_id=pt.id`,instanceId).QueryRow(&pointType)
	return pointType,error
}

func getNextPoint(instanceId int64) (int64,error){

	pt,err := getPointTypeByInstance(instanceId)
	if err!=nil{
		return int64(0),err
	}
	if pt=="exclusivegateway"{
		return getNextPointFromExclusiveGateway(instanceId)
	}

	pointId,err := getPointByInstance(instanceId)
	if err!=nil{
		return 0,err
	}
	o := orm.NewOrm()
	o.Using("default")

	nextPointId := int64(0)

	err = o.Raw(`select s2.point_id from
			bp_point_sfs s1,
			bp_point_sfs s2,
			bp_points p1,
			bp_points p2
			where
			s1.point_id=? and
			s1.sf_id=s2.sf_id and
			s1.is_incoming=0 and
			s2.is_incoming=1
			and s1.point_id=p1.id
			and s2.point_id=p2.id
			and p1.is_active=1
			and p2.is_active=1
	`,pointId).QueryRow(&nextPointId)


	return nextPointId,err
}

func setInstancePoint(instanceId,point int64) error{
	o := orm.NewOrm()
	o.Using("default")
	_,err := o.Raw("update bp_instances set point_id=? where id=?",point,instanceId).Exec()

	pt,err := getPointTypeByInstance(instanceId)

	if pt=="endevent"{
		_,err = o.Raw("update bp_instances set is_finished=1 where id=?",instanceId).Exec()
	}

	return err
}
func gotoNextPoint(instanceId int64) error {
	o := orm.NewOrm()
	o.Using("default")
	point,err := getNextPoint(instanceId)
	if err!=nil{
		return err
	}
	return setInstancePoint(instanceId,point)
}
func findFirstPoint(processId int64) (int64,error) {
	o := orm.NewOrm()
	o.Using("default")
	pointId := int64(0)

	err := o.Raw(`select id from bp_points
			where process_id=?
			 and type_id=(select id from bp_point_types where code='startevent')
			 `,processId).QueryRow(&pointId)
	if err!=nil{
		log.Println("findFirstPoint. No next task found "+err.Error())
		return pointId,err
	}

	return pointId,nil



}

func terminateInstance(instanceId int64 ) error {
	o := orm.NewOrm()
	o.Using("default")
	_,err := o.Raw("update bp_instance_tasks set is_terminated = 1 where instance_id=?",instanceId).Exec()
	_,err = o.Raw("update bp_instances set is_terminated = 1 where id=?",instanceId).Exec()
	return err
}
