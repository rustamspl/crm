package utils
import (
	"github.com/astaxie/beego/orm"
	"os"
	"log"

)

var db = os.Getenv("OPENSHIFT_APP_NAME")

func CheckEntity(entityCode string) bool{

	o := orm.NewOrm()
	o.Using("default")
	ok := 0
	err := o.Raw("select 1 from entities where code=?",entityCode).QueryRow(&ok)
	return err == nil && ok == 1


/*	func main() {
		var validID = regexp.MustCompile(`^[a-z|0-9]+$`)
		fmt.Println(validID.MatchString("1sal/em"))
	}*/
}
func doCreateUQIndexes(entityCode string)(error){

	sql :=
	`select concat ( 'CREATE UNIQUE INDEX ', concat(e.code,'_',ea.code, '_uindex'),' ON ', e.code , ' (', ea.code,')' ) res from entity_attrs ea,entities e where e.id=ea.entity_id and ea.uq='1'
	and e.code=?
and not exists(
SELECT *
FROM information_schema.TABLE_CONSTRAINTS
WHERE constraint_type = 'UNIQUE' and table_schema=?
and constraint_name=concat(e.code,'_',ea.code, '_uindex') COLLATE utf8_unicode_ci
)`

	o := orm.NewOrm()
	o.Using("default")
	log.Println(sql)
	type addFieldsRows struct {
		Res string `json:"res"`
	}
	var ws = [] addFieldsRows{}
	//log.Println(sql)
	_,err:= o.Raw(sql, entityCode,db).QueryRows(&ws)

	if err != nil {
		log.Println("vata")
		return err
	}
	for _,element := range ws {
		if element.Res==""{
			log.Println("CONTINUE")
			continue
		}
		sql := element.Res
		log.Println("#################"+sql)
		_,err := o.Raw(sql).Exec();
		if err!=nil{
			return err
		}
	}
	return err

}

func doAlterAddFields(entityCode string)(error){


	sql := `
select
      CASE
      WHEN (dt.code = "reference")
        THEN
          concat(ea.code, ' ', dt.db_data_type, ' ', dt.addon)
      WHEN (ea.len > 0)
        THEN
          CONCAT(ea.code, ' ', dt.db_data_type, '(', ea.len, ')')
      WHEN (ea.len = 0)
        THEN
          concat(ea.code, ' ', dt.db_data_type, ' ', dt.addon) END res
from entity_attrs ea,entities e,data_types dt where e.id=ea.entity_id
                                                    and dt.id=ea.data_type_id
                                                    and e.code=?
and not exists
(select 1 from information_schema.columns i where i.table_schema=?
  and i.table_name=e.code COLLATE utf8_unicode_ci and i.column_name=ea.code  COLLATE utf8_unicode_ci
)
	`

	o := orm.NewOrm()
	o.Using("default")
	type addFieldsRows struct {
		Res string `json:"res"`
	}
	var ws = [] addFieldsRows{}
	//log.Println(sql)
	_,err:= o.Raw(sql, entityCode,db).QueryRows(&ws)

	if err != nil {
		log.Println("vata")
		return err
	}
	for _,element := range ws {
		if element.Res==""{
			log.Println("CONTINUE")
			continue
		}
		sql := "alter table "+entityCode+" add "+element.Res
		log.Println("@@@@@@@@@@"+sql)
		_,err := o.Raw(sql).Exec();
		if err!=nil{
			return err
		}
	}
	return err

}

func newFields(entityCode string) (string,error){

	sql := `
select
  GROUP_CONCAT(
      CASE
      WHEN (dt.code = "reference")
        THEN
          concat(ea.code, ' ', dt.db_data_type, ' ', dt.addon)
      WHEN (ea.len > 0)
        THEN
          CONCAT(ea.code, ' ', dt.db_data_type, '(', ea.len, ')')
      WHEN (ea.len = 0)
        THEN
          concat(ea.code, ' ', dt.db_data_type, ' ', dt.addon) END
  )
from entity_attrs ea,entities e,data_types dt where e.id=ea.entity_id
and dt.id=ea.data_type_id
and e.code=?
`
	o := orm.NewOrm()
	o.Using("default")

	res := ""
	err:= o.Raw(sql,entityCode).QueryRow(&res)
	return res,err

}

func createTable(entityCode string) error{

	o := orm.NewOrm()
	o.Using("default")
	newFld, err := newFields(entityCode)
	if err!=nil{
		return err
	}
	sql :="create table "+entityCode+" ("+newFld+")"
	log.Println("=============="+sql)
	_,err = o.Raw(sql).Exec()
	return err

}


func tableExists(entityCode string) bool{

	o := orm.NewOrm()
	o.Using("default")
	i := 0
	err := o.Raw("SELECT 1 FROM information_schema.tables	WHERE table_schema = ? AND table_name = ? LIMIT 1",db,entityCode).QueryRow(&i)
	return err == nil
}

func GenerateDDL(entityCode string) error{
	o := orm.NewOrm()
	o.Using("default")

	if !tableExists(entityCode) {
		err := createTable(entityCode)
		if err!=nil{
			return err
		}
	}else { //Existing Table
		err := doAlterAddFields(entityCode)
		if err != nil {
			return err
		}
	}

	err :=  doCreateUQIndexes(entityCode)
	if err!=nil{
		return err
	}
	return nil
}
