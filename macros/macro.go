package macros
import (
	"github.com/tealeg/xlsx"
//	"strconv"
	"strconv"
)

func RunMacro(templateId int, pk int64, row *xlsx.Row){
	if templateId == 1{
		str1,_ := row.Cells[17].String()
		str2,_ := row.Cells[18].String()
		pk := []string  {strconv.Itoa(int(pk)), str1,str2}
		CreateAccountBy2Phone(pk)
	}

}
