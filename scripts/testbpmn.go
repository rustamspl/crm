package main

import 	_ "github.com/go-sql-driver/mysql"

import (
	"github.com/yeldars/crm/bpms"
	"github.com/astaxie/beego/orm"
	"os"
	"log"
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


 bpms.ImportBPMN2(

	 `<?xml version="1.0" encoding="UTF-8"?>
<definitions xmlns="http://www.omg.org/spec/BPMN/20100524/MODEL" xmlns:bpmndi="http://www.omg.org/spec/BPMN/20100524/DI" xmlns:omgdc="http://www.omg.org/spec/DD/20100524/DC" xmlns:omgdi="http://www.omg.org/spec/DD/20100524/DI" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" targetNamespace="" xsi:schemaLocation="http://www.omg.org/spec/BPMN/20100524/MODEL http://www.omg.org/spec/BPMN/2.0/20100501/BPMN20.xsd">
  <process id="Process_0ugra0m">
    <startEvent id="StartEvent_0gn6c38">
      <outgoing>SequenceFlow_1m1uui4</outgoing>
      <outgoing>SequenceFlow_000u2lx</outgoing>
      <outgoing>SequenceFlow_1b5sxty</outgoing>
    </startEvent>
    <endEvent id="EndEvent_1g4svdd">
      <incoming>SequenceFlow_15bk45l</incoming>
      <incoming>SequenceFlow_00erru5</incoming>
      <incoming>SequenceFlow_0w9gwte</incoming>
    </endEvent>
    <sequenceFlow id="SequenceFlow_1m1uui4" sourceRef="StartEvent_0gn6c38" targetRef="ExclusiveGateway_1nd0dxe" />
    <userTask id="UserTask_0lutumk">
      <incoming>SequenceFlow_13buy3h</incoming>
      <outgoing>SequenceFlow_15bk45l</outgoing>
    </userTask>
    <sequenceFlow id="SequenceFlow_15bk45l" sourceRef="UserTask_0lutumk" targetRef="EndEvent_1g4svdd" />
    <exclusiveGateway id="ExclusiveGateway_1nd0dxe">
      <incoming>SequenceFlow_1m1uui4</incoming>
      <outgoing>SequenceFlow_13buy3h</outgoing>
    </exclusiveGateway>
    <sequenceFlow id="SequenceFlow_13buy3h" name="test" sourceRef="ExclusiveGateway_1nd0dxe" targetRef="UserTask_0lutumk" />
    <sequenceFlow id="SequenceFlow_000u2lx" sourceRef="StartEvent_0gn6c38" targetRef="UserTask_1xxtdyn" />
    <sequenceFlow id="SequenceFlow_00erru5" name="конец" sourceRef="UserTask_1xxtdyn" targetRef="EndEvent_1g4svdd" />
    <userTask id="UserTask_1xxtdyn">
      <incoming>SequenceFlow_000u2lx</incoming>
      <outgoing>SequenceFlow_00erru5</outgoing>
    </userTask>
    <userTask id="UserTask_1muigf1">
      <incoming>SequenceFlow_1b5sxty</incoming>
      <outgoing>SequenceFlow_0w9gwte</outgoing>
    </userTask>
    <sequenceFlow id="SequenceFlow_1b5sxty" sourceRef="StartEvent_0gn6c38" targetRef="UserTask_1muigf1" />
    <sequenceFlow id="SequenceFlow_0w9gwte" sourceRef="UserTask_1muigf1" targetRef="EndEvent_1g4svdd" />
  </process>
  <bpmndi:BPMNDiagram id="sid-74620812-92c4-44e5-949c-aa47393d3830">
    <bpmndi:BPMNPlane id="sid-cdcae759-2af7-4a6d-bd02-53f3352a731d" bpmnElement="Process_0ugra0m">
      <bpmndi:BPMNShape id="StartEvent_0gn6c38_di" bpmnElement="StartEvent_0gn6c38">
        <omgdc:Bounds x="244" y="176" width="36" height="36" />
        <bpmndi:BPMNLabel>
          <omgdc:Bounds x="217" y="212" width="90" height="20" />
        </bpmndi:BPMNLabel>
      </bpmndi:BPMNShape>
      <bpmndi:BPMNShape id="EndEvent_1g4svdd_di" bpmnElement="EndEvent_1g4svdd">
        <omgdc:Bounds x="541" y="305" width="36" height="36" />
        <bpmndi:BPMNLabel>
          <omgdc:Bounds x="514" y="341" width="90" height="20" />
        </bpmndi:BPMNLabel>
      </bpmndi:BPMNShape>
      <bpmndi:BPMNEdge id="SequenceFlow_1m1uui4_di" bpmnElement="SequenceFlow_1m1uui4">
        <omgdi:waypoint xsi:type="omgdc:Point" x="280" y="194" />
        <omgdi:waypoint xsi:type="omgdc:Point" x="345" y="194" />
        <bpmndi:BPMNLabel>
          <omgdc:Bounds x="343" y="184" width="90" height="20" />
        </bpmndi:BPMNLabel>
      </bpmndi:BPMNEdge>
      <bpmndi:BPMNShape id="UserTask_0lutumk_di" bpmnElement="UserTask_0lutumk">
        <omgdc:Bounds x="509.11976047904193" y="86.85163007318698" width="100" height="80" />
      </bpmndi:BPMNShape>
      <bpmndi:BPMNEdge id="SequenceFlow_15bk45l_di" bpmnElement="SequenceFlow_15bk45l">
        <omgdi:waypoint xsi:type="omgdc:Point" x="559" y="167" />
        <omgdi:waypoint xsi:type="omgdc:Point" x="559" y="305" />
        <bpmndi:BPMNLabel>
          <omgdc:Bounds x="490.5" y="209" width="90" height="20" />
        </bpmndi:BPMNLabel>
      </bpmndi:BPMNEdge>
      <bpmndi:BPMNShape id="ExclusiveGateway_1nd0dxe_di" bpmnElement="ExclusiveGateway_1nd0dxe" isMarkerVisible="true">
        <omgdc:Bounds x="345.11976047904193" y="169" width="50" height="50" />
        <bpmndi:BPMNLabel>
          <omgdc:Bounds x="325.11976047904193" y="219" width="90" height="20" />
        </bpmndi:BPMNLabel>
      </bpmndi:BPMNShape>
      <bpmndi:BPMNEdge id="SequenceFlow_13buy3h_di" bpmnElement="SequenceFlow_13buy3h">
        <omgdi:waypoint xsi:type="omgdc:Point" x="370" y="169" />
        <omgdi:waypoint xsi:type="omgdc:Point" x="370" y="127" />
        <omgdi:waypoint xsi:type="omgdc:Point" x="509" y="127" />
        <bpmndi:BPMNLabel>
          <omgdc:Bounds x="325" y="138" width="90" height="20" />
        </bpmndi:BPMNLabel>
      </bpmndi:BPMNEdge>
      <bpmndi:BPMNEdge id="SequenceFlow_000u2lx_di" bpmnElement="SequenceFlow_000u2lx">
        <omgdi:waypoint xsi:type="omgdc:Point" x="262" y="212" />
        <omgdi:waypoint xsi:type="omgdc:Point" x="262" y="241" />
        <omgdi:waypoint xsi:type="omgdc:Point" x="262" y="241" />
        <omgdi:waypoint xsi:type="omgdc:Point" x="262" y="283" />
        <bpmndi:BPMNLabel>
          <omgdc:Bounds x="233" y="231" width="90" height="20" />
        </bpmndi:BPMNLabel>
      </bpmndi:BPMNEdge>
      <bpmndi:BPMNEdge id="SequenceFlow_00erru5_di" bpmnElement="SequenceFlow_00erru5">
        <omgdi:waypoint xsi:type="omgdc:Point" x="312" y="323" />
        <omgdi:waypoint xsi:type="omgdc:Point" x="427" y="323" />
        <omgdi:waypoint xsi:type="omgdc:Point" x="427" y="323" />
        <omgdi:waypoint xsi:type="omgdc:Point" x="541" y="323" />
        <bpmndi:BPMNLabel>
          <omgdc:Bounds x="397.5" y="299" width="90" height="20" />
        </bpmndi:BPMNLabel>
      </bpmndi:BPMNEdge>
      <bpmndi:BPMNShape id="UserTask_1xxtdyn_di" bpmnElement="UserTask_1xxtdyn">
        <omgdc:Bounds x="212.30838323353294" y="283" width="100" height="80" />
      </bpmndi:BPMNShape>
      <bpmndi:BPMNShape id="UserTask_1muigf1_di" bpmnElement="UserTask_1muigf1">
        <omgdc:Bounds x="212.074" y="42.27499999999998" width="100" height="80" />
      </bpmndi:BPMNShape>
      <bpmndi:BPMNEdge id="SequenceFlow_1b5sxty_di" bpmnElement="SequenceFlow_1b5sxty">
        <omgdi:waypoint xsi:type="omgdc:Point" x="262" y="176" />
        <omgdi:waypoint xsi:type="omgdc:Point" x="262" y="148" />
        <omgdi:waypoint xsi:type="omgdc:Point" x="262" y="148" />
        <omgdi:waypoint xsi:type="omgdc:Point" x="262" y="122" />
        <bpmndi:BPMNLabel>
          <omgdc:Bounds x="206" y="138" width="90" height="20" />
        </bpmndi:BPMNLabel>
      </bpmndi:BPMNEdge>
      <bpmndi:BPMNEdge id="SequenceFlow_0w9gwte_di" bpmnElement="SequenceFlow_0w9gwte">
        <omgdi:waypoint xsi:type="omgdc:Point" x="312" y="82" />
        <omgdi:waypoint xsi:type="omgdc:Point" x="795" y="82" />
        <omgdi:waypoint xsi:type="omgdc:Point" x="795" y="323" />
        <omgdi:waypoint xsi:type="omgdc:Point" x="577" y="323" />
        <bpmndi:BPMNLabel>
          <omgdc:Bounds x="371" y="184" width="90" height="20" />
        </bpmndi:BPMNLabel>
      </bpmndi:BPMNEdge>
    </bpmndi:BPMNPlane>
    <bpmndi:BPMNLabelStyle id="sid-e0502d32-f8d1-41cf-9c4a-cbb49fecf581">
      <omgdc:Font name="Arial" size="11" isBold="false" isItalic="false" isUnderline="false" isStrikeThrough="false" />
    </bpmndi:BPMNLabelStyle>
    <bpmndi:BPMNLabelStyle id="sid-84cb49fd-2f7c-44fb-8950-83c3fa153d3b">
      <omgdc:Font name="Arial" size="12" isBold="false" isItalic="false" isUnderline="false" isStrikeThrough="false" />
    </bpmndi:BPMNLabelStyle>
  </bpmndi:BPMNDiagram>
</definitions>`)
}



