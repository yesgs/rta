package dhh

import (
	"errors"
	"fmt"
	"github.com/yesgs/rta/taobao"
	"github.com/yesgs/rta/taobao/request"
	"github.com/yesgs/rta/taobao/response"
	"strings"
)

type Client struct {
	*taobao.Client
}

func NewClient(taobaoClient *taobao.Client) *Client {
	return &Client{taobaoClient}
}

func (c *Client) Ask(payload interface{}) (data []byte, err error) {
	return c.Invoke("taobao.usergrowth.dhh.delivery.batchask", payload)
}

func (c *Client) ConvertRequest(body interface{}) (interface{}, error) {
	batchAskRequest := body.(request.TaobaoUsergrowthDhhDeliveryBatchaskRequest)
	return batchAskRequest.ToMap(), nil
}

func (c *Client) ConvertResponse(body []byte, output interface{}) (err error) {
	_ = c.ExtractContent(body, output)
	return c.ResponseHasBusinessError(output)
}

func (c *Client) ResponseHasBusinessError(body interface{}) error {
	switch body.(type) {
	case *response.TaobaoUsergrowthDhhDeliveryBatchaskResponse:
		respStruct := body.(*response.TaobaoUsergrowthDhhDeliveryBatchaskResponse)
		if respStruct.RequestId == "" {
			err := errors.New("RequestId is empty")
			return err
		}
		if respStruct.Result.Errcode == nil {
			err := errors.New(fmt.Sprintf("RequestId: %v Errcode is nil", respStruct.RequestId))
			return err
		}

		if *(respStruct.Result.Errcode) != 0 {
			err := errors.New(fmt.Sprintf("RequestId: %v Errcode: %v != 0", respStruct.RequestId, *(respStruct.Result.Errcode)))
			return err
		}
		return nil
	default:
		return nil
	}
}

func NewPlatformRequest(adSpaceId, channel string, imei []string, oaid []string, idfa []string, caid []string) request.TaobaoUsergrowthDhhDeliveryBatchaskRequest {
	req := request.TaobaoUsergrowthDhhDeliveryBatchaskRequest{}

	if len(imei) > 0 {
		req.SetImeiMd5(strings.Join(imei, ","))
	}
	if len(oaid) > 0 {
		req.SetOaidMd5(strings.Join(oaid, ","))
	}
	if len(idfa) > 0 {
		req.SetIdfaMd5(strings.Join(idfa, ","))
	}
	if len(caid) > 0 {
		req.SetCaidMd5(strings.Join(caid, ","))
	}

	req.SetAdvertisingSpaceId(adSpaceId)
	req.SetChannel(channel)

	return req
}
