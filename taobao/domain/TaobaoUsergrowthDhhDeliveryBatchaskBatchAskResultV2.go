package domain


type TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultV2 struct {
    /*
        错误码， 0： 成功；1：限流；2：服务不可用     */
    Errcode  *int64 `json:"errcode,omitempty" `

    /*
        匹配的设备与其任务信息列表     */
    Results  *[]TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultItem `json:"results,omitempty" `

}

func (s *TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultV2) SetErrcode(v int64) *TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultV2 {
    s.Errcode = &v
    return s
}
func (s *TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultV2) SetResults(v []TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultItem) *TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultV2 {
    s.Results = &v
    return s
}
