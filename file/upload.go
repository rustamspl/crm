package file
import (
	"net/http"
	"fmt"
	"io"
	"os"
	"github.com/julienschmidt/httprouter"
	"github.com/astaxie/beego/orm"
	"github.com/yeldars/crm/utils"
	"github.com/yeldars/crm/restapi"
	"runtime"
	"path"
	"strings"
	"log"
)


func Upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		r.ParseMultipartForm(32 << 20)
	    if r.ContentLength > 20000000{
			fmt.Fprint(w,"{result:\"TOO LARGE FILE\" }");
			return;
		}

		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}

	o := orm.NewOrm()
	o.Using("default")

	fileName := os.TempDir()+handler.Filename
	if r.Form.Get("dir")!="" {
		dir :=""
		path := "unix_path"
		if runtime.GOOS == "windows" {
			path = "win_path"
		}
		err:= o.Raw("select "+path+" from dirs where code=?", r.Form.Get("dir")).QueryRow(&dir)
		if (restapi.RestCheckPanic(err,w)){
			return
		}
		fileName = dir+handler.Filename
	}


	defer file.Close()

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	log.Println(fileName)
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w,"{\"result\": \""+err.Error()+"\"}")
		return
	}
	defer f.Close()
	io.Copy(f, file)


	var file_id = int64(0)
	if (r.Form.Get("dir")!="") {

		baseName := path.Base(strings.Replace(fileName,"\\","/",-1))
		rs, err := o.Raw("insert into files (dir_id,code,title,filename) values ((select id from dirs where code=?),?,?,?)", r.Form.Get("dir"),utils.Uuid(), "TEST123", baseName).Exec()

		if (restapi.RestCheckPanic(err,w)){
			return
		}

		file_id,err = rs.LastInsertId()
	}


	r.ParseForm()
	if r.Form.Get("action")=="import_deals01" {
		ImportDeal01(w, r, fileName)
	}	else if r.Form.Get("action")=="import_deals01_xls" {
		ImportDeal01XLS(w, r, fileName)
	}else if r.Form.Get("action")=="universalimport" {
		ImportUniversal(w, r, fileName)
	}

	if r.Form.Get("action")=="profile_image" {
		ImportProfileImage(w, r, int64(file_id))
	}
		/*o := orm.NewOrm()
		o.Using("default")

		var ff  models.Files
		ff.FileName = "test222.txt";
		//ff.Data,err =  ioutil.ReadFile(fileName)
		ff.Data = "fdsafdsfoijdsfidshfidsh"
		pi,err := o.QueryTable("files").PrepareInsert();
		pi.Insert(&ff)




		*/


}
