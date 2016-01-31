package restapi
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/Shopify/go-lua"
	"fmt"
)

func test(L *lua.State) int {
	fmt.Println("hello world! from go!")
	return 0
}

func LuaTest(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	l := lua.NewState()
	lua.OpenLibraries(l)
	l.Register("test",test)


	if err := lua.DoFile(l, "lua/hello.lua"); err != nil {
		panic(err)
	}
}
