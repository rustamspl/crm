package auth
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/sessions"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"fmt"
	"github.com/yeldars/crm/models"
	"github.com/yeldars/crm/utils"
	"golang.org/x/crypto/bcrypt"
)


var store = sessions.NewCookieStore([]byte("asdjkjkl39090wejiosdfklo"))

var SessionName = "BAPPSSessionId"

func GetStore() * sessions.CookieStore{
	return store
}
func UserId(req *http.Request) int64{
	session, _ := store.Get(req, SessionName)
	if (session.Values["user_id"]==nil){
		session.Values["user_id"]=int64(0)
	}else{
		//session.Options.MaxAge =	3600
	}
	return session.Values["user_id"].(int64)
}

func System(req *http.Request) string{
	session, _ := store.Get(req, SessionName)
	if (session.Values["system"]==nil){
		session.Values["system"]=""
	}else{
		//session.Options.MaxAge =	3600
	}
	return session.Values["system"].(string)
}

type SessionInfo struct {
	UserId int64 `json:"user_id"`
}

func GetSessionInfo(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  var s SessionInfo
	s.UserId = UserId(req)
	jsonData, _ := json.Marshal(s)
	//checkErr(err)
	fmt.Fprint(res,string(jsonData))
}
func Login(res http.ResponseWriter, req *http.Request, _ httprouter.Params){

	req.ParseForm();

	DoLoginLog(UserId(req),1)
	o := orm.NewOrm()
	o.Using("default")
	session, _ := store.Get(req, SessionName)
	//session.Options.MaxAge =	3600
	user_id :=int64(0)
	//err := o.Raw("select id from users u where u.email=? and `password`=?",req.PostForm.Get("email"),req.PostForm.Get("password")).QueryRow(&user_id )

	oldPassword := ""
	loginWithoutPassword := 0
	err := o.Raw("select id,`password`,`login_without_password` from users where email=?",req.PostForm.Get("email")).QueryRow(&user_id,&oldPassword,&loginWithoutPassword)
	//log.Println(oldPassword)


	if loginWithoutPassword==0 {
		oldPasswordByte1 := []byte(oldPassword)
		oldPasswordByte2 := []byte(req.PostForm.Get("password"))
		err = bcrypt.CompareHashAndPassword(oldPasswordByte1, oldPasswordByte2)
	}

	if err!=nil{
		//fmt.Fprint(res,err);
		http.Redirect(res, req, utils.GetDomainParamValue(req.Host,"loginpage")+"#invalidlogin", 301);
	}	else {
		session.Values["user_id"] = user_id
		session.Save(req, res)
		http.Redirect(res, req, utils.GetDomainParamValue(req.Host,"homepage"), 301)
	}
}


func GetLanguage2(req *http.Request) string{
	session, _ := store.Get(req, SessionName)
	if (session.Values["lang"]==nil){
		session.Values["lang"]="ru"
	}

	return session.Values["lang"].(string)
}


func GetLanguage(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	session, _ := store.Get(req, SessionName)
	//req.ParseForm();

	fmt.Fprint(res, "{\"lang\":\""+ session.Values["lang"].(string) + "\"}")
}

func SetLanguage(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	session, _ := store.Get(req, SessionName)
	req.ParseForm();
	session.Values["lang"]=req.Form.Get("lang")
	//fmt.Fprint(res,"TEST"+req.Form.Get("lang"))
	session.Save(req, res)
}

func Logout(res http.ResponseWriter, req *http.Request, _ httprouter.Params){

	DoLoginLog(UserId(req),2)
	session, _ := store.Get(req, SessionName)
	session.Values["user_id"] = int64(0)
	session.Save(req, res)

	http.Redirect(res, req, utils.GetDomainParamValue(req.Host,"loginpage"), 301)
}

func DoLoginLog(user_id int64,login_type int64){

	var userLog models.LoginLogs

	userLog.UserId = user_id
	userLog.LoginType = login_type

	o := orm.NewOrm()
	o.Using("default")

	qs := o.QueryTable("login_logs")
	i, _ := qs.PrepareInsert()
	i.Insert(&userLog)
	i.Close()

}
