package file
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"io/ioutil"
	"os"
	"strconv"
	"github.com/yeldars/crm/restapi"
)

func UserPic(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if restapi.RestCheckAuth(w,r){
		return
	}

	r.ParseForm();
	var b  [] byte
	userId,err := (strconv.Atoi(r.Form.Get("id"))) //AntiHack
	if restapi.RestCheckPanic(err,w) {
		return
	}
	filename :="uploads/users/"+ strconv.Itoa(userId) +".jpg";
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		filename = "uploads/users/default.png"
	}
	b,err =  ioutil.ReadFile(filename)
	if restapi.RestCheckPanic(err,w) {
		return
	}
	w.Write(b)


}
