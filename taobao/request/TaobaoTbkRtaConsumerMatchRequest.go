package request

type TaobaoTbkRtaConsumerMatchRequest struct {
	AdZoneId       *int64  `json:"adzone_id,omitempty" required:"false" `
	SpecialId      *string `json:"special_id,omitempty" required:"false" `
	DeviceValue    *string `json:"device_value,omitempty" required:"false" `
	DeviceType     *string `json:"device_type,omitempty" required:"false" `
	StrategyIdList *string `json:"strategy_id_list,omitempty" required:"false" `
}

func (req *TaobaoTbkRtaConsumerMatchRequest) SetAdZoneId(v int64) *TaobaoTbkRtaConsumerMatchRequest {
	req.AdZoneId = &v
	return req
}
func (req *TaobaoTbkRtaConsumerMatchRequest) SetSpecialId(v string) *TaobaoTbkRtaConsumerMatchRequest {
	req.SpecialId = &v
	return req
}
func (req *TaobaoTbkRtaConsumerMatchRequest) SetDeviceValue(v string) *TaobaoTbkRtaConsumerMatchRequest {
	req.DeviceValue = &v
	return req
}
func (req *TaobaoTbkRtaConsumerMatchRequest) SetDeviceType(v string) *TaobaoTbkRtaConsumerMatchRequest {
	req.DeviceType = &v
	return req
}
func (req *TaobaoTbkRtaConsumerMatchRequest) SetStrategyIdList(v string) *TaobaoTbkRtaConsumerMatchRequest {
	req.StrategyIdList = &v
	return req
}

func (req *TaobaoTbkRtaConsumerMatchRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.AdZoneId != nil {
		paramMap["adzone_id"] = *req.AdZoneId
	}
	if req.SpecialId != nil {
		paramMap["special_id"] = *req.SpecialId
	}
	if req.DeviceValue != nil {
		paramMap["device_value"] = *req.DeviceValue
	}
	if req.DeviceType != nil {
		paramMap["device_type"] = *req.DeviceType
	}
	if req.StrategyIdList != nil {
		paramMap["strategy_id_list"] = *req.StrategyIdList
	}
	return paramMap
}
