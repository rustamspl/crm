package utils
import "github.com/astaxie/beego/orm"

func GetDomainParamValue(domain string , param string ) string{
	o := orm.NewOrm()
	o.Using("default")
	val := ""
	err:=o.Raw("select dpv.value from domain_param_values dpv,params p,domains d where d.id=dpv.domain_id and p.id=dpv.param_id and d.domain=? and p.code=?",domain,param).QueryRow(&val)
	if err!=nil	{
		 return GetParamValue(param)
	}
	return  val
}


func GetParamValue(param string )  ( string  ){
	o := orm.NewOrm()
	o.Using("default")
	val := ""
	o.Raw("select p.value from params p where p.code=?",param).QueryRow(&val)
	return val
}



func GetUserParamValue(user int64 , param string ) ( string  ){
	o := orm.NewOrm()
	o.Using("default")
	val := ""
	err := o.Raw("select up.value from user_params up,params p where p.id=up.param_id and up.user_id=? and p.code=?",user,param).QueryRow(&val)
	if err!=nil	{
		return GetParamValue(param)
	}
	return  val
}
