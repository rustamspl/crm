package auth
import (
	"github.com/astaxie/beego/orm"
	"strings"
	"strconv"
)

func GetSqlGrantByTableName(userId int64, tableName string) string {

	o := orm.NewOrm()
	o.Using("default")
	s := ""
	o.Raw("SELECT GROUP_CONCAT(concat('(',gs.sql_text,')') SEPARATOR ' and ') res  FROM `grant_sql_roles` gsr,`grant_sqls` gs, user_roles ur "+
"	where gs.id=gsr.sql_id and ur.role_id=gsr.role_id "+
"	and ur.user_id=?" ,userId).QueryRow(&s)
	s = strings.Replace(s,":user_id",strconv.Itoa(int(userId)),-1)
	return s
}
