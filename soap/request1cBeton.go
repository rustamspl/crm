package soap

//http://play.golang.org/p/957GWzfdvN
import (
"encoding/xml"

)

type Request1CBetonEnvelope struct {
	XMLName xml.Name
	Body    body
}

type body struct {
	XMLName     xml.Name
	GetResponse createNewRequestResponse `xml:"CreateNewRequestResponse"`
}

type createNewRequestResponse struct {
	Return   string   `xml:"return"`
}

func ParseRequest1CBetonEnvelope(s string) (error,*Request1CBetonEnvelope){
	Soap  := []byte(s)
	res := &Request1CBetonEnvelope{}
	err := xml.Unmarshal(Soap, res)
	return err,res
}

