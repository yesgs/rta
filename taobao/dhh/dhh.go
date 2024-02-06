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
	switch body.(type) {
	case request.TaobaoUsergrowthDhhDeliveryBatchaskRequest:
		batchAskRequest := body.(request.TaobaoUsergrowthDhhDeliveryBatchaskRequest)
		return batchAskRequest.ToMap(), nil
	case request.TaobaoUsergrowthDhhDeliveryAskRequest:
		singleAskRequest := body.(request.TaobaoUsergrowthDhhDeliveryAskRequest)
		return singleAskRequest.ToMap(), nil
	}
	return nil, errors.New("unknown")
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

func NewPlatformRequestSingleQuery(adSpaceId, channel string, imeiMd5 string, oaidMd5 string, idfaMd5 string, caidMd5 string) request.TaobaoUsergrowthDhhDeliveryAskRequest {
	req := request.TaobaoUsergrowthDhhDeliveryAskRequest{
		Profile:            nil,
		OaidMd5:            nil,
		IdfaMd5:            nil,
		ImeiMd5:            nil,
		AdvertisingSpaceId: &adSpaceId,
		Channel:            &channel,
		CaidMd5:            nil,
	}
	if len(imeiMd5) > 0 {
		req.ImeiMd5 = &imeiMd5
	}
	if len(oaidMd5) > 0 {
		req.OaidMd5 = &oaidMd5
	}
	if len(idfaMd5) > 0 {
		req.IdfaMd5 = &idfaMd5
	}
	if len(caidMd5) > 0 {
		req.CaidMd5 = &caidMd5
	}
	return req
}
