package dhhSingleAsk

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
	return c.Invoke("taobao.usergrowth.dhh.delivery.ask", payload)
}

func (c *Client) ConvertRequest(body interface{}) (interface{}, error) {
	singleAskRequest := body.(request.TaobaoUsergrowthDhhDeliveryAskRequest)
	return singleAskRequest.ToMap(), nil
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
	case *response.TaobaoUsergrowthDhhDeliveryAskResponse:
		respStruct := body.(*response.TaobaoUsergrowthDhhDeliveryAskResponse)
		if respStruct.RequestId == "" {
			err := errors.New("RequestId is empty")
			return err
		}

		if respStruct.Errcode != 0 {
			err := errors.New(fmt.Sprintf("RequestId: %v Errcode: %v != 0", respStruct.RequestId, respStruct.Errcode))
			return err
		}
		return nil
	default:
		return nil
	}
}

func NewPlatformRequest(adSpaceId, channel string, imei string, oaid string, idfa string, caid string) request.TaobaoUsergrowthDhhDeliveryAskRequest {
	req := request.TaobaoUsergrowthDhhDeliveryAskRequest{}

	req.SetAdvertisingSpaceId(adSpaceId)
	req.SetChannel(channel)

	if len(imei) > 0 {
		req.SetImeiMd5(imei)
		return req
	}
	if len(oaid) > 0 {
		req.SetOaidMd5(oaid)
		return req
	}
	if len(idfa) > 0 {
		req.SetIdfaMd5(idfa)
		return req
	}
	if len(caid) > 0 {
		req.SetCaidMd5(caid)
		return req
	}

	return req
}
