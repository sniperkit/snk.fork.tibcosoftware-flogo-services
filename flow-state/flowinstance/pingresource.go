/*
Sniperkit-Bot
- Status: analyzed
*/

package flowinstance

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Ping(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	fmt.Fprintf(response, "%s", "{\"status\":\"ok\"}")
}
