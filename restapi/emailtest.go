package restapi
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/yeldars/crm/utils/email"
	"net/smtp"
	"github.com/astaxie/beego/orm"
	"strconv"
	"log"
)



func EmailTest(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	o := orm.NewOrm()
	o.Using("default")

	req.ParseForm()
	taskId,errCnv := strconv.Atoi(req.Form.Get("task_id"))
	if errCnv != nil {
		panic(errCnv)
	}
	type TParams struct {
		Subject string `json:"subject"`
		Template string `json:"template"`
		SmtpFrom string `json:"smtp_from"`
		SmtpHost string `json:"smtp_host"`
		SmtpPort string `json:"smtp_port"`
		SmtpUser string `json:"smtp_user"`
		SmtpPassword string `json:"smtp_password"`

	}
	var params TParams
	err := o.Raw("select t.subject,t.template,ch.smtp_from,ch.smtp_host,ch.smtp_port,ch.smtp_user,ch.smtp_password from di_templates t,di_chs ch,di_tasks ta  where ta.id=? and ta.template_id=t.id and ch.id=t.ch_id",taskId).QueryRow(&params)
	if err != nil {
		panic(err)
	}

	m := email.NewHTMLMessage(params.Subject, params.Template)


	address := ""
	recId := 0
	err = o.Raw("select id,address from di_task_recs where task_id=? and status=? limit 1",taskId,0).QueryRow(&recId,&address)
	if err !=nil{
		//panic(err)
		return
	}

	m.From = params.SmtpFrom
	//m.To = []string{"nabievnurlan7@gmail.com"}
	m.To = []string{address}

	_,err = o.Raw("update di_task_recs set status=? where id=?",3,recId).Exec() //Queue
	if err != nil {
		panic(err)
	}

	err = email.Send(params.SmtpHost +  ":" +params.SmtpPort,
		smtp.PlainAuth("", params.SmtpUser,
			params.SmtpPassword,
			params.SmtpHost), m)
	log.Println("sent to "+m.To[0])
	if err != nil {
		o.Raw("update di_task_recs set status=?,err_txt=? where id=?",2,err.Error(),recId).Exec() //Err
		panic(err)
	}

	o.Raw("update di_task_recs set status=? where id=?",1,recId).Exec() //Ok
}
