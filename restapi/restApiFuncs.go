package restapi
import (
	"net/http"
	"github.com/yeldars/crm/auth"
)

func IsMobile (req *http.Request) bool {
	return auth.System(req) == "android" || auth.System(req) == "ios"
}
