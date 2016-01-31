package restapi
import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"log"
	"main/auth"
)

type TErrorResponse struct {
	ErrorText string `json:"error_text"`
	Error string `json:"error"`
}
func RestCheckDBPanic( err error,res http.ResponseWriter,o orm.Ormer ) bool{
	if err!=nil {
		//fmt.Println("error goi")
		errRes := TErrorResponse{Error:"1",ErrorText:err.Error()}
		jsonData, _ := json.Marshal(errRes)
		fmt.Fprint(res, string(jsonData))
		o.Rollback()
		log.Println(err)
		//panic(err)
		return true
	}
	return false
}

func RestCheckAuth(res http.ResponseWriter, req *http.Request ) bool{
	if auth.UserId(req) ==0 {
		//fmt.Println("error goi")
		errRes := TErrorResponse{Error:"2",ErrorText:"NEED AUTH"}
		jsonData, _ := json.Marshal(errRes)
		fmt.Fprint(res, string(jsonData))
		log.Println("AUTH NEED")
		//panic(err)
		return true
	}
	return false
}


func RestCheckPanic( err error,res http.ResponseWriter ) bool{
	if err!=nil {
		//fmt.Println("error goi")
		errRes := TErrorResponse{Error:"1",ErrorText:err.Error()}
		jsonData, _ := json.Marshal(errRes)
		fmt.Fprint(res, string(jsonData))
		log.Println(err)
		//panic(err)
		return true
	}
	return false
}

func CheckPanic( err error) bool{
	if err!=nil {
		//fmt.Println("error goi")
		log.Println(err)
		panic(err)
		return true
	}
	return false
}
