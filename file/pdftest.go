package file
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/jung-kurt/gofpdf"
)


func TestPDF(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {


	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("hello.pdf")
	if err!=nil{
		panic(err)
	}
}

