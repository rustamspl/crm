package bpms



import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
	"log"
)


type TypeSequenceFlow struct {
	Id   string `xml:"id,attr"`
	Name   string `xml:"name,attr"`
	SourceRef   string `xml:"sourceRef,attr"`
	TargetRef   string `xml:"targetRef,attr"`

}
type TypeStartEvent struct {
	Id   string `xml:"id,attr"`
	Name   string `xml:"name,attr"`
	Outgoing [] string

}

type TypeStopEvent struct {
	Id   string `xml:"id,attr"`
	Name   string `xml:"name,attr"`
	Incoming [] string
}

type TypeUserTask struct {
	Id   string `xml:"id,attr"`
	Name   string `xml:"name,attr"`
	Outgoing [] string `xml:"incoming"`
	Incoming [] string `xml:"outgoing"`
}

type TypeServiceTask struct {
	Id   string `xml:"id,attr"`
	Name   string `xml:"name,attr"`
	Outgoing [] string `xml:"incoming"`
	Incoming [] string `xml:"outgoing"`
	StandardLoopCharacteristics [] string `xml:"standardLoopCharacteristics"`

}

type TypeScriptTask struct {
	Id   string `xml:"id,attr"`
	Name   string `xml:"name,attr"`
	Outgoing [] string `xml:"incoming"`
	Incoming [] string `xml:"outgoing"`
}

type TypeManualTask struct {
	Id   string `xml:"id,attr"`
	Name   string `xml:"name,attr"`
	Outgoing [] string `xml:"incoming"`
	Incoming [] string `xml:"outgoing"`
}


type TypeExclusiveGateway struct {
	Id   string `xml:"id,attr"`
	Name   string `xml:"name,attr"`
	Outgoing [] string `xml:"incoming"`
	Incoming [] string `xml:"outgoing"`
}

type TypeProcess struct {
	Id   string `xml:"id,attr"`
	UserTask []TypeUserTask `xml:"userTask"`
	ScriptTask []TypeScriptTask `xml:"scriptTask"`
	ManualTask []TypeManualTask `xml:"manualTask"`
	ServiceTask []TypeServiceTask `xml:"serviceTask"`
	StartEvent []TypeStartEvent `xml:"startEvent"`
	EndEvent []TypeStopEvent `xml:"endEvent"`
	SequenceFlow []TypeSequenceFlow `xml:"sequenceFlow"`
	ExclusiveGateway  []TypeExclusiveGateway  `xml:"exclusiveGateway"`


}

type TypeBPMN2 struct {
	XMLName   xml.Name `xml:"definitions"`
	Process   TypeProcess `xml:"process"`
}


func importPoint(titleText string, typeText string, elementId string, processId int64, loop bool) error{

		o := orm.NewOrm()
		o.Using("default")

		cnt := 0

		iLoop := 0
		if loop {
			iLoop = 1
		}
		err := o.Raw("select count(1) cnt from bp_points where code=?",elementId).QueryRow(&cnt)
		if err!=nil{
			panic(err)
		}
		if cnt == 0{
			_,err := o.Raw(
				`insert into bp_points (is_active,is_loop,code,title,type_id,process_id)
				values (1,?,?,?,(select id from bp_point_types where code=?),?)`,
				iLoop,elementId,titleText,typeText,processId).Exec()
			if err!=nil{
				panic(err)
			}
		}else{
			_,err = o.Raw(
				`update bp_points  set is_active=1, is_loop=?,title=?, type_id=(select id from bp_point_types where code=?),process_id=?
				where code=?`,
				iLoop,titleText,typeText,processId,elementId).Exec()
		}
	return nil
}


