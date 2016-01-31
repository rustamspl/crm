package auth
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"fmt"
	"log"
	"golang.org/x/crypto/bcrypt"
)
type TPasswordResetResult struct {
 Error int64 `json:"error"`
}

type TPasswordResetRequest struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
	UserId string `json:"user_id"`
}


func ResetPassword(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	var result TPasswordResetResult
	var request TPasswordResetRequest
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&request)

	o := orm.NewOrm()
	o.Using("default")

	oldPassword := ""
	password_not_set := 0
	o.Raw("select `password`,`password_not_set` from users where id=?",request.UserId).QueryRow(&oldPassword,&password_not_set)

	if password_not_set==0 {
		log.Println(oldPassword)
		oldPasswordByte1 := []byte(oldPassword)
		oldPasswordByte2 := []byte(request.OldPassword)

		err = bcrypt.CompareHashAndPassword(oldPasswordByte1, oldPasswordByte2)
		//fmt.Println(err) // nil means it is a match
		if err != nil {
			result.Error = 1
			jsonData, _ := json.Marshal(result)
			fmt.Fprint(res, string(jsonData))
			log.Println("password incorrect")
			return
		}
	}

	password := []byte(request.NewPassword)

	hashedPassword, err := bcrypt.GenerateFromPassword(password,  bcrypt.DefaultCost)

	_,err =o.Raw("update users set `password_not_set`=0,`password`=? where id=?",string(hashedPassword),request.UserId).Exec();
	if err!=nil{
		panic(err)
	}
	result.Error = 0
	jsonData, err := json.Marshal(result)
	if err!=nil{
		panic(err)
	}
	//checkErr(err)
	fmt.Fprint(res,string(jsonData))

}