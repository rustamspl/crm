package file
import (
	"net/http"
	"fmt"
	"bufio"
	"os"
	"strings"
	"github.com/astaxie/beego/orm"
	"main/models"
	"log"
	"encoding/json"
)

func ImportDeal01(w http.ResponseWriter, r *http.Request, fileName string) (error) {

   //fmt.Fprint(w,"{\"result\": \"importok\"}")

	type resultT struct {
		Result string `json:"result"`
		OkCnt   int `json:"ok_cnt"`
		ErrCnt  int `json:"err_cnt"`
		SkipCnt int	`json:"skip_cnt"`
	}

	var result resultT
	file, err := os.Open(fileName)
	if err != nil {
		return  err
	}
	defer file.Close()

	var lines []string

	o := orm.NewOrm()
	o.Using("default")

	dealStageId := 0
	o.Raw("select id from deal_stages where code='notassigned'").QueryRow(&dealStageId)

	accId := int64(0)



	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		ss := strings.Split(scanner.Text(),"\t")
		var element models.Deals
		if len(ss)>5 {

			objType := ss[0]
			dealType := ss[1]
			city := ss[2]
			region := ss[3]
			dscR := ss[4]
			phone := ss[5]
			phone = strings.Replace(phone, "(", "", -1);
			phone = strings.Replace(phone, ")", "", -1);
			phone = strings.Replace(phone, " ", "", -1);
			phone2 := ""

			if len(ss) > 6 {
				phone2 = ss[6]
				phone2 = strings.Replace(phone2, "(", "", -1);
				phone2 = strings.Replace(phone2, ")", "", -1);
				phone2 = strings.Replace(phone2, " ", "", -1);
			}


			cnt1 := 0
			cnt2 := 0
			o.Raw("select count(1) cnt from accountconts where cont=?",phone).QueryRow(&cnt1)
			o.Raw("select count(1) cnt from accountconts where cont=?",phone2).QueryRow(&cnt2)

			if cnt1+cnt2> 0 {
				log.Println("Propusk " + phone)
				result.SkipCnt ++
				continue
			} else{
				var acc models.Accounts
				acc.Name = "Клиент "+phone
				accId,_= o.Insert(&acc)

				var phoneT models.Accountconts
				phoneT.ContTypeId = 1
				phoneT.AccountId = accId
				phoneT.Cont = phone
				o.Insert(&phoneT)
				if phone2!="" {
					phoneT.Cont = phone2
					o.Insert(&phoneT)
				}

			}

			/*fmt.Fprintln(w, objType)
			fmt.Fprintln(w, dealType)
			fmt.Fprintln(w, city)
			fmt.Fprintln(w, region)
			fmt.Fprintln(w, dscR)
			fmt.Fprintln(w, phone)*/

			dealTypeValue := 0
			objTypeValue := 0
			regionValue := 0
			cityValue := 0

			o.Raw("select id from list_values where list_id=3 and value=?",dealType).QueryRow(&dealTypeValue)
			o.Raw("select id from list_values where list_id=4 and value=?",objType).QueryRow(&objTypeValue)
			o.Raw("select id from list_values where list_id=2 and value=?",region).QueryRow(&regionValue)
			o.Raw("select id from list_values where list_id=1 and value=?",city).QueryRow(&cityValue)


			element.Id = 0
			element.AccountId = accId
			element.Amount = 0
			element.DealStageId = dealStageId
			element.Title=dscR
			element.XDealType = dealTypeValue
			element.XObjectType = objTypeValue
			element.XCity = cityValue

			element.XRegion = regionValue
			element.XDscr =dscR

			o.Insert(&element)
			if err==nil{
				result.OkCnt ++
			}else{
				result.ErrCnt++
			}
		}
	}


	result.Result="ok"

	jsonData, _ := json.Marshal(result)
	fmt.Fprint(w,string(jsonData))

	return  scanner.Err()
	//return lines, scanner.Err()
}
