package main
import (
	"net/http"
	"log"
	"bytes"
	"io/ioutil"
)

func main(){
	buf := bytes.NewBufferString(
`<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:obm="obmen">
   <soap:Header/>
   <soap:Body>
      <obm:CreateNewRequest>
         <obm:Client>Привет</obm:Client>
         <obm:Quantity>100</obm:Quantity>
         <obm:Address></obm:Address>
      </obm:CreateNewRequest>
   </soap:Body>
</soap:Envelope>`	)
	resp, err := http.Post("http://89.218.77.139:8008/testmadi/ws/ws1.1cws", "application/soap+xml;charset=UTF-8;action=\"obmen#obmen:CreateNewRequest", buf)
	if err!=nil{
		panic(err)
	}
	response, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(response))
}
