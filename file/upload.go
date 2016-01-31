package file
import (
	"net/http"
	"fmt"
	"io"
	"os"
	"github.com/julienschmidt/httprouter"
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
		defer file.Close()
		//fmt.Fprintf(w, "%v", handler.Header)
		fileName := "./uploads/"+handler.Filename;
		f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			//fmt.Println(err)
			fmt.Fprint(w,"{\"result\": \"err\"}")
			return
		}
		defer f.Close()
		io.Copy(f, file)

	r.ParseForm()
	if r.Form.Get("action")=="import_deals01" {
		ImportDeal01(w, r, fileName)
	}	else if r.Form.Get("action")=="import_deals01_xls" {
		ImportDeal01XLS(w, r, fileName)
	}

	if r.Form.Get("action")=="profile_image" {
		ImportProfileImage(w, r, fileName)
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
