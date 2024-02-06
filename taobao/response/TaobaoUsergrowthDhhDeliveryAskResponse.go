package response

type TaobaoUsergrowthDhhDeliveryAskResponse struct {
	Result    *Result `json:"result"`
	RequestId string  `json:"request_id"`
}

type Result struct {
	Errcode int          `json:"errcode"`
	Results []ResultItem `json:"results"`
}

type ResultItem struct {
	OaidMd5    string   `json:"oaid_md5"`
	IdfaMd5    string   `json:"idfa_md5"`
	CaidMd5    string   `json:"caid_md5"`
	ImeiMd5    string   `json:"imei_md5"`
	TaskIdList []string `json:"task_id_list"`
}
