package restapi
import (
"github.com/julienschmidt/httprouter"
"github.com/astaxie/beego/orm"
"net/http"
"github.com/yeldars/crm/auth"
"golang.org/x/crypto/bcrypt"
"encoding/json"
	"fmt"
	"github.com/yeldars/crm/utils"
)


func Login(res http.ResponseWriter, req *http.Request, _ httprouter.Params){

	type LoginRequest struct {
		Login string "json:`login`"
		Password string "json:`password`"
		System string "json:`system`"
	}
	type LoginResponse struct {
		Result string "json:`result`"
		RedirectURL string "json:`redirectURL`"
	}
	const loginIncorrect  = "incorrect"
	const loginOk  = "ok"
	const loginLocked  = "locked"
	const loginUnknownError  = "locked"


	decoder := json.NewDecoder(req.Body)
	var request LoginRequest
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}
	var result LoginResponse
	result.Result = loginUnknownError
	req.ParseForm();
	auth.DoLoginLog(auth.UserId(req),1)
	o := orm.NewOrm()
	o.Using("default")
	session, _ := auth.GetStore().Get(req, auth.SessionName)
	//session.Options.MaxAge =	3600
	user_id :=int64(0)
	oldPassword := ""
	loginWithoutPassword := 0
	err = o.Raw("select id,`password`,`login_without_password` from users where email=?",request.Login).QueryRow(&user_id,&oldPassword,&loginWithoutPassword)

	if loginWithoutPassword==0 {
		oldPasswordByte1 := []byte(oldPassword)
		oldPasswordByte2 := []byte(request.Password)
		err = bcrypt.CompareHashAndPassword(oldPasswordByte1, oldPasswordByte2)
	}


	if err!=nil{
		result.Result = loginIncorrect
	}	else {
		session.Values["user_id"] = user_id
		session.Values["system"] = request.System
		session.Save(req, res)
		result.Result = loginOk
		result.RedirectURL = utils.GetDomainParamValue(req.Host,"homepage")
	}


	jsonData, _ := json.Marshal(result)
	fmt.Fprint(res, string(jsonData))
}

