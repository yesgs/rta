package response

type TaobaoUsergrowthDhhDeliveryAskResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        在大航海平台可投放的任务ID列表
    */
    TaskIdList  []string `json:"task_id_list,omitempty" `
    /*
        错误码， 0： 成功；1：限流；2：服务不可用
    */
    Errcode  int64 `json:"errcode,omitempty" `
    /*
        true: 目标用户；false: 非目标用户
    */
    Result  bool `json:"result,omitempty" `
    /*
        在大航海平台推荐的任务ID
    */
    TaskId  string `json:"task_id,omitempty" `
}
