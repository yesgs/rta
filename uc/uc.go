package uc

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yesgs/rta"
)

type Request struct {
	Channel  string `json:"channel"`
	Platform string `json:"platform"` //ANDROID IOS UNKNOWN
	DidType  string `json:"didType"`  //IMEI IMEI_MD5 OAID OAID_MD5 ANDROID_ID ANDROID_ID_MD5 IDFA IDFA_MD5
	Did      string `json:"did"`
}

type Response struct {
	Message    string  `json:"message"`
	AcIds      []int64 `json:"acIds"` //投放的⼴告账号ID列表
	AdIds      []int64 `json:"adIds"` //投放的⼴告计划ID列表
	StatusCode int     `json:"status_code"`
	//
	//状态码:
	//0 : 需要投放
	//1 : 不需要投放
	//>1 : 其他错误ID
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
		if body.(*Response).StatusCode > 1 || body.(*Response).StatusCode < 0 {
			return errors.New(fmt.Sprintf("code: %v err: %v", body.(*Response).StatusCode, body.(*Response).Message))
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

func NewPlatformRequest(channel string, platform string, didType string, did string) Request {
	req := Request{
		Channel:  channel,
		Platform: platform,
		DidType:  didType,
		Did:      did,
	}
	return req
}
