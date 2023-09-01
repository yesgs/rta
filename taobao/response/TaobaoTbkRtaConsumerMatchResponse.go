package response

import (
	"github.com/yesgs/rta/taobao/domain"
)

type TaobaoTbkRtaConsumerMatchResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回结果描述信息
	*/
	ResultMsg string `json:"result_msg,omitempty" `
	/*
	   返回结果封装对象
	*/
	Data domain.TaobaoTbkRtaConsumerMatchData `json:"data,omitempty" `
}
