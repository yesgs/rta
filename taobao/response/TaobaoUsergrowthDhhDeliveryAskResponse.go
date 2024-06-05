package response

type TaobaoUsergrowthDhhDeliveryAskResponse struct {
	Errcode    int      `json:"errcode"`
	Result     bool     `json:"result"`
	RequestId  string   `json:"request_id"`
	TaskIdList []string `json:"task_id_list"`
}
