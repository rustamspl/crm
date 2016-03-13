package restapi
import (
	"github.com/astaxie/beego/orm"
	"database/sql"
	"errors"
	"log"
)


func AdvancedImportCaseUpdate(entityCode string, o orm.Ormer , element orm.Params) (sql.Result, error) {
	if entityCode == "accounts" {
		sql := "update " + entityCode + " set title=?,bin=?,kpp=?,fullname=?,address_fiz=?,address_jur=?,is_apt=?,main_contact_id=(select id from contacts where code=?),is_provider=? where code=?";
		return o.Raw(sql, element["title"], element["bin"],element["kpp"],element["fullname"],element["address_fiz"],element["address_jur"],element["is_apt"],element["main_contact_id"],element["is_provider"], element["code"]).Exec()
	}else if 	//Proceed standart references
				entityCode == "bi_mobilities" ||
				entityCode == "bi_constructions" ||
				entityCode == "bi_ind_sites"  {
		sql := "update " + entityCode + " set title=? where code=?";
		return o.Raw(sql, element["title"],element["code"]).Exec()
	}else if 	entityCode == "bi_nomens"  {
		sql := "update " + entityCode + " set title=?,article=?,model=?,frost=?,water=?,mobility=?,unit=? where code=?";
		return o.Raw(sql, element["title"], element["article"], element["model"],element["frost"],element["water"],element["mobility"],element["unit"],element["code"]).Exec()
	}else if 	entityCode == "bi_addresses"  {
		sql := "update " + entityCode + " set title=?,lat=?,lon=? where code=?";
		return o.Raw(sql, element["title"], element["lat"], element["lon"],element["code"]).Exec()
	}else if 	entityCode == "bi_drivers"  {
		sql := "update " + entityCode + " set title=?,account_id=(select id from accounts where code=?),contact_id=(select id from contacts where code=?) where code=?";
		return o.Raw(sql, element["title"], element["account_code"], element["contact_code"],element["code"]).Exec()
	}else if 	entityCode == "bi_deals"  {
		sql := "update " + entityCode + " set title=?,account_id=(select id from accounts where code=?),active=? where code=?";
		return o.Raw(sql, element["title"], element["account_code"],element["active"], element["code"]).Exec()
	}else if 	entityCode == "contacts"  {
		sql := "update " + entityCode + " set title=?,account_id=(select id from accounts where code=?),lastname=?,firstname=?,middlename=?,dscr=?,delivery_address_id=(select id from bi_addresses where code=?) where code=?";
		return o.Raw(sql, element["title"], element["account_code"],element["lastname"],element["firstname"],element["middlename"],element["dscr"],element["delivery_address_code"], element["code"]).Exec()
	}else if 	entityCode == "bi_vehicles"  {
		sql := "update " + entityCode + " set title=?,vechicle_type_id=(select id from bi_vehicle_vids where code=?) where code=?";
		return o.Raw(sql, element["title"], element["bi_vid_ts_code"], element["code"]).Exec()
	}else if 	entityCode == "bi_vehicle_vids"  {
		sql := "update " + entityCode + " set title=?,bi_tip_ts_code=?,volume=? where code=?";
		return o.Raw(sql, element["title"], element["bi_tip_ts_code"],element["volume"], element["code"]).Exec()
	}else if 	entityCode == "bi_gosnum"  {
		sql := "update " + entityCode + " set title=?,reg_at=?,vehicle_id=(select id from bi_vehicles where code=?) where code=?";
		return o.Raw(sql, element["title"], element["reg_at"],element["vehicles_code"], element["code"]).Exec()
	}else if 	entityCode == "bi_individuals"  {
		sql := "update " + entityCode + " set title=?,lastname=?,firstname=?,middlename=?,position=?,rolset=? where code=?";
		return o.Raw(sql, element["title"], element["lastname"],element["firstname"],element["middlename"],element["position"],element["rolset"], element["code"]).Exec()
	}else if 	entityCode == "bi_beton_invoices"  {
		/*                START BETON INVOICE DOCUMENT */
		/*                ****************       */
		sql := "update " + entityCode + `
		set
			doc_at=?,
		 	account_id=(select id from accounts where code=?),
		 	ind_site_id=(select id from bi_ind_sites where code=?),
		 	deal_id=(select id from bi_deals where code=?),
		 	seal=?,
		 	is_central=?,
		 	delivery_address_id=(select id from bi_addresses where code=?),
		 	driver_id=(select id from bi_drivers where code=?),
		 	apt_account_id=(select id from accounts where code=?),
		 	vehicle_id=(select id from bi_vehicles where code=?),
		 	delivery_type=?,
		 	departure_at=?,
		 	beton_req_id=(select id from bi_beton_reqs where code=?),
		 	shipped_quantity=?,
		 	dscr=?,
		 	mod_nomen_id=(select id from bi_nomens where code=?),
		 	addon_nomen_id=(select id from bi_nomens where code=?),
		 	owner_1c=?,
		 	node=?,
		 	is_cancel=?,
		 	delivery_amount=?,
		 	is_rеturned=?,
		 	return_at=?,
		 	individual_id=(select id from bi_nomens where code=?),
		 	delay_amount=?,
		 	deal_logist_id=(select id from bi_deals where code=?),
		 	recipe=?,
		 	ported=?

		 where code=?`;
		return o.Raw(sql,
			element["doc_at"],
			element["account_code"],
			element["ind_site_code"],
			element["deal_code"],
			element["seal"],
			element["is_central"],
			element["delivery_address_code"],
			element["driver_code"],
			element["apt_account_code"],
			element["vehicle_code"],
			element["delivery_type"],
			element["departure_at"],
			element["beton_req_code"],
			element["shipped_quantity"],
			element["dscr"],
			element["mod_nomen_code"],
			element["addon_nomen_code"],
			element["owner_1c"],
			element["node"],
			element["is_cancel"],
			element["delivery_amount"],
			element["is_rеturned"],
			element["return_at"],
			element["individual_code"],
			element["delay_amount"],
			element["deal_logist_code"],
			element["recipe"],
			element["ported"],
			element["code"],


		).Exec()
		/*                ****************       */
		/*                END   INVOICE DOCUMENT */
	}else{
		return nil,errors.New("entity "+entityCode+" not importable")
	}
	return nil,errors.New("entity "+entityCode+" not importable")
}

