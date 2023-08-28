package vip

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yesgs/rta"
)

type Client struct {
	rta.DefaultRtaClient
}
type Request struct {
	AdReqs  []AdReqItem `json:"adReqs"`
	Id      string      `json:"id"`
	DidType int         `json:"didType"`
	Did     string      `json:"did"`
}

type AdReqItem struct {
	Token string `json:"token"`
}

type Response struct {
	BidResp BidResp `json:"bidResp"`
	Errno   int     `json:"errno"`
	Msg     string  `json:"msg"`
}

type BidResp struct {
	AdResp  []AdRespItem `json:"adResps"`
	BidId   string       `json:"bidId"`
	Id      string       `json:"id"`
	ResType int          `json:"resType"`
}

type AdRespItem struct {
	Ac    string `json:"ac"`
	Token string `json:"token"`
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
		if body.(*Response).Errno != 0 {
			return errors.New(fmt.Sprintf("Msg: %v", body.(*Response).Msg))
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

func NewPlatformRequest(uuid string, tokenList []string, didType int, did string) Request {
	var adReqs = make([]AdReqItem, 0, len(tokenList))
	for i := 0; i < len(tokenList); i++ {
		adReqs = append(adReqs, AdReqItem{
			Token: tokenList[i],
		})
	}
	req := Request{
		AdReqs:  adReqs,
		Id:      uuid,
		DidType: didType,
		Did:     did,
	}
	return req
}
