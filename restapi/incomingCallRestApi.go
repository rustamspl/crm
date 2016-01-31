package restapi
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"log"
	"github.com/astaxie/beego/orm"
)


func IncomingCallRestApiGet(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	req.ParseForm();
	log.Println("INCOMING ccid = "+req.Form.Get("ccid"))
	log.Println("INCOMING channel = "+req.Form.Get("channel"))
	log.Println("INCOMING answerstate = "+req.Form.Get("answerstate"))


	o := orm.NewOrm()
	o.Using("default")

	_,err := o.Raw("insert into col_events (caller_id,channel,answerstate) values (?,?,?)",req.Form.Get("ccid"),req.Form.Get("channel"),req.Form.Get("answerstate")).Exec()
	if err == nil{
	fmt.Fprint(res,"OK")
	}else{
		fmt.Fprint(res,err.Error())
		log.Print(err)
	}

}
