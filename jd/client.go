package jd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yesgs/rta"
	"github.com/yesgs/rta/jd/domain/request"
	"github.com/yesgs/rta/jd/domain/response"
	"github.com/yesgs/rta/jd/internal"
	"github.com/yesgs/rta/jd/util"
	"net/http"
	"time"
)

type Client struct {
	rta.DefaultRtaClient

	serverUrl string

	accessToken string
	appKey      string
	appSecret   string

	version    string
	format     string
	signMethod string
}

type Request struct {
	DeviceType int    `json:"deviceType"`
	DeviceId   string `json:"deviceId"`
}

type BUResponse struct {
	Code           string `json:"code"`
	ValidateResult string `json:"validateResult"`
}

type Response struct {
	ErrorResponse                           *ErrorResponse `json:"error_response"`
	JdUnionOpenUserRegisterValidateResponce *BUResponse    `json:"jd_union_open_user_register_validate_responce"`
}

type ErrorResponse struct {
	Code   string `json:"code"`
	ZhDesc string `json:"zh_desc"`
	EnDesc string `json:"en_desc"`
}

func (c *Client) GetHttpClient() *http.Client {
	return c.Opts.HttpClient
}

func NewClient(opt *rta.Options, appKey, appSecret string) rta.ClientInterface {
	opt.Init()
	return &Client{
		DefaultRtaClient: rta.DefaultRtaClient{
			Opts: opt,
		},
		serverUrl:   "https://api.jd.com/routerjson",
		accessToken: "",
		appKey:      appKey,
		appSecret:   appSecret,
		version:     "1.0",
		format:      "json",
		signMethod:  "md5",
	}
}

func (c *Client) Execute(req interface{}) ([]byte, error) {
	httpClient := c.GetHttpClient()
	// 请求JD Api服务器
	respBytes, err := util.HttpGet(httpClient, c.serverUrl, req)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}

func (c *Client) Ask(payload interface{}) (data []byte, err error) {
	return c.Execute(payload)
}

func (c *Client) ConvertRequest(body interface{}) (interface{}, error) {
	var plainReq = body.(Request)
	var req = request.UnionOpenUserRegisterValidateRawRequest{
		UserStateReq: &request.UserStateRequest{
			UserId:     plainReq.DeviceId,
			UserIdType: plainReq.DeviceType,
		},
	}

	// get business params
	jsonParams, err := req.JsonParams()
	if err != nil {
		return nil, err
	}

	// sign
	timestamp := time.Now().Format("2006-01-02 15:04:05") + ".000+0800"
	signParams := map[string]string{
		"method":            req.RequestMethod(),
		"app_key":           c.appKey,
		"timestamp":         timestamp,
		"format":            c.format,
		"v":                 c.version,
		"sign_method":       c.signMethod,
		"360buy_param_json": string(jsonParams),
	}
	signValue := util.Sign(signParams, c.appSecret)

	params := internal.SystemConfig{
		Method:      req.RequestMethod(),
		AppKey:      c.appKey,
		AccessToken: c.accessToken,
		Timestamp:   timestamp,
		Format:      c.format,
		Version:     c.version,
		SignMethod:  c.signMethod,
		Sign:        signValue,
		ParamJson:   string(jsonParams),
	}

	return params, nil
}

func (c *Client) ConvertResponse(body []byte, output interface{}) (err error) {
	err = c.ExtractContent(body, output)
	if err != nil {
		return err
	}
	return c.ResponseHasBusinessError(output)
}

func (c *Client) ExtractContent(content []byte, v interface{}) (err error) {
	var r1 = Response{}
	err = json.Unmarshal(content, &r1)
	if err != nil {
		return err
	}

	if r1.ErrorResponse != nil {
		return fmt.Errorf("code: %v zh_desc: %v", r1.ErrorResponse.Code, r1.ErrorResponse.ZhDesc)
	}

	if r1.JdUnionOpenUserRegisterValidateResponce == nil {
		return fmt.Errorf("JdUnionOpenUserRegisterValidateResponce is nil")
	}

	if r1.JdUnionOpenUserRegisterValidateResponce.Code != "0" {
		return fmt.Errorf("jd_union_open_user_register_validate_responce.code is %v", r1.JdUnionOpenUserRegisterValidateResponce.Code)
	}

	rst := r1.JdUnionOpenUserRegisterValidateResponce.ValidateResult

	return json.Unmarshal([]byte(rst), v)
}

func (c *Client) ResponseHasBusinessError(body interface{}) error {
	switch body.(type) {
	case *response.UnionOpenUserRegisterValidateRowResponse:
		v := body.(*response.UnionOpenUserRegisterValidateRowResponse)
		if v.Code != 200 {
			return errors.New(fmt.Sprintf("code: %v err: %v", v.Code, v.Message))
		}
		return nil
	default:
		return fmt.Errorf("unknown error %v", body)
	}
}

func NewPlatformRequest(deviceType int, deviceId string) Request {
	req := Request{
		DeviceType: deviceType,
		DeviceId:   deviceId,
	}
	return req
}
