package macros
import (
	"github.com/astaxie/beego/orm"

	"log"
)


func CreateAccountBy2Phone(str []string) (int64,error){

	dealId := str[0]
	phone1 :=str[1]
	phone2  := str[2]

	o := orm.NewOrm()
	o.Using("default")
	r,err := o.Raw("insert into accounts (name) values (?)","").Exec()
	if err!=nil{
		return 0,err
	}
	lid,err := r.LastInsertId()
	contTypeId := 0;
	err = o.Raw("select id from cont_types where code=? ","mobile").QueryRow(&contTypeId)
	if err!=nil{
		log.Println("vata 1")
		return 0,err
	}
	_,err = o.Raw("insert into accountconts (account_id,cont,cont_type_id) values (?,?,?)",lid,phone1,contTypeId).Exec()
	if err!=nil{
		log.Println("vata 2")
		return 0,err
	}
	_,err = o.Raw("insert into accountconts (account_id,cont,cont_type_id) values (?,?,?)",lid,phone2,contTypeId).Exec()
	if err!=nil{
		log.Println("vata 3")
		return 0,err
	}
	_,err = o.Raw("update deals set account_id=? where id=?",lid,dealId).Exec()
	if err!=nil{
		log.Println("vata 4")
		return 0,err
	}
	return lid,nil

}