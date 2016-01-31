package file
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"main/restapi"
)

func UserPic(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if restapi.RestCheckAuth(w,r){
		return
	}

	r.ParseForm();
	var b  [] byte
	userId,err := (strconv.Atoi(r.Form.Get("id"))) //AntiHack
	filename :="uploads/users/"+ strconv.Itoa(userId) +".jpg";
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		filename = "uploads/users/default.png"
	}
	b,err =  ioutil.ReadFile(filename)
	w.Write(b)
	log.Println(err)

}
