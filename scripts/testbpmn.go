package main
import "github.com/yeldars/crm/bpms"


func main(){

 bpms.ImportBPMN2(

	 `<definitions xmlns="http://www.omg.org/spec/BPMN/20100524/MODEL" xmlns:bpmndi="http://www.omg.org/spec/BPMN/20100524/DI" xmlns:omgdc="http://www.omg.org/spec/DD/20100524/DC" xmlns:omgdi="http://www.omg.org/spec/DD/20100524/DI" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" targetNamespace="" xsi:schemaLocation="http://www.omg.org/spec/BPMN/20100524/MODEL http://www.omg.org/spec/BPMN/2.0/20100501/BPMN20.xsd">
  <process id="Process_0ugra0m">
    <startEvent id="StartEvent_0gn6c38">
      <outgoing>SequenceFlow_1m1uui4</outgoing>
    </startEvent>
    <endEvent id="EndEvent_1g4svdd">
      <incoming>SequenceFlow_1ghv9eh</incoming>
      <incoming>SequenceFlow_15bk45l</incoming>
    </endEvent>
    <sequenceFlow id="SequenceFlow_1m1uui4" sourceRef="StartEvent_0gn6c38" targetRef="ExclusiveGateway_0ihq7ch" />
    <exclusiveGateway id="ExclusiveGateway_0ihq7ch">
      <incoming>SequenceFlow_1m1uui4</incoming>
      <outgoing>SequenceFlow_1ghv9eh</outgoing>
      <outgoing>SequenceFlow_09dspz3</outgoing>
    </exclusiveGateway>
    <sequenceFlow id="SequenceFlow_1ghv9eh" sourceRef="ExclusiveGateway_0ihq7ch" targetRef="EndEvent_1g4svdd" />
    <sequenceFlow id="SequenceFlow_09dspz3" sourceRef="ExclusiveGateway_0ihq7ch" targetRef="UserTask_0lutumk" />
    <userTask id="UserTask_0lutumk">
      <incoming>SequenceFlow_09dspz3</incoming>
      <outgoing>SequenceFlow_15bk45l</outgoing>
    </userTask>
    <sequenceFlow id="SequenceFlow_15bk45l" sourceRef="UserTask_0lutumk" targetRef="EndEvent_1g4svdd" />
  </process></definitions>`)
}



