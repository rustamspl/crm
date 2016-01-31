package restapi
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/astaxie/beego/orm"
)

func RemoveAll(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	req.ParseForm();
	if RestCheckAuth(res,req){
		return
	}
	o := orm.NewOrm()
	o.Using("default")
	_,err := o.Raw("delete from "+req.Form.Get("code")).Exec()
	if err!=nil{
		RestCheckDBPanic(err,res,o)
	}else{
		OkResponse(res,"ALL REMOVED")
	}

}

