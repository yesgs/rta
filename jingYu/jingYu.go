package jingYu

import (
	"errors"
	"fmt"
	"rta"
	"rta/util"
)

type Request struct {
	Channel  string        `json:"channel"`
	Platform string        `json:"platform"`
	Device   RequestDevice `json:"device"`
}

type RequestDevice struct {
	ImeiMd5 *string `json:"imeiMd5,omitempty"`
	OaidMd5 *string `json:"oaidMd5,omitempty"`
	IdfaMd5 *string `json:"idfaMd5,omitempty"`
}

type Response struct {
	Code     string `json:"code"`
	Error    string `json:"error"`
	Platform string `json:"platform"`
	Status   bool   `json:"status"`
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
		if len(body.(*Response).Error) > 0 {
			return errors.New(fmt.Sprintf("code: %v err: %v", body.(*Response).Code, body.(*Response).Error))
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

func NewPlatformRequest(channel string, platform string, deviceType, deviceId string) Request {
	req := Request{
		Channel:  channel,
		Platform: platform,
	}
	device := RequestDevice{}

	switch deviceType {
	case util.DeviceIMEI:
		if len(deviceId) > 0 {
			device.ImeiMd5 = &deviceId
		}
	case util.DeviceOAID:
		if len(deviceId) > 0 {
			device.OaidMd5 = &deviceId
		}

	case util.DeviceIDFA:
		if len(deviceId) > 0 {
			device.IdfaMd5 = &deviceId
		}
	}

	req.Device = device
	return req
}
