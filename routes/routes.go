package routes


import "net/http"
import (
	"github.com/yeldars/crm/restapi"
	"github.com/yeldars/crm/file"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"log"
	"os"
	"github.com/yeldars/crm/auth"
	"github.com/astaxie/beego/orm"
	"github.com/yeldars/crm/utils"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func HandleInit(){

	router := httprouter.New()
	router.GET("/",RedirectHome)
	router.ServeFiles("/static/*filepath",http.Dir("static"))
	router.GET("/restapi/accounts/get",restapi.AccountRestApiGet)
	router.POST("/restapi/accounts/insert",restapi.AccountRestApiInsert)
	router.POST("/restapi/accounts/update",restapi.AccountRestApiUpdate)
	router.POST("/restapi/accounts/delete",restapi.AccountRestApiDelete)
	router.GET("/restapi/accounts/detail",restapi.AccountRestApiDetail)
	router.GET("/restapi/accounts/removeall",restapi.AccountRestApiRemoveAll)

	router.GET("/restapi/query/get",restapi.QueryRestApiGet)
	router.POST("/restapi/update",restapi.UpdateRestApi)
	router.POST("/restapi/update_v_1_1",restapi.UpdateRestApi_v_1_1)
	router.GET("/restapi/detail",restapi.DetailRestApi)
	router.GET("/restapi/removeall",restapi.RemoveAll)

	router.GET("/restapi/deals/get",restapi.DealRestApiGet)
	router.GET("/restapi/deals/detail",restapi.DealRestApiDetail)
	router.POST("/restapi/deals/runop",restapi.DealRestApiRunOp)
	router.GET("/restapi/deals/takeone",restapi.DealRestApiTakeOne)
	router.GET("/restapi/deals/removeall",restapi.DealRestApiRemoveAll)
	router.POST("/restapi/deals/insert",restapi.DealRestApiInsert)
	router.POST("/restapi/deals/update",restapi.DealRestApiUpdate)

	router.GET("/restapi/waitcall",restapi.WaitCallRestApiGet)
	router.GET("/restapi/incomingcall",restapi.IncomingCallRestApiGet)


	router.GET("/userpic",file.UserPic)
	router.POST("/auth/login",auth.Login)
	router.GET("/auth/session_info",auth.GetSessionInfo)
	router.POST("/auth/resetpassword",auth.ResetPassword)
	router.GET("/auth/logout",auth.Logout)

	router.GET("/auth/getlanguage",auth.GetLanguage)
	router.GET("/auth/setlanguage",auth.SetLanguage)

	router.GET("/restapi/luatest",restapi.LuaTest)
	router.GET("/restapi/emailtest",restapi.EmailTest)
	router.POST("/cdr",restapi.Cdr)

	router.GET("/restapi/translates/get",restapi.TranslateRestApiGet)
	router.GET("/restapi/pages/get",restapi.PageRestApiGet)
	router.GET("/restapi/pagetemplate",restapi.PageRestApiGetPageTemplate)
	router.GET("/restapi/widgettemplate",restapi.WidgetRestApiGetWidgetTemplate)

	router.GET("/showpage",restapi.PageRestApiGetPageTemplate)

	router.GET("/restapi/menus/tree",restapi.MenuRestApiGetTree)
	router.GET("/restapi/widget/get",restapi.WidgetRestApiGet)
	router.POST("/upload",file.Upload)
	router.GET("/exportall",file.ExportAll)

	bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
	//err := http.ListenAndServe(bind, nil)
	log.Fatal(http.ListenAndServe(bind, router))



}


func RedirectHome(res http.ResponseWriter, req *http.Request, _ httprouter.Params){

	o := orm.NewOrm()
	o.Using("default")



	if auth.UserId(req)!=0 {
		http.Redirect(res, req, utils.GetDomainParamValue(req.Host,"homepage"), 301)
	}else{
		http.Redirect(res, req, utils.GetDomainParamValue(req.Host,"loginpage"), 301)
	}
}