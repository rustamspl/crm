package restapi
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/yeldars/crm/utils/email"
	"net/smtp"
	"github.com/astaxie/beego/orm"
)



func EmailTest(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	o := orm.NewOrm()
	o.Using("default")


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
	err := o.Raw("select t.subject,t.template,ch.smtp_from,ch.smtp_host,ch.smtp_port,ch.smtp_user,ch.smtp_password from di_tasks t,di_chs ch  where t.id=1 and ch.id=t.ch_id").QueryRow(&params)
		if err != nil {
			panic(err)
		}

	m := email.NewHTMLMessage(params.Subject, params.Template)

	m.From = params.SmtpFrom
	m.To = []string{"yeldar@bk.ru"}
	err = email.Send(params.SmtpHost +  ":" +params.SmtpPort,
		smtp.PlainAuth("", params.SmtpUser,
			params.SmtpPassword,
			params.SmtpHost), m)

	if err != nil {
		panic(err)
	}
}
