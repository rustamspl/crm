package main
import (
	"github.com/yeldars/crm/soap"
	"log"
)

func main(){

	log.Println("Start...")
	_,x := soap.ParseRequest1CBetonEnvelope(`<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope">
   <soap:Header/>
   <soap:Body>
      <m:CreateNewRequestResponse xmlns:m="obmen">
         <m:return xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">0000000034</m:return>
         <m:KebRequestNumber xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">000000003</m:KebRequestNumber>
      </m:CreateNewRequestResponse>
   </soap:Body>
</soap:Envelope>`)

	log.Println(x.Body.GetResponse.Return)
}
