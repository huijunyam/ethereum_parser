package apis

import (
	"encoding/json"
	"errors"
	"ethereum_parser/dal"
	"net/http"
)

func getTransactionsByAddress(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		bindErrorResponse(w, errors.New("invalid request method"), http.StatusMethodNotAllowed)
		return
	}
	var tr GetTransactionsRequest
	tr.PageInfo = PageInfo{
		Page:     1,
		PageSize: 10,
	}
	err := json.NewDecoder(r.Body).Decode(&tr)
	if err != nil {
		bindErrorResponse(w, errors.New("address is required"), http.StatusBadRequest)
		return
	}

	if tr.Address == "" {
		bindErrorResponse(w, errors.New("address cannot be empty"), http.StatusBadRequest)
		return
	}

	offset := (tr.PageInfo.Page - 1) * tr.PageInfo.PageSize
	transactionList, total, err := dal.StorageClient.GetTransactionList(tr.Address, offset, offset+tr.PageInfo.PageSize)
	if err != nil {
		bindErrorResponse(w, err, http.StatusInternalServerError)
	}
	tr.PageInfo.Total = total
	bindResponse(w, transactionList, &tr.PageInfo)
}
