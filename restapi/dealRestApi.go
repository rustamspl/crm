package restapi
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/astaxie/beego/orm"
	"strconv"
	"fmt"
	"encoding/json"
	"github.com/yeldars/crm/models"


	"github.com/yeldars/crm/auth"
	"log"

)

type dealInsertRequest struct {
	Items [] models.Deals `json:"items"`
}


type DealsList struct {
	Id          int64 `orm:"auto" json:"id"`
	Title        string `json:"title"`
	Dscr string `json:"dscr"`
	Amount float64 `json:"amount"`
	Account string `json:"account"`
	AccountId int64 `json:"account_id"`
	DealStage string `json:"deal_stage"`
	City string `json:"city"`
	Region string `json:"region"`
	DealType string `json:"deal_type"`
	ObjectType string `json:"object_type"`
	Category string `json:"category"`
	DealStageColor string `json:"deal_stage_color"`
	OwnerName string `json:"owner_name"`
	OwnerID int64 `json:"owner_id"`
}



type dealsGetResponse struct {
	PageCount   int `json:"pageCount"`
	Items [] DealsList `json:"items"`
}

type DealOp struct {
	Id           int64 `orm:"auto" json:"id"`
	Title        string `json:"title"`
	OperCode        string `json:"oper_code"`
	FromStageId  int64 `json:"from_stage_id"`
	ToStageId    int64 `json:"to_stage_id"`
	FormUrl      string  `json:"form_url"`
	Attrs map[string]string
}

type DealOpJrnList struct {
	Id           int64 `orm:"auto" json:"id"`
	Title        string `json:"title"`
	OpId        int64 `json:"op_id"`
	DealId  int64 `json:"deal_id"`
	UserId    int64 `json:"user_id"`
}
type Attrs struct {
	DateReceipt string `json:"date_receipt"`
}
type DealOpRun struct {
	DealId          int64 `json:"deal_id"`
	UserId          int64 `json:"user_id"`
	JrnId 			int64 `json:"jrn_id"`
	Op DealOp `json:"op"`

}



type DealsDetails struct {
	Id          int64 `orm:"auto" json:"id"`
	Title        string `json:"title"`
	AccountId int64 `json:"account_id"`
	Account string `json:"account"`
	Amount float64 `json:"amount"`
	DealStageId int64 `json:"deal_stage_id"`
	DealStage string `json:"deal_stage"`
	XCity int64 `json:"x_city"`
	XCityTitle string `json:"x_city_title"`
	XRegion int64 `json:"x_region"`
	XRegionTitle string `json:"x_region_title"`
	XObjectType int64 `json:"x_object_type"`
	XObjectTypeTitle string `json:"x_object_type_title"`
	XObjectPrice string `json:"x_object_price"`
	XObjectSquare string `json:"x_object_square"`
	OwnerId int64 `json:"owner_id"`
	Owner string `json:"owner"`
	Ops [] DealOp `json:"ops"`
	OpJrns [] DealOpJrnList `json:"opjrns"`
}


func DealRunOper (t * DealOpRun){


	o := orm.NewOrm()
	o.Using("default")


	log.Print("before t.op.id=")
	log.Print(t.Op.Id)

	err := o.Raw("SELECT * from stage_opers so where so.id=?",t.Op.Id).QueryRow(&t.Op)
	log.Print("t.op.id=")
	log.Print(t.Op.Id)
	CheckPanic(err)
	o.QueryTable("deals").Filter("id", t.DealId).Update(orm.Params{
		"deal_stage_id":t.Op.ToStageId,
		"owner_id":t.UserId,
	})
	CheckPanic(err)
	ops := o.QueryTable("stage_op_jrns")
	i,_ := ops.PrepareInsert()
	var stageOpJrn = models.StageOpJrns{}
	stageOpJrn.OpId = t.Op.Id
	stageOpJrn.UserId = t.UserId
	stageOpJrn.DealId =t.DealId
	i.Insert(&stageOpJrn)
	//log.Println("stageOpJrn")
	//log.Println(stageOpJrn.Id)
	CheckPanic(err)
	t.JrnId = stageOpJrn.Id
	run(t)
}

