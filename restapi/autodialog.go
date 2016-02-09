package restapi
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"github.com/astaxie/beego/orm"
)

func AutoDialogRestApiGet(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	o := orm.NewOrm()
	err := o.Using("default")
	if err!=nil{
		panic(err)
	}
	play := ""
	id := 0
	err = o.Raw("select id,play from col_autodialogs limit 1").QueryRow(&id,&play)
	if err!=nil{
		panic(err)
	}

	//fmt.Fprint(res,"voicemail/vm-goodbye.wav")
	o.Raw("delete from col_autodialogs where id=?",id).Exec()
	fmt.Fprint(res,play)

}
