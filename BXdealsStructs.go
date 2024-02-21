package bx_worker_go

type CrmDealHistoryResponse struct {
	Result struct {
		Items []interface{} `json:"items"`
	} `json:"result"`
	Total int `json:"total"`
}