func DealRestApiRunOp(res http.ResponseWriter, req *http.Request, _ httprouter.Params){

	if RestCheckAuth(res,req){
		return
	}
	decoder := json.NewDecoder(req.Body)
	var t DealOpRun
	t.Op.Attrs = make(map[string]string)


	err := decoder.Decode(&t)
	CheckPanic(err)
	//t.Attrs = make(map[string]string)
	//err := json.Unmarshal(data, &objmap)




	t.UserId = auth.UserId(req)
	DealRunOper(&t)
	fmt.Fprint(res,"{}")

}

func DealRestApiDetail(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	req.ParseForm()
	item  := DealsDetails{}
	o := orm.NewOrm()
	o.Using("default")
	err := o.Raw("SELECT * from deals a where a.id=?",req.Form.Get("id")).QueryRow(&item)
	if RestCheckDBPanic(err ,res ,o ) {
		return
	}
	_, err = o.Raw("SELECT id,title,form_url from stage_opers cg where from_stage_id=?",item.DealStageId).QueryRows(&item.Ops)
	if RestCheckDBPanic(err ,res ,o ) {
		return
	}
	_, err = o.Raw("SELECT so.title from stage_opers so,stage_op_jrns soj where soj.op_id=so.id and soj.deal_id=?",item.Id).QueryRows(&item.OpJrns)
	if RestCheckDBPanic(err ,res ,o ) {
		return
	}
	err = o.Raw("SELECT title from deal_stages ds where id=?",item.DealStageId).QueryRow(&item.DealStage)
	if RestCheckDBPanic(err ,res ,o ) {
		return
	}
	err = o.Raw("SELECT value from list_values ds where id=?",item.XCity).QueryRow(&item.XCityTitle)
	if RestCheckDBPanic(err ,res ,o ) {
		return
	}
	err = o.Raw("SELECT value from list_values ds where id=?",item.XRegion).QueryRow(&item.XRegionTitle)
	if RestCheckDBPanic(err ,res ,o ) {
		return
	}
	err = o.Raw("SELECT value from list_values ds where id=?",item.XObjectType).QueryRow(&item.XObjectTypeTitle)
	if RestCheckDBPanic(err ,res ,o ) {
		return
	}
	err = o.Raw("SELECT name from users a where id=?",item.OwnerId).QueryRow(&item.Owner)
	if RestCheckDBPanic(err ,res ,o ) {
		return
	}
	jsonData, err := json.Marshal(item)
	if RestCheckDBPanic(err ,res ,o ) {
		return
	}
	fmt.Fprint(res,string(jsonData))

}