func ImportBPMN2(xmlStr string,processId int64 )error{

	v := &TypeBPMN2{}
	err := xml.Unmarshal([]byte(xmlStr), &v)
	if err!=nil{
		return err
	}


	o := orm.NewOrm()
	o.Using("default")

	_,err = o.Raw("update bp_points set is_active=0 where process_id=?",processId).Exec()
	if err!=nil{
		return err
	}

	_,err = o.Raw("update bp_sequence_flows set is_active=0 where process_id=?",processId).Exec()
	if err!=nil{
		return err
	}

	cnt := 0
	for _, element := range v.Process.StartEvent {
		importPoint(element.Name, "startevent", element.Id ,processId,false)
	}
	for _, element := range v.Process.EndEvent {
		importPoint(element.Name, "endevent", element.Id ,processId,false)
	}
	for _, element := range v.Process.UserTask {
		importPoint(element.Name, "usertask", element.Id ,processId,false)
	}
	for _, element := range v.Process.ServiceTask {
		importPoint(element.Name, "servicetask", element.Id ,processId,false)
		if len(element.StandardLoopCharacteristics)>0{
			log.Println("LOOP"+element.Name)
			importPoint(element.Name, "servicetask", element.Id ,processId,true)
		}else{
			importPoint(element.Name, "servicetask", element.Id ,processId,false)
		}
	}
	for _, element := range v.Process.ScriptTask {
		importPoint(element.Name, "scripttask", element.Id ,processId,false)
	}
	for _, element := range v.Process.ManualTask {
		importPoint(element.Name, "manualtask", element.Id ,processId,false)
	}
	for _, element := range v.Process.ExclusiveGateway {
		importPoint(element.Name, "exclusivegateway", element.Id ,processId,false)
	}
		for _, element := range v.Process.SequenceFlow {
			log.Println("SourceRef="+element.SourceRef)
		o.Raw("select count(1) cnt from bp_sequence_flows where code=?",element.Id).QueryRow(&cnt)
		if cnt==0{
			_, err := o.Raw("insert into bp_sequence_flows (is_active,code,title,process_id) values (1,?,?,?)",element.Id,element.Name,processId).Exec()
			if err!=nil{
				panic(err)
			}
		}else{
			_,err = o.Raw("update bp_sequence_flows set is_active=1, title=?,process_id=? where code=?",element.Name,processId,element.Id).Exec()
			if err!=nil{
				panic(err)
			}
		}
		cnt  = 0
		if element.SourceRef!="" {
			log.Println("element.SourceRef="+element.SourceRef)
			o.Raw("select count(1) cnt from bp_point_sfs where is_incoming=0 and sf_id=(select id from bp_sequence_flows where code=?) ", element.Id).QueryRow(&cnt)
			if cnt == 0 {
				o.Raw(`insert into bp_point_sfs
			(is_incoming,sf_id,point_id)
			 values
			 (0,(select id from bp_sequence_flows where code=?),(select id from bp_points where code=?) ) `,
					element.Id, element.SourceRef).Exec()
			}else{
				o.Raw(`update bp_point_sfs
			set point_id=(select id from bp_points where code=?)
			 where sf_id=(select id from bp_sequence_flows where code=?) and is_incoming=0`,
					element.SourceRef,element.Id).Exec()
			}
		}
		if element.TargetRef!="" {
			log.Println("element.TargetRef="+element.TargetRef)
			o.Raw("select count(1) cnt from bp_point_sfs where is_incoming=1 and sf_id=(select id from bp_sequence_flows where code=?) ", element.Id).QueryRow(&cnt)
			if cnt == 0 {
				o.Raw(`insert into bp_point_sfs
			(is_incoming,sf_id,point_id)
			 values
			 (1,(select id from bp_sequence_flows where code=?),(select id from bp_points where code=?) ) `,
					element.Id, element.TargetRef).Exec()
			}else{
				o.Raw(`update bp_point_sfs
			set point_id=(select id from bp_points where code=?)
			 where sf_id=(select id from bp_sequence_flows where code=?) and is_incoming=1`,
					element.TargetRef,element.Id).Exec()
			}
		}
	}
	return nil


}