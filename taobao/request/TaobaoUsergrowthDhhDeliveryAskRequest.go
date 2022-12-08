package request


type TaobaoUsergrowthDhhDeliveryAskRequest struct {
    /*
        预留json参数，与手淘团队单独沟通     */
    Profile  *string `json:"profile,omitempty" required:"false" `
    /*
        oaid的md5值， 32位小写     */
    OaidMd5  *string `json:"oaid_md5,omitempty" required:"false" `
    /*
        idfa的md5值， 32位小写     */
    IdfaMd5  *string `json:"idfa_md5,omitempty" required:"false" `
    /*
        imei的md5值， 32位小写     */
    ImeiMd5  *string `json:"imei_md5,omitempty" required:"false" `
    /*
        oaid的原生值     */
    Oaid  *string `json:"oaid,omitempty" required:"false" `
    /*
        idfa的原生值     */
    Idfa  *string `json:"idfa,omitempty" required:"false" `
    /*
        imei的原生值     */
    Imei  *string `json:"imei,omitempty" required:"false" `
    /*
        用户所使用设备的系统， 0： android, 1: ios, 2: windowsphone, 3: other     */
    Os  *string `json:"os,omitempty" required:"false" `
    /*
        大航海广告位，在大航海平台申请     */
    AdvertisingSpaceId  *string `json:"advertising_space_id" required:"true" `
    /*
        渠道标识，在大航海平台申请     */
    Channel  *string `json:"channel" required:"true" `
    /*
        md5加密后的caid， 32位小写，前面拼接上caid版本号，当前支持20220111、20211207版本     */
    CaidMd5  *string `json:"caid_md5,omitempty" required:"false" `
}

func (s *TaobaoUsergrowthDhhDeliveryAskRequest) SetProfile(v string) *TaobaoUsergrowthDhhDeliveryAskRequest {
    s.Profile = &v
    return s
}
func (s *TaobaoUsergrowthDhhDeliveryAskRequest) SetOaidMd5(v string) *TaobaoUsergrowthDhhDeliveryAskRequest {
    s.OaidMd5 = &v
    return s
}
func (s *TaobaoUsergrowthDhhDeliveryAskRequest) SetIdfaMd5(v string) *TaobaoUsergrowthDhhDeliveryAskRequest {
    s.IdfaMd5 = &v
    return s
}
func (s *TaobaoUsergrowthDhhDeliveryAskRequest) SetImeiMd5(v string) *TaobaoUsergrowthDhhDeliveryAskRequest {
    s.ImeiMd5 = &v
    return s
}
func (s *TaobaoUsergrowthDhhDeliveryAskRequest) SetOaid(v string) *TaobaoUsergrowthDhhDeliveryAskRequest {
    s.Oaid = &v
    return s
}
func (s *TaobaoUsergrowthDhhDeliveryAskRequest) SetIdfa(v string) *TaobaoUsergrowthDhhDeliveryAskRequest {
    s.Idfa = &v
    return s
}
func (s *TaobaoUsergrowthDhhDeliveryAskRequest) SetImei(v string) *TaobaoUsergrowthDhhDeliveryAskRequest {
    s.Imei = &v
    return s
}
func (s *TaobaoUsergrowthDhhDeliveryAskRequest) SetOs(v string) *TaobaoUsergrowthDhhDeliveryAskRequest {
    s.Os = &v
    return s
}
func (s *TaobaoUsergrowthDhhDeliveryAskRequest) SetAdvertisingSpaceId(v string) *TaobaoUsergrowthDhhDeliveryAskRequest {
    s.AdvertisingSpaceId = &v
    return s
}
func (s *TaobaoUsergrowthDhhDeliveryAskRequest) SetChannel(v string) *TaobaoUsergrowthDhhDeliveryAskRequest {
    s.Channel = &v
    return s
}
func (s *TaobaoUsergrowthDhhDeliveryAskRequest) SetCaidMd5(v string) *TaobaoUsergrowthDhhDeliveryAskRequest {
    s.CaidMd5 = &v
    return s
}

func (req *TaobaoUsergrowthDhhDeliveryAskRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Profile != nil) {
        paramMap["profile"] = *req.Profile
    }
    if(req.OaidMd5 != nil) {
        paramMap["oaid_md5"] = *req.OaidMd5
    }
    if(req.IdfaMd5 != nil) {
        paramMap["idfa_md5"] = *req.IdfaMd5
    }
    if(req.ImeiMd5 != nil) {
        paramMap["imei_md5"] = *req.ImeiMd5
    }
    if(req.Oaid != nil) {
        paramMap["oaid"] = *req.Oaid
    }
    if(req.Idfa != nil) {
        paramMap["idfa"] = *req.Idfa
    }
    if(req.Imei != nil) {
        paramMap["imei"] = *req.Imei
    }
    if(req.Os != nil) {
        paramMap["os"] = *req.Os
    }
    if(req.AdvertisingSpaceId != nil) {
        paramMap["advertising_space_id"] = *req.AdvertisingSpaceId
    }
    if(req.Channel != nil) {
        paramMap["channel"] = *req.Channel
    }
    if(req.CaidMd5 != nil) {
        paramMap["caid_md5"] = *req.CaidMd5
    }
    return paramMap
}

func (req *TaobaoUsergrowthDhhDeliveryAskRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}