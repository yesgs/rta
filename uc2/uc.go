package uc2

import (
	"errors"
	"fmt"
	"github.com/yesgs/rta"
	"net/url"
	"strconv"
	"time"
)

type Request struct {
	OaidMd5   string `json:"oaid_md5"`
	Caid      string `json:"caid"`
	Timestamp int64  `json:"timestamp"`
	Sign      string `json:"sign"`
}

type Response struct {
	Code int `json:"code"`
	//
	//状态码:
	//0 : 需要投放
	//1 : 不需要投放
	//>1 : 其他错误ID

	Msg       string   `json:"msg"`
	RequestId string   `json:"requestId"`
	Tasks     []string `json:"tasks"` //投放的广告任务D列表
}

type Client struct {
	rta.DefaultRtaClient
	Channel string
	Secret  string
}

func NewClient(opt *rta.Options, channel string, secret string) rta.ClientInterface {
	opt.Init()
	return &Client{
		DefaultRtaClient: rta.DefaultRtaClient{
			Opts: opt,
		},
		Channel: channel,
		Secret:  secret,
	}
}

func (c *Client) Execute(req interface{}) ([]byte, error) {
	httpClient := c.GetHttpClient()
	respBytes, err := HttpGet(httpClient, c.Opts.BaseUrl+"?"+req.(url.Values).Encode())
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}

func (c *Client) Ask(payload interface{}) (data []byte, err error) {
	return c.Execute(payload)
}

func (c *Client) ConvertRequest(body interface{}) (interface{}, error) {
	var reqBody = body.(Request)
	var paramMap = make(map[string]string)
	var uv = url.Values{}

	reqBody.Timestamp = time.Now().Unix()
	paramMap["timestamp"] = strconv.FormatInt(reqBody.Timestamp, 10)
	uv.Set("timestamp", strconv.FormatInt(reqBody.Timestamp, 10))

	if len(reqBody.OaidMd5) > 0 {
		paramMap["oaid_md5"] = reqBody.OaidMd5
		uv.Set("oaid_md5", reqBody.OaidMd5)
	}
	if len(reqBody.Caid) > 0 {
		paramMap["caid"] = reqBody.Caid
		uv.Set("caid", reqBody.Caid)
	}

	sign := Sign(paramMap, c.Secret)
	uv.Set("sign", sign)

	reqBody.Sign = sign
	return uv, nil
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
		if body.(*Response).Code > 1 || body.(*Response).Code < 0 {
			return errors.New(fmt.Sprintf("code: %v err: %v", body.(*Response).Code, body.(*Response).Msg))
		}
		return nil
	default:
		return nil
	}
}

func NewPlatformRequest(oaidMd5 string, caid string) Request {
	req := Request{
		OaidMd5: oaidMd5,
		Caid:    caid,
	}
	return req
}
