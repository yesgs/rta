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

		if respStruct.ResultMsg != "成功" {
			err := errors.New(fmt.Sprintf("RequestId: %v ResultMsg: %v != 0", respStruct.RequestId, respStruct.ResultMsg))
			return err
		}
		return nil
	default:
		return nil
	}
}

func NewPlatformRequest(pid, deviceType, deviceValue string) request.TaobaoTbkDgVegasSendStatusRequest {
	var bizCode = "1"
	req := request.TaobaoTbkDgVegasSendStatusRequest{
		DeviceValue: &deviceValue,
		DeviceType:  &deviceType,
		ThorBizCode: &bizCode,
		Pid:         &pid,
	}
	return req
}
