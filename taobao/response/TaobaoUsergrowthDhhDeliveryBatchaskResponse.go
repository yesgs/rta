package response

import (
	"github.com/yesgs/rta/taobao/domain"
)

type TaobaoUsergrowthDhhDeliveryBatchaskResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回值
	*/
	Result domain.TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultV2 `json:"result,omitempty" `
}
