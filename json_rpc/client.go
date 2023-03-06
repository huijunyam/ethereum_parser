package json_rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"ethereum_parser/models"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	url                       = "https://cloudflare-eth.com"
	jsonRpcVersion            = "2.0"
	EthBlockNumberMethod      = "eth_blockNumber"
	EthGetBlockByNumberMethod = "eth_getBlockByNumber"
)

type JsonRpcRequest struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Id      string        `json:"id"`
	Params  []interface{} `json:"params"`
}

type JsonRpcResponse struct {
	JsonRpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Id      string      `json:"id"`
}

func GetCurrentBlock(id string) (*string, error) {
	reqBody, err := json.Marshal(JsonRpcRequest{
		JsonRpc: jsonRpcVersion,
		Method:  EthBlockNumberMethod,
		Id:      id,
	})
	if err != nil {
		return nil, err
	}
	resBody, err := postRequest(url, reqBody)
	if err != nil {
		return nil, err
	}
	var res JsonRpcResponse
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, err
	}
	field, ok := res.Result.(string)
	if !ok {
		return nil, errors.New("unable to convert field type")
	}
	return &field, nil
}

func GetBlockByNumber(id, blockNum string, withTransaction bool) (models.BlockInfo, error) {
	reqBody, err := json.Marshal(JsonRpcRequest{
		JsonRpc: jsonRpcVersion,
		Method:  EthGetBlockByNumberMethod,
		Id:      id,
		Params:  []interface{}{blockNum, withTransaction},
	})
	if err != nil {
		return models.BlockInfo{}, err
	}
	resBody, err := postRequest(url, reqBody)
	if err != nil {
		return models.BlockInfo{}, err
	}
	var res JsonRpcResponse
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return models.BlockInfo{}, err
	}
	blockInfo := models.BlockInfo{}
	resultByteData, err := json.Marshal(res.Result)
	err = json.Unmarshal(resultByteData, &blockInfo)
	if err != nil {
		return models.BlockInfo{}, errors.New("unable to convert block info type")
	}
	return blockInfo, nil
}

func postRequest(url string, reqBody []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Printf("client: could not create request: %s\n", err)
		return nil, err
	}

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("client: error making http request: %s\n", err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("client: error getting response body: %s\n", err)
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return body, err
	}
	return body, nil
}
