package restapi
import (
"net/http"
"github.com/julienschmidt/httprouter"
	"time"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"encoding/json"
	"log"
	"github.com/yeldars/crm/auth"
)


func WaitCallRestApiGet(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	o := orm.NewOrm()
	o.Using("default")
	like := ""
	err := o.Raw("select up.value from user_params up,params p WHERE p.code='sip_channel_like' and p.id=up.param_id and up.user_id=?",auth.UserId(req)).QueryRow(&like)

	if err!=nil{
		return
	}

	//fmt.Fprint(res,"OK")

	type Resp struct{
		CallerId string `json:"caller_id"`
		Ok       bool `json:"ok"`
		AnswerState string `json:"answer_state"`
		AccountId string `json:"account_id"`
		AccountName string `json:"account_name"`
	}

	var r Resp


	_,_ = o.Raw("delete from col_events WHERE extime<DATE_SUB(now(), INTERVAL 2 MINUTE)").Exec()
	//ip := strings.Split(req.RemoteAddr,":")[0]

	//ip = "192.168.1.2"
	//ip = "82.200.155.246"
	//ip = "192.168.1.46"



	//for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 2)
		r.CallerId = "EMPTY"
		r.Ok =false
		r.AnswerState = "EMPTY"
		err = o.Raw("select answerstate,caller_id from col_events WHERE channel like ? limit 1",like).QueryRow(&r.AnswerState, &r.CallerId)




		if err==nil {

			if len(r.CallerId)==11 && strings.HasPrefix(r.CallerId,"8") {
				r.CallerId = "7" + r.CallerId[1:len(r.CallerId)]
				log.Println("changed CAllderId to "+r.CallerId)
			}

			log.Println("try to find Account by CAllderId "+r.CallerId)
			o.Raw("select a.name, ac.account_id from accountconts ac,accounts a where ac.cont = ? and ac.account_id=a.id limit 1",r.CallerId).QueryRow(&r.AccountName, &r.AccountId)

			o.Raw("delete from col_events WHERE channel like ?",like).Exec()


			r.Ok = true
			//break
		}

	//}


	jsonData, _ := json.Marshal(r)
	fmt.Fprint(res,string(jsonData))


}
