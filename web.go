package main

import 	_ "github.com/go-sql-driver/mysql"
import 	"os"
import 	"log"




import (
	"github.com/astaxie/beego/orm"
	"github.com/yeldars/crm/routes"
)

func init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err!=nil{
		panic(err)
	}

	err = orm.RegisterDataBase("default", "mysql", os.Getenv("OPENSHIFT_MYSQL_DB_USERNAME")+":"+os.Getenv("OPENSHIFT_MYSQL_DB_PASSWORD")+"@tcp("+os.Getenv("OPENSHIFT_MYSQL_DB_HOST")+":"+os.Getenv("OPENSHIFT_MYSQL_DB_PORT")+")/"+os.Getenv("OPENSHIFT_APP_NAME")+"?charset=utf8")
	if err!=nil{
		panic(err)
	}else{
		log.Println("ok... openshift_port="+os.Getenv("OPENSHIFT_GO_PORT"))
	}

}

func main() {
	routes.HandleInit()
}
