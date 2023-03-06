package apis

type GetBlockInfoReq struct {
	Hex string `json:"hex"`
}

type SubscriptionReq struct {
	Address string `json:"address"`
}

type GetTransactionsRequest struct {
	Address  string   `json:"address"`
	PageInfo PageInfo `json:"page_info"`
}

type PageInfo struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"page_size"`
	Total    int64 `json:"total"`
	Desc     bool  `json:"desc"`
}

type Response struct {
	ErrMsg   string      `json:"error_message,omitempty"`
	Data     interface{} `json:"data"`
	PageInfo *PageInfo   `json:"page_info,omitempty"`
}
