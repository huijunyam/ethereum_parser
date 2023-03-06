package apis

import (
	"encoding/json"
	"net/http"
)

func bindResponse(w http.ResponseWriter, data interface{}, pageInfo *PageInfo) {
	resp := Response{
		Data:     data,
		PageInfo: nil,
	}
	if pageInfo != nil {
		resp.PageInfo = pageInfo
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func bindErrorResponse(w http.ResponseWriter, err error, errorCode int) {
	var resp Response
	if err != nil {
		resp.ErrMsg = err.Error()
	}
	w.WriteHeader(errorCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
