package tbk

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
	return c.Invoke("taobao.tbk.dg.vegas.send.status", payload)
}

func (c *Client) ConvertRequest(body interface{}) (interface{}, error) {
	askRequest := body.(request.TaobaoTbkDgVegasSendStatusRequest)
	return askRequest.ToMap(), nil
}

func (c *Client) ConvertResponse(body []byte, output interface{}) (err error) {
	_ = c.ExtractContent(body, output)
	return c.ResponseHasBusinessError(output)
}

func (c *Client) ResponseHasBusinessError(body interface{}) error {
	switch body.(type) {
	case *response.TaobaoTbkDgVegasSendStatusResponse:
		respStruct := body.(*response.TaobaoTbkDgVegasSendStatusResponse)
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

func NewPlatformRequest(pid, deviceType, deviceValue, bizCode string) request.TaobaoTbkDgVegasSendStatusRequest {
	req := request.TaobaoTbkDgVegasSendStatusRequest{
		DeviceValue: &deviceValue,
		DeviceType:  &deviceType,
		ThorBizCode: &bizCode,
		Pid:         &pid,
	}
	return req
}
