package file
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"io/ioutil"
	"github.com/yeldars/crm/restapi"
	"github.com/astaxie/beego/orm"
	"runtime"
)

func GetFile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {


	path := "unix_path"
	if runtime.GOOS == "windows" {
		path = "win_path"
	}

	o := orm.NewOrm()
	o.Using("default")

	fileName := ""
	r.ParseForm()
	err := o.Raw("select concat(d."+path+",f.filename) fileName from files f,dirs d where d.id=f.dir_id and f.code=?",r.Form.Get("code")).QueryRow(&fileName)
	if restapi.RestCheckPanic(err,w){
		return
	}

	b,err :=  ioutil.ReadFile(fileName)
	if restapi.RestCheckPanic(err,w) {
		return
	}
	w.Write(b)
}
