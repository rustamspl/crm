package file
import (
	"net/http"
	"fmt"
	"os"
	"github.com/astaxie/beego/orm"
	"log"
	"encoding/json"
	"github.com/tealeg/xlsx"
	"github.com/yeldars/crm/restapi"
)

func ImportUniversal(w http.ResponseWriter, r *http.Request, fileName string) (error) {

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


	o := orm.NewOrm()
	o.Using("default")

	dealStageId := 0
	o.Raw("select id from deal_stages where code='notassigned'").QueryRow(&dealStageId)

	//accId := int64(0)




	xlFile, err := xlsx.OpenFile(fileName)
	restapi.CheckPanic(err)


	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {


			_,err := o.Raw("insert into deals "+
			"(deal_stage_id,x_object_type,x_deal_type,x_city,x_room_count,x_region,x_address,"+
			"x_object_price,x_price_meter,x_sq1,x_sq2,x_sq3,x_floor,x_floor_count,x_sost,x_plan,x_year,x_mebel,amount,title) "+
			" values ("+
			"(select id from deal_stages where code='notassigned'),"+
			"(select id from list_values lv where lv.value=? and lv.list_id=4),"+
			"(select id from list_values lv where lv.value=? and lv.list_id=3),"+
			"(select id from list_values lv where lv.value=? and lv.list_id=1),"+
			"?,"+
			"(select id from list_values lv where lv.value=? and lv.list_id=2),"+
			"?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
				row.Cells[0].String(),
				row.Cells[1].String(),
				row.Cells[2].String(),
				row.Cells[3].String(),
				row.Cells[4].String(),
				row.Cells[5].String(),
				row.Cells[6].String(),
				row.Cells[7].String(),
				row.Cells[8].String(),
				row.Cells[9].String(),
				row.Cells[10].String(),
				row.Cells[11].String(),
				row.Cells[12].String(),
				row.Cells[13].String(),
				row.Cells[14].String(),
				row.Cells[15].String(),
				row.Cells[16].String(),
				row.Cells[19].String(),
				row.Cells[20].String(),
			).Exec()

			if err==nil{
				result.OkCnt ++
			}else{
				result.ErrCnt++
				log.Println(err)
			}


			//row.Cells

			/*var element models.Deals

			if len(row.Cells)>6 {

				objType := row.Cells[0].String()
				dealType := row.Cells[1].String()
				city := row.Cells[2].String()
				region := row.Cells[3].String()
				dscR := row.Cells[4].String()
				amount,err := row.Cells[5].Float()
				phone := row.Cells[6].String()
				phone = strings.Replace(phone, "(", "", -1);
				phone = strings.Replace(phone, ")", "", -1);
				phone = strings.Replace(phone, " ", "", -1);
				phone2 := ""

				if len(row.Cells) > 7 {
					phone2 = row.Cells[7].String()
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
				element.Amount = amount

				element.XRegion = regionValue
				element.XDscr =dscR

				o.Insert(&element)
				if err==nil{
					result.OkCnt ++
				}else{
					result.ErrCnt++
					log.Println(err)
				}
			}*/
		}
	}


	result.Result="ok"

	jsonData, err := json.Marshal(result)
	fmt.Fprint(w,string(jsonData))

	return  err
	//return lines, scanner.Err()
}
