package file
import (
	"net/http"
	"fmt"
	"os"
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

func ImportProfileImage(w http.ResponseWriter, r *http.Request, fileName string) (error) {


	type resultT struct {
		Result string `json:"result"`
	}

	var result resultT
	o := orm.NewOrm()
	o.Using("default")
	os.Remove("uploads/users/"+r.Form.Get("user_id")+".jpg")
	os.Link(fileName,"uploads/users/"+r.Form.Get("user_id")+".jpg")
	result.Result="ok"
	jsonData, _ := json.Marshal(result)
	fmt.Fprint(w,string(jsonData))

	return nil

}
