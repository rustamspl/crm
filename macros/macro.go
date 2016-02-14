package macros
import (
	"github.com/tealeg/xlsx"
	"strconv"
)

func RunMacro(templateId int, pk int64, row *xlsx.Row){
	if templateId == 1{
		pk := []string  {strconv.Itoa(int(pk)), row.Cells[17].String(),row.Cells[18].String()}
		CreateAccountBy2Phone(pk)
	}

}
