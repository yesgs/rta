package talkingData

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yesgs/rta"
)

type Request struct {
	MediaId      string   `json:"mediaId"`
	CampaignId   []string `json:"campaignId"`
	DeviceIdType int      `json:"deviceIdType"`
	DeviceId     string   `json:"deviceId"`
}

//对应DeviceIdType的设备ID，设备ID规则如下：
//IDFA：32 位的数字 + 大写字母串，保留中划线"-"；
//IMEI：14 位或 15 位的纯数字串，或者 14 位或 15 位数字 + 小写字母串；
//OAID：保留原始值（不要转换大小写），由数字字母和连接线构成，具体格式取决于各收集厂商和系统版本，长度有别；
//IDFA_MD5：即加密后的 IDFA，加密前需要格式转化成 32 位的数字 + 大写字母，加密后为不计大小写的 32 位数字字母串；
//IMEI_MD5：即加密后的 IMEI，加密前需要格式转化成 14 位或 15 位数字 + 小写字母串，加密后为不计大小写的 32 位数字字母串；
//OAID_MD5：即加密后的 OAID，加密前请保留OAID原值（不要转换大小写或去连接符），加密后为不计大小写的 32 位字符串；

type CampaignMatchItem struct {
	CampaignId string `json:"campaignId"`
	Matched    string `json:"matched"`
}

type Response struct {
	Code   int                 `json:"code"`
	Result []CampaignMatchItem `json:"result"`
}

type Client struct {
	rta.DefaultRtaClient
}

func (c *Client) ConvertRequest(body interface{}) (interface{}, error) {
	var reqBody = body.(Request)
	return json.Marshal(reqBody)
}

func (c *Client) ConvertResponse(body []byte, output interface{}) (err error) {
	err = c.DefaultRtaClient.ConvertResponse(body, output)
	if err != nil {
		return err
	}
	return c.ResponseHasBusinessError(output)
}

func (c *Client) ResponseHasBusinessError(body interface{}) error {
	switch body.(type) {
	case *Response:
		if body.(*Response).Code != 2000 {
			return errors.New(fmt.Sprintf("code: %v", body.(*Response).Code))
		}
		return nil
	default:
		return nil
	}
}

func NewClient(opt *rta.Options) rta.ClientInterface {
	opt.Init()
	return &Client{
		DefaultRtaClient: rta.DefaultRtaClient{
			Opts: opt,
		},
	}
}

func NewPlatformRequest(mediaId string, campaignId []string, deviceIdType int, deviceId string) Request {
	req := Request{
		MediaId:      mediaId,
		CampaignId:   campaignId,
		DeviceIdType: deviceIdType,
		DeviceId:     deviceId,
	}
	return req
}