func AdvancedImportCaseInsert(entityCode string, o orm.Ormer , element orm.Params) (sql.Result, error) {
	if entityCode == "accounts" {
		sql := "insert into " + entityCode + " (code,title,bin,kpp,fullname,address_fiz,address_jur,is_apt,main_contact_id,is_provider) values (?,?,?,?,?,?,?,?,(select id from contacts where code=?),?)";
		return o.Raw(sql, element["code"], element["title"], element["bin"],  element["kpp"], element["fullname"], element["address_fiz"], element["address_jur"], element["is_apt"], element["main_contact_id"],element["is_provider"]).Exec()
	}else if 	//Proceed standart references
				entityCode == "bi_mobilities" ||
				entityCode == "bi_constructions" ||
				entityCode == "bi_ind_sites"  {
		sql := "insert into " + entityCode + " (code,title) values (?,?)";
		return o.Raw(sql, element["code"], element["title"]).Exec()
	}else if 	entityCode == "bi_addresses" {
		sql := "insert into " + entityCode + " (code,title,lat,lon) values (?,?,?,?)";
		return o.Raw(sql, element["code"], element["title"], element["lat"], element["lon"]).Exec()
	}else if 	entityCode == "bi_nomens" {
		sql := "insert into " + entityCode + " (code,title,article,model,frost,water,mobility,unit) values (?,?,?,?,?,?,?,?)";
		return o.Raw(sql, element["code"], element["title"], element["article"], element["model"], element["frost"], element["water"], element["mobility"], element["unit"]).Exec()
	}else if 	entityCode == "bi_drivers"  {
		sql := "insert into " + entityCode + " (code,title,account_id,contact_id) values "+
		"(?,?,(select id from accounts where code=?),(select id from contacts where code=?))";
		log.Println(sql)
		return o.Raw(sql, element["code"], element["title"], element["account_code"], element["contact_code"]).Exec()
	}else if 	entityCode == "bi_deals"  {
		sql := "insert into " + entityCode + " (code,title,account_id,active) values "+
		"(?,?,(select id from accounts where code=?),?)";
		return o.Raw(sql, element["code"], element["title"], element["account_code"], element["active"]).Exec()
	}else if 	entityCode == "contacts"  {
		sql := "insert into " + entityCode + " (code,title,account_id,lastname,firstname,middlename,dscr,delivery_address_id) values "+
		"(?,?,(select id from accounts where code=?),?,?,?,?,(select id from bi_addresses where code=?))";
		return o.Raw(sql, element["code"], element["title"], element["account_code"], element["lastname"], element["firstname"], element["middlename"], element["dscr"], element["delivery_address_code"]).Exec()
	}else if 	entityCode == "bi_vehicles"  {
		sql := "insert into " + entityCode + " (code,title,vechicle_type_id) values "+
		"(?,?,(select id from bi_vehicle_vids where code=?))";
		return o.Raw(sql, element["code"], element["title"], element["bi_vid_ts_code"]).Exec()
	}else if 	entityCode == "bi_vehicle_vids"  {
		sql := "insert into " + entityCode + " (code,title,bi_tip_ts_code,volume) values "+
		"(?,?,?,?)";
		return o.Raw(sql, element["code"], element["title"], element["bi_tip_ts_code"],element["volume"]).Exec()
	}else if 	entityCode == "bi_gosnum"  {
		sql := "insert into " + entityCode + " (code,title,reg_at,vehicle_id) values "+
		"(?,?,?,(select id from bi_vehicles where code=?))";
		return o.Raw(sql, element["code"], element["title"], element["reg_at"],element["vehicles_code"]).Exec()
	}else if 	entityCode == "bi_individuals"  {
		sql := "insert into " + entityCode + " (code,title,lastname,firstname,middlename,position,rolset) values "+
		"(?,?,?,?,?,?,?)";
		return o.Raw(sql, element["code"], element["title"], element["lastname"], element["firstname"], element["middlename"], element["position"], element["rolset"]).Exec()
	}else if 	entityCode == "bi_beton_invoices"  {
		/*                START BETON INVOICE DOCUMENT */
		/*                ****************       */
		sql := "insert into " + entityCode +
		`
		(code,
		doc_at,
		account_id,
		ind_site_id,
		deal_id,
		seal,
		is_central,
		delivery_address_id,
		driver_id,
		apt_account_id,
		vehicle_id,
		delivery_type,
		departure_at,
		beton_req_id,
		shipped_quantity,
		dscr,
		mod_nomen_id,
		addon_nomen_id,
		owner_1c,
		node,
		is_cancel,
		delivery_amount,
		is_rеturned,
		return_at,
		individual_id,
		delay_amount,
		deal_logist_id,
		recipe,
		ported
		)

values
		(
			?,
			?,
			(select id from accounts where code=?),
			(select id from bi_ind_sites where code=?),
			(select id from bi_deals where code=?),
			?,
			?,
			(select id from bi_addresses where code=?),
			(select id from bi_drivers where code=?),
			(select id from accounts where code=?),
			(select id from bi_vehicles where code=?),
			?,
			?,
			(select id from bi_beton_reqs where code=?),
			?,
			?,
			(select id from bi_nomens where code=?),
			(select id from bi_nomens where code=?),
			?,
			?,
			?,
			?,
			?,
			?,
			(select id from bi_nomens where code=?),
			?,
			(select id from bi_deals where code=?),
			?,
			? )`;
		return o.Raw(sql,
			element["code"],
			element["doc_at"],
			element["account_code"],
			element["ind_site_code"],
			element["deal_code"],
			element["seal"],
			element["is_central"],
			element["delivery_address_code"],
			element["driver_code"],
			element["apt_account_code"],
			element["vehicle_code"],
			element["delivery_type"],
			element["departure_at"],
			element["beton_req_code"],
			element["shipped_quantity"],
			element["dscr"],
			element["mod_nomen_code"],
			element["addon_nomen_code"],
			element["owner_1c"],
			element["node"],
			element["is_cancel"],
			element["delivery_amount"],
			element["is_rеturned"],
			element["return_at"],
			element["individual_code"],
			element["delay_amount"],
			element["deal_logist_code"],
			element["recipe"],
			element["ported"],

		).Exec()
		/*                ****************       */
		/*                END BETON INVOICE DOCUMENT */
	}else{
		return nil,errors.New("entity "+entityCode+" not importable")
	}
	return nil,errors.New("entity "+entityCode+" not importable")
}