func DealRestApiTakeOne(res http.ResponseWriter, req *http.Request, _ httprouter.Params){

	if RestCheckAuth(res,req){
		return
	}

	o := orm.NewOrm()
	o.Using("default")


	id := int64(0)



	err := o.Raw("SELECT id FROM deals where deal_stage_id=(select id from deal_stages where code='prolonged') and now()>=x_prolong_datetime  limit 1").QueryRow(&id)
	if err==nil {
		log.Println("vata emes2?")
		var t DealOpRun
		t.UserId = auth.UserId(req)
		t.DealId = id
		t.Op.Id =  getOperIdByCode("return_prolong_to_work")
		log.Println("before prolonged")
		log.Println(id)
		DealRunOper(&t)
		log.Println("after prolonged")

		log.Println(id)
		fmt.Fprint(res,"{\"id\":"+strconv.Itoa(int(id))+"}")
		return
	}
	err = o.Raw("SELECT id FROM deals where `owner_id`=? and deal_stage_id=(select id from deal_stages where code='process') limit 1",auth.UserId(req)).QueryRow(&id)
	if err==nil {
		log.Println("vata emes3?")
		log.Println(id)
		fmt.Fprint(res,"{\"id\":"+strconv.Itoa(int(id))+"}")
		return
	}
	log.Println("before1")
	err = o.Raw("SELECT id FROM deals where deal_stage_id=(select id from deal_stages where code='notassigned') limit 1").QueryRow(&id)
	log.Println("after1")
	if err == nil {

		var t DealOpRun

		t.UserId = auth.UserId(req)
		t.DealId = id
		t.Op.Id = getOperIdByCode("takework")
		log.Println("before takework")
		log.Println(id)
		DealRunOper(&t)
		log.Println("after takework")
		fmt.Fprint(res, "{\"id\":" + strconv.Itoa(int(id)) + "}")
	}

}
func DealRestApiGet(res http.ResponseWriter, req *http.Request, _ httprouter.Params){

	if RestCheckAuth(res,req){
		return
	}

	o := orm.NewOrm()
	o.Using("default")
	var arr [] DealsList

	grantFilter := auth.GetSqlGrantByTableName(auth.UserId(req),"deals")
	log.Println(grantFilter)
	if grantFilter == ""{
		grantFilter = "(1 = 1)"
	}

	req.ParseForm()
	limitFrom := "1"
	limitTo := "5"

	iPerPage,_ := strconv.ParseInt(req.Form.Get("perpage"),10,32);
	iPage,_ := strconv.ParseInt(req.Form.Get("page"),10,32);

	limitFrom = strconv.Itoa(int(iPerPage*(iPage-1)))
	limitTo = req.Form.Get("perpage");

	pageCount := 10;
	o.Raw("SELECT ceil(count(1)/?) FROM deals d where "+grantFilter,limitTo).QueryRow(&pageCount)



	_, err := o.Raw("SELECT d.*" +
	",(select name from accounts a where a.id=d.account_id) account"+
	",(select name from users u where u.id=d.owner_id) owner_name"+
	",(select s.title from deal_stages s where s.id=d.deal_stage_id) deal_stage"+
	",(select s.color from deal_stages s where s.id=d.deal_stage_id) deal_stage_color"+
	",(select lv.value from list_values lv where lv.id=d.x_city) city"+
	",(select lv.value from list_values lv where lv.id=d.x_region) region"+
	",(select lv.value from list_values lv where lv.id=d.x_deal_type) deal_type"+
	",(select lv.value from list_values lv where lv.id=d.x_object_type) object_type"+
	",(select lv.value from list_values lv where lv.id=d.x_deal_cat) deal_cat"+
	" FROM deals d  where "+grantFilter+" limit ?,?",limitFrom,limitTo).QueryRows(&arr)

	if RestCheckDBPanic(err ,res ,o ) {
		return
	}
	respO := dealsGetResponse{}
	respO.Items = arr
	respO.PageCount = pageCount
	jsonData, err := json.Marshal(respO)
	if RestCheckDBPanic(err ,res ,o ) {
		return
	}
	fmt.Fprint(res,string(jsonData))

}

func DealRestApiRemoveAll(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if RestCheckAuth(res,req){
		return
	}

   	o := orm.NewOrm()
	o.Using("default")
	o.Raw("delete from deals").Exec()
}


func DealRestApiInsert(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	if RestCheckAuth(res,req){
		return
	}

	decoder := json.NewDecoder(req.Body)
	var t dealInsertRequest
	err := decoder.Decode(&t)
	if RestCheckPanic(err ,res ) {
		return
	}

	o := orm.NewOrm()
	o.Using("default")
	o.Begin()


	qs := o.QueryTable("deals")
	i, _ := qs.PrepareInsert()
	for index, element := range  t.Items {
		element.Id = 0;
		id,err := i.Insert(&element)
		element.Id=id

		t.Items[index].Id=id

		fmt.Printf("%d",id)
		if err != nil{
			o.Rollback()
			panic(err)
		}
	}
	i.Close()
	o.Commit()

	j,_ := json.Marshal(t)
	fmt.Fprint(res,string(j))
}

type dealUpdateRequest struct {
	Items [] models.Deals`json:"items"`
}
func DealRestApiUpdate(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if RestCheckAuth(res,req){
		return
	}

	decoder := json.NewDecoder(req.Body)
	var t dealUpdateRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	o := orm.NewOrm()
	o.Using("default")
	o.Begin()

	for _,element := range t.Items {

		o.QueryTable("deals").Filter("id", element.Id).Update(orm.Params{
			"title":element.Title,
			"amount": element.Amount,
		})

	}

	o.Commit()

}