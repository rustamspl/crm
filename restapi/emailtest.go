package restapi
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/yeldars/crm/utils/email"
	"net/smtp"
	"github.com/yeldars/crm/utils"
	"github.com/yeldars/crm/auth"
)



func EmailTest(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	m := email.NewHTMLMessage("Текст по-русски. Параметры?", "<h4>Содержимое по-русски</h4>")
	m.From = utils.GetUserParamValue(auth.UserId(req), "smtp_from")
	m.To = []string{"yeldar@bk.ru"}


//	err := m.Attach("email_test.go")
//	if err != nil {
//		panic(err)
//	}

	//err := email.Send("smtp.mail.ru:587", smtp.PlainAuth("", "kz.bapsdadps", "dasdasdasf2", "smtp.mail.ru"), m)
	err := email.Send(utils.GetUserParamValue(auth.UserId(req), "smtp_url"),
		smtp.PlainAuth("", utils.GetUserParamValue(auth.UserId(req), "smtp_user"),
			utils.GetUserParamValue(auth.UserId(req), "smtp_password"),
			utils.GetUserParamValue(auth.UserId(req), "smtp_host")), m)

	if err != nil {
		panic(err)
	}
}
