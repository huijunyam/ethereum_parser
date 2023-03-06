package apis

import (
	"encoding/json"
	"errors"
	"ethereum_parser/helper"
	"ethereum_parser/json_rpc"
	"net/http"
)

func getCurrentBlock(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		bindErrorResponse(w, errors.New("invalid request method"), http.StatusMethodNotAllowed)
		return
	}

	hexVal, err := json_rpc.GetCurrentBlock(helper.GenId())
	if err != nil {
		bindErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	decimalVal, err := helper.ConvertHexStringToInt(*hexVal)
	if err != nil {
		bindErrorResponse(w, err, http.StatusBadRequest)
	}
	bindResponse(w, decimalVal, nil)
}

func getBlockInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		bindErrorResponse(w, errors.New("invalid request method"), http.StatusMethodNotAllowed)
		return
	}
	var b GetBlockInfoReq
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		bindErrorResponse(w, errors.New("block number is required"), http.StatusBadRequest)
		return
	}

	if b.Hex == "" {
		bindErrorResponse(w, errors.New("block number cannot be empty"), http.StatusBadRequest)
		return
	}

	blockInfo, err := json_rpc.GetBlockByNumber(helper.GenId(), b.Hex, true)
	if err != nil {
		bindErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	bindResponse(w, blockInfo, nil)
}
