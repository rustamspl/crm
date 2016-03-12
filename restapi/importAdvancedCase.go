package restapi
import (
	"github.com/astaxie/beego/orm"
	"database/sql"
	"errors"
	"log"
)


func AdvancedImportCaseUpdate(entityCode string, o orm.Ormer , element orm.Params) (sql.Result, error) {
	if entityCode == "accounts" {
		sql := "update " + entityCode + " set title=?,bin=? where code=?";
		return o.Raw(sql, element["title"], element["bin"], element["code"]).Exec()
	}else if 	entityCode == "bi_addresses" || //Proceed standart references
				entityCode == "bi_nomens" ||
				entityCode == "bi_mobilities" ||
				entityCode == "bi_constructions" ||
				entityCode == "bi_ind_sites"  {
		sql := "update " + entityCode + " set title=? where code=?";
		return o.Raw(sql, element["title"],element["code"]).Exec()
	}else if 	entityCode == "bi_drivers"  {
		sql := "update " + entityCode + " set title=?,account_id=(select id from accounts where code=?),contact_id=(select id from contacts where code=?) where code=?";
		return o.Raw(sql, element["title"], element["account_code"], element["contact_code"],element["code"]).Exec()
	}else if 	entityCode == "bi_deals"  {
		sql := "update " + entityCode + " set title=?,account_id=(select id from accounts where code=?) where code=?";
		return o.Raw(sql, element["title"], element["account_code"], element["code"]).Exec()
	}else if 	entityCode == "contacts"  {
		sql := "update " + entityCode + " set title=?,account_id=(select id from accounts where code=?),lastname=?,firstname=?,middlename=? where code=?";
		return o.Raw(sql, element["title"], element["account_code"],element["lastname"],element["firstname"],element["middlename"], element["code"]).Exec()
	}else if 	entityCode == "bi_vehicles"  {
		sql := "update " + entityCode + " set title=?,vechicle_type_id=(select id from bi_vehicle_vids where code=?) where code=?";
		return o.Raw(sql, element["title"], element["bi_vid_ts_code"], element["code"]).Exec()
	}else if 	entityCode == "bi_vehicle_vids"  {
		sql := "update " + entityCode + " set title=?,bi_tip_ts_code=?,volume=? where code=?";
		return o.Raw(sql, element["title"], element["bi_tip_ts_code"],element["volume"], element["code"]).Exec()
	} else{
		return nil,errors.New("entity "+entityCode+" not importable")
	}
	return nil,errors.New("entity "+entityCode+" not importable")
}

func AdvancedImportCaseInsert(entityCode string, o orm.Ormer , element orm.Params) (sql.Result, error) {
	if entityCode == "accounts" {
		sql := "insert into " + entityCode + " (code,title,bin) values (?,?,?)";
		return o.Raw(sql, element["code"], element["title"], element["bin"]).Exec()
	}else if 	entityCode == "bi_addresses" || //Proceed standart references
				entityCode == "bi_nomens" ||
				entityCode == "bi_mobilities" ||
				entityCode == "bi_constructions" ||
				entityCode == "bi_ind_sites"  {
		sql := "insert into " + entityCode + " (code,title) values (?,?)";
		return o.Raw(sql, element["code"], element["title"]).Exec()
	}else if 	entityCode == "bi_drivers"  {
		sql := "insert into " + entityCode + " (code,title,account_id,contact_id) values "+
		"(?,?,(select id from accounts where code=?),(select id from contacts where code=?))";
		log.Println(sql)
		return o.Raw(sql, element["code"], element["title"], element["account_code"], element["contact_code"]).Exec()
	}else if 	entityCode == "bi_deals"  {
		sql := "insert into " + entityCode + " (code,title,account_id) values "+
		"(?,?,(select id from accounts where code=?))";
		return o.Raw(sql, element["code"], element["title"], element["account_code"]).Exec()
	}else if 	entityCode == "contacts"  {
		sql := "insert into " + entityCode + " (code,title,account_id,lastname,firstname,middlename) values "+
		"(?,?,(select id from accounts where code=?),?,?,?)";
		return o.Raw(sql, element["code"], element["title"], element["account_code"], element["lastname"], element["firstname"], element["middlename"]).Exec()
	}else if 	entityCode == "bi_vehicles"  {
		sql := "insert into " + entityCode + " (code,title,vechicle_type_id) values "+
		"(?,?,(select id from bi_vehicle_vids where code=?))";
		return o.Raw(sql, element["code"], element["title"], element["bi_vid_ts_code"]).Exec()
	}else if 	entityCode == "bi_vehicle_vids"  {
		sql := "insert into " + entityCode + " (code,title,bi_tip_ts_code,volume) values "+
		"(?,?,?,?)";
		return o.Raw(sql, element["code"], element["title"], element["bi_tip_ts_code"],element["volume"]).Exec()
	}	else{
		return nil,errors.New("entity "+entityCode+" not importable")
	}
	return nil,errors.New("entity "+entityCode+" not importable")
}
