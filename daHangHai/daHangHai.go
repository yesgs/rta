package daHangHai

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"rta"
	"rta/daHangHai/request"
	"rta/daHangHai/response"

	"net/http"
	"net/url"
	"rta/util"
	"sort"
	"strings"
	"time"
	"unsafe"
)

type Client struct {
	rta.DefaultRtaClient

	AppKey     string
	AppSecret  string
	SignMethod string
	Version    string
	Format     string
	Simplify   bool
}

func (c *Client) Ask(payload interface{}) (data []byte, err error) {
	var (
		sign        string
		publicParam = make(map[string]interface{})
		payloadMap  = payload.(map[string]interface{})
	)

	publicParam["method"] = "taobao.usergrowth.dhh.delivery.batchask"
	publicParam["app_key"] = c.AppKey
	publicParam["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	publicParam["v"] = c.Version
	publicParam["sign_method"] = c.SignMethod
	publicParam["format"] = c.Format
	publicParam["simplify"] = c.Simplify
	publicParam["partner_id"] = "new_go_sdk"

	sign = c.getSign(publicParam, payloadMap)

	serverUrl, _ := url.Parse(c.Opts.BaseUrl)
	urlValues := url.Values{}
	urlValues.Add("sign", sign)
	for k, v := range publicParam {
		urlValues.Add(k, fmt.Sprint(v))
	}
	serverUrl.RawQuery = urlValues.Encode()
	urlPath := serverUrl.String()

	bodyParam := url.Values{}
	for k, v := range payloadMap {
		bodyParam.Add(k, fmt.Sprint(v))
	}
	var httpClient = c.DefaultRtaClient.GetHttpClient()

	req, _ := http.NewRequest(c.Opts.HttpMethod, urlPath, strings.NewReader(bodyParam.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) ConvertRequest(body interface{}) (interface{}, error) {
	batchAskRequest := body.(request.TaobaoUsergrowthDhhDeliveryBatchaskRequest)
	return batchAskRequest.ToMap(), nil
}

func (c *Client) getSign(publicParam map[string]interface{}, data map[string]interface{}) string {
	var (
		initialSize = 0
		allParamMap = make(map[string]interface{})
		keyList     []string
		sb          strings.Builder
		sbStr       string
	)

	for k, v := range data {
		allParamMap[k] = v
		keyList = append(keyList, k)
		switch v.(type) {
		case string:
			initialSize += len(v.(string))
		default:
			initialSize += int(unsafe.Sizeof(v))
		}
	}

	for k, v := range publicParam {
		allParamMap[k] = v
		keyList = append(keyList, k)
		switch v.(type) {
		case string:
			initialSize += len(v.(string))
		default:
			initialSize += int(unsafe.Sizeof(v))
		}
	}

	sort.Strings(keyList)

	sb.Grow(initialSize + 100)

	for _, key := range keyList {
		value := allParamMap[key]
		sb.WriteString(fmt.Sprintf("%v%v", key, value))
	}

	h := hmac.New(sha256.New, util.StringToBytes(&c.AppSecret))
	sbStr = sb.String()
	h.Write(util.StringToBytes(&sbStr))
	sign := hex.EncodeToString(h.Sum(nil))
	return strings.ToUpper(sign)
}

func (c *Client) ConvertResponse(body []byte, output interface{}) (err error) {
	_ = c.extractContent(body, output)
	return c.ResponseHasBusinessError(output)
}

func (c *Client) extractContent(content []byte, v interface{}) (err error) {
	err = json.Unmarshal(content, v)
	if err == nil {
		return nil
	} else {
		txtStr := string(content)
		if len(txtStr) < 20 {
			return errors.New("response content too short")
		} else {
			if strings.Contains(txtStr[0:20], "error_response") {
				reqErr := &TopApiRequestError{}
				txtStr = txtStr[18 : len(txtStr)-1]
				err2 := json.Unmarshal([]byte(txtStr), reqErr)
				if err2 != nil {
					return err2
				}
				return err
			}
			return errors.New("unknown error")
		}
	}
}

func (c *Client) ResponseHasBusinessError(body interface{}) error {
	switch body.(type) {
	case *response.TaobaoUsergrowthDhhDeliveryBatchaskResponse:
		respStruct := body.(*response.TaobaoUsergrowthDhhDeliveryBatchaskResponse)
		if respStruct.RequestId == "" || (respStruct.Result.Errcode != nil && *(respStruct.Result.Errcode) != 0) {
			respStruct.Body = fmt.Sprint(body)
			err := errors.New(fmt.Sprintf("unknown error: %v", respStruct.Body))
			return err
		}
		return nil
	default:
		return nil
	}
}

func NewClient(opt *rta.Options, appKey, appSecret string) rta.ClientInterface {
	opt.Init()
	return &Client{
		DefaultRtaClient: rta.DefaultRtaClient{
			Opts: opt,
		},
		AppKey:     appKey,
		AppSecret:  appSecret,
		SignMethod: "hmac-sha256",
		Version:    "2.0",
		Format:     "json",
		Simplify:   true,
	}
}

func NewPlatformRequest(adSpaceId, channel string, imei []string, oaid []string, idfa []string) request.TaobaoUsergrowthDhhDeliveryBatchaskRequest {
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

	req.SetAdvertisingSpaceId(adSpaceId)
	req.SetChannel(channel)

	return req
}
