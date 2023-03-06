package apis

import (
	"net/http"
)

func getPing(w http.ResponseWriter, r *http.Request) {
	bindResponse(w, "OK!", nil)
}
