package request


type TaobaoUsergrowthDhhDeliveryBatchaskRequest struct {
    /*
        md5加密后的oaid列表， 32位小写， 多个使用,分隔， 最多支持20个     */
    OaidMd5  *string `json:"oaid_md5,omitempty" required:"false" `
    /*
        md5加密后的oaid列表， 32位小写， 多个使用,分隔， 最多支持20个     */
    IdfaMd5  *string `json:"idfa_md5,omitempty" required:"false" `
    /*
        md5加密后的imei列表， 32位小写， 多个使用,分隔， 最多支持20个     */
    ImeiMd5  *string `json:"imei_md5,omitempty" required:"false" `
    /*
        巨浪广告位id,在巨浪平台申请     */
    AdvertisingSpaceId  *string `json:"advertising_space_id" required:"true" `
    /*
        巨浪渠道id,在巨浪平台申请     */
    Channel  *string `json:"channel" required:"true" `
}

func (s *TaobaoUsergrowthDhhDeliveryBatchaskRequest) SetOaidMd5(v string) *TaobaoUsergrowthDhhDeliveryBatchaskRequest {
    s.OaidMd5 = &v
    return s
}
func (s *TaobaoUsergrowthDhhDeliveryBatchaskRequest) SetIdfaMd5(v string) *TaobaoUsergrowthDhhDeliveryBatchaskRequest {
    s.IdfaMd5 = &v
    return s
}
func (s *TaobaoUsergrowthDhhDeliveryBatchaskRequest) SetImeiMd5(v string) *TaobaoUsergrowthDhhDeliveryBatchaskRequest {
    s.ImeiMd5 = &v
    return s
}
func (s *TaobaoUsergrowthDhhDeliveryBatchaskRequest) SetAdvertisingSpaceId(v string) *TaobaoUsergrowthDhhDeliveryBatchaskRequest {
    s.AdvertisingSpaceId = &v
    return s
}
func (s *TaobaoUsergrowthDhhDeliveryBatchaskRequest) SetChannel(v string) *TaobaoUsergrowthDhhDeliveryBatchaskRequest {
    s.Channel = &v
    return s
}

func (req *TaobaoUsergrowthDhhDeliveryBatchaskRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OaidMd5 != nil) {
        paramMap["oaid_md5"] = *req.OaidMd5
    }
    if(req.IdfaMd5 != nil) {
        paramMap["idfa_md5"] = *req.IdfaMd5
    }
    if(req.ImeiMd5 != nil) {
        paramMap["imei_md5"] = *req.ImeiMd5
    }
    if(req.AdvertisingSpaceId != nil) {
        paramMap["advertising_space_id"] = *req.AdvertisingSpaceId
    }
    if(req.Channel != nil) {
        paramMap["channel"] = *req.Channel
    }
    return paramMap
}

func (req *TaobaoUsergrowthDhhDeliveryBatchaskRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}