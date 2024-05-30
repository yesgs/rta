package autohome

import (
	"errors"
	"fmt"
	"github.com/yesgs/rta"
	"net/url"
)

type Request struct {
	Id      string `json:"id"`
	IdType  string `json:"idType"`
	Channel string `json:"channel"`
}

//IdType 类型枚举：
//OAID
//IDFA
//IMEI
//ANDROID_ID
//IMEI_MD5（IMEI md5加密）
//IDFA_MD5（IDFA md5加密）
//OAID_MD5（OAID md5加密）
//ANDROID_ID_MD5（ANDROID_ID md5加密）

type Response struct {
	ReturnCode int    `json:"returncode"` //业务级别的返回码。0 表示成功，非 0 表示失败
	Message    string `json:"message"`    //返回码详细描述
	Result     Result `json:"result"`
}

type Result struct {
	Allowed       int      `json:"allowed"` //0:请求设备不存在, 1:准入投放
	MaterialCodes []string `json:"materialCodes"`
}

type Client struct {
	rta.DefaultRtaClient
}

func (c *Client) ConvertRequest(body interface{}) (interface{}, error) {
	return nil, nil
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
		if body.(*Response).ReturnCode != 0 {
			return errors.New(fmt.Sprintf("code: %v err: %v", body.(*Response).ReturnCode, body.(*Response).Message))
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

func NewPlatformRequest(channel string, idType string, id string) string {
	var uv = url.Values{}
	uv.Add("id", id)
	uv.Add("idType", idType)
	uv.Add("channel", channel)
	return uv.Encode()
}
