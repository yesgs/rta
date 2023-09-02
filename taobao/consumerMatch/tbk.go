package consumerMatch

import (
	"errors"
	"fmt"
	"github.com/yesgs/rta/taobao"
	"github.com/yesgs/rta/taobao/request"
	"github.com/yesgs/rta/taobao/response"
)

type Client struct {
	*taobao.Client
}

func NewClient(taobaoClient *taobao.Client) *Client {
	return &Client{taobaoClient}
}

func (c *Client) Ask(payload interface{}) (data []byte, err error) {
	return c.Invoke("taobao.tbk.rta.consumer.match", payload)
}

func (c *Client) ConvertRequest(body interface{}) (interface{}, error) {
	askRequest := body.(request.TaobaoTbkRtaConsumerMatchRequest)
	return askRequest.ToMap(), nil
}

func (c *Client) ConvertResponse(body []byte, output interface{}) (err error) {
	err = c.ExtractContent(body, output)
	if err != nil {
		return err
	}
	return c.ResponseHasBusinessError(output)
}

func (c *Client) ResponseHasBusinessError(body interface{}) error {
	switch body.(type) {
	case *response.TaobaoTbkRtaConsumerMatchResponse:
		respStruct := body.(*response.TaobaoTbkRtaConsumerMatchResponse)
		if respStruct.RequestId == "" {
			err := errors.New("RequestId is empty")
			return err
		}
		if respStruct.Data.ResultList == nil {
			err := errors.New(fmt.Sprintf("RequestId: %v ResultList is nil", respStruct.RequestId))
			return err
		}

		//{"data":{"result_list":[]},"request_id":"16mcb0anm8xhi"}
		//存在 result_list 为空的情况
		//if len(*respStruct.Data.ResultList) == 0 {
		//	err := errors.New(fmt.Sprintf("RequestId: %v ResultList 为空", respStruct.RequestId))
		//	return err
		//}
		return nil
	default:
		return nil
	}
}

func NewPlatformRequest(adZoneId int64, deviceType, deviceValue, strategyIdList, specialId string) request.TaobaoTbkRtaConsumerMatchRequest {
	var specialIdPtr *string
	if len(specialId) > 0 {
		specialIdPtr = &specialId
	}
	req := request.TaobaoTbkRtaConsumerMatchRequest{
		AdZoneId:       &adZoneId,
		SpecialId:      specialIdPtr,
		DeviceValue:    &deviceValue,
		DeviceType:     &deviceType,
		StrategyIdList: &strategyIdList,
	}
	return req
}
