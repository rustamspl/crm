package main
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/yeldars/crm/macros"
	"log"
	"os"
"github.com/astaxie/beego/orm"
)

func Init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	log.Println(err)
	log.Println("openshift_port="+os.Getenv("OPENSHIFT_GO_PORT"))

	err = orm.RegisterDataBase("default", "mysql", os.Getenv("OPENSHIFT_MYSQL_DB_USERNAME")+":"+os.Getenv("OPENSHIFT_MYSQL_DB_PASSWORD")+"@tcp("+os.Getenv("OPENSHIFT_MYSQL_DB_HOST")+":"+os.Getenv("OPENSHIFT_MYSQL_DB_PORT")+")/golang?charset=utf8")
	if err!=nil{
		panic(err)
	}else{
		log.Println("ok")
	}

}

func main(){
	Init()
	i,err := macros.CreateAccountBy2Phone(12312,"Просто Гой","77772825520","77772825521")
	log.Println(i)
	log.Println(err)
}
