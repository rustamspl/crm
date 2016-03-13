package file
import (
	"net/http"
	"fmt"
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

func ImportProfileImage(w http.ResponseWriter, r *http.Request, file_id int64) (error) {


	type resultT struct {
		Result string `json:"result"`
	}

	var result resultT
	o := orm.NewOrm()
	o.Using("default")
	o.Raw("update users set user_pic_file_id=? where id=?",file_id,r.Form.Get("user_id")).Exec()
	result.Result="ok"
	jsonData, _ := json.Marshal(result)
	fmt.Fprint(w,string(jsonData))

	return nil

}
