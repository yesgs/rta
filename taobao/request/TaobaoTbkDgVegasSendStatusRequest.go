package request

type TaobaoTbkDgVegasSendStatusRequest struct {
	/*
	   渠道管理id     */
	RelationId *string `json:"relation_id,omitempty" required:"false" `
	/*
	   会员运营id     */
	SpecialId *string `json:"special_id,omitempty" required:"false" `
	/*
	   加密后的值(ALIPAY_ID除外)，需用MD5加密，32位小写     */
	DeviceValue *string `json:"device_value,omitempty" required:"false" `
	/*
	   入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE 5. ALIPAY_ID     */
	DeviceType *string `json:"device_type,omitempty" required:"false" `
	/*
	   thor平台业务码， 1：coupon 超红     */
	ActivityCategory *int64 `json:"activity_category,omitempty" required:"false" `
	/*
	   媒体pid     */
	Pid *string `json:"pid,omitempty" required:"false" `
}

func (req *TaobaoTbkDgVegasSendStatusRequest) SetRelationId(v string) *TaobaoTbkDgVegasSendStatusRequest {
	req.RelationId = &v
	return req
}
func (req *TaobaoTbkDgVegasSendStatusRequest) SetSpecialId(v string) *TaobaoTbkDgVegasSendStatusRequest {
	req.SpecialId = &v
	return req
}
func (req *TaobaoTbkDgVegasSendStatusRequest) SetDeviceValue(v string) *TaobaoTbkDgVegasSendStatusRequest {
	req.DeviceValue = &v
	return req
}
func (req *TaobaoTbkDgVegasSendStatusRequest) SetDeviceType(v string) *TaobaoTbkDgVegasSendStatusRequest {
	req.DeviceType = &v
	return req
}
func (req *TaobaoTbkDgVegasSendStatusRequest) SetActivityCategory(v int64) *TaobaoTbkDgVegasSendStatusRequest {
	req.ActivityCategory = &v
	return req
}
func (req *TaobaoTbkDgVegasSendStatusRequest) SetPid(v string) *TaobaoTbkDgVegasSendStatusRequest {
	req.Pid = &v
	return req
}

func (req *TaobaoTbkDgVegasSendStatusRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RelationId != nil {
		paramMap["relation_id"] = *req.RelationId
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
	if req.ActivityCategory != nil {
		paramMap["activity_category"] = *req.ActivityCategory
	}
	if req.Pid != nil {
		paramMap["pid"] = *req.Pid
	}
	return paramMap
}

func (req *TaobaoTbkDgVegasSendStatusRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
