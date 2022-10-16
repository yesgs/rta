package iWingSky

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yesgs/rta"
)

type Request struct {
	Token   string `json:"token"`
	Id      string `json:"id"`
	IdKind  int    `json:"id_kind"` //1:设备唯一标识 2:设备唯一标识Md5
	Channel string `json:"channel"`
}

type Response struct {
	Code    int    `json:"code"` //0:不存在 1:已存在 500:异常
	Message string `json:"message"`
}

type Client struct {
	rta.DefaultRtaClient
	Token   string
	Channel string
}

func (c *Client) ConvertRequest(body interface{}) (interface{}, error) {
	var reqBody = body.(Request)
	reqBody.Channel = c.Channel
	reqBody.Token = c.Token
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
		if body.(*Response).Code > 0 {
			return errors.New(fmt.Sprintf("code: %v err: %v", body.(*Response).Code, body.(*Response).Message))
		}
		return nil
	default:
		return nil
	}
}

func NewClient(opt *rta.Options, channel, token string) rta.ClientInterface {
	opt.Init()
	return &Client{
		DefaultRtaClient: rta.DefaultRtaClient{
			Opts: opt,
		},
		Channel: channel,
		Token:   token,
	}
}

func NewPlatformRequest(id, idKind string) Request {
	req := Request{
		Id:     id,
		IdKind: idKind,
	}
	return req
}
