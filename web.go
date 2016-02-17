package main

import 	_ "github.com/go-sql-driver/mysql"
import 	"os"
import 	"fmt"
import 	"log"
import 	"net/http"



import (
	"github.com/astaxie/beego/orm"
	"github.com/yeldars/crm/models"
	"github.com/yeldars/crm/routes"
)

func init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	log.Println(err)
	log.Println("openshift_port="+os.Getenv("OPENSHIFT_GO_PORT"))

	err = orm.RegisterDataBase("default", "mysql", os.Getenv("OPENSHIFT_MYSQL_DB_USERNAME")+":"+os.Getenv("OPENSHIFT_MYSQL_DB_PASSWORD")+"@tcp("+os.Getenv("OPENSHIFT_MYSQL_DB_HOST")+":"+os.Getenv("OPENSHIFT_MYSQL_DB_PORT")+")/"+os.Getenv("OPENSHIFT_APP_NAME")+"?charset=utf8")
	if err!=nil{
		panic(err)
	}else{
		log.Println("ok")
	}

}

/*type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}
*/


func main() {

	routes.HandleInit()

	//beego.Handler("/",)
	//beego.Handler("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//beego.Router("/",&MainController{})
	//beego.Run()

	//bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
	//fmt.Printf("listening on %s...", bind)
	//err := http.ListenAndServe(bind, nil)
	//checkErr(err)

}

func checkErr(err error){
	if err != nil {
		panic (err)
	}
}

func restApi(res http.ResponseWriter, req *http.Request) {

	o := orm.NewOrm()
	o.Using("default") // Using default, you can use other database

	//user := new("models.Users")
	user := new(models.Users)
	user.Name = "slene55512"
	user.DeptId = 1
	user.CompanyId = 1
	user.Email = "test555@bk.aa31"
	o.Insert(user)


	fmt.Fprintf(res, "! Сәлем, !!! Қалың қалай? from %s"+user.Email)
}
