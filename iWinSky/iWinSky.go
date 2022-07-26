package iWinSky

import (
	"errors"
	"fmt"
	"github.com/yesgs/rta"
)

type Request struct {
	Token   string `json:"token"`
	Id      string `json:"id"`
	IdKind  string `json:"id_kind"` //1:设备唯一标识 2:设备唯一标识Md5
	Channel string `json:"channel"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Client struct {
	rta.DefaultRtaClient
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

func NewClient(opt *rta.Options) rta.ClientInterface {
	opt.Init()
	return &Client{
		DefaultRtaClient: rta.DefaultRtaClient{
			Opts: opt,
		},
	}
}

func NewPlatformRequest(token string, id string, idKind, channel string) Request {
	req := Request{
		Token:   token,
		Id:      id,
		IdKind:  idKind,
		Channel: channel,
	}
	return req
}
