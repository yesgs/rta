package taobao

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yesgs/rta"
	"github.com/yesgs/rta/util"
	"io/ioutil"
	"net/http"
	"net/url"
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

func (c *Client) GetSign(publicParam map[string]interface{}, data map[string]interface{}) string {
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

func (c *Client) GetPublicParam() map[string]interface{} {
	var publicParam = make(map[string]interface{})
	publicParam["app_key"] = c.AppKey
	publicParam["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	publicParam["v"] = c.Version
	publicParam["sign_method"] = c.SignMethod
	publicParam["format"] = c.Format
	publicParam["simplify"] = c.Simplify
	publicParam["partner_id"] = "new_go_sdk"
	return publicParam
}

func (c *Client) ExtractContent(content []byte, v interface{}) (err error) {
	txtStr := string(content)
	if strings.Contains(txtStr, "error_response") {
		reqErr := &TopApiRequestError{}
		txtStr = txtStr[18 : len(txtStr)-1]
		err2 := json.Unmarshal([]byte(txtStr), reqErr)
		if err2 != nil {
			return err2
		}
		return fmt.Errorf("code: %v msg: %v sub_code: %v sub_msg: %v request_id: %v", reqErr.TopCode, reqErr.Msg, reqErr.SubCode, reqErr.SubMsg, reqErr.RequestId)
	}

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

func (c *Client) Invoke(method string, payload interface{}) (data []byte, err error) {
	var (
		serverUrl   *url.URL
		sign        string
		urlPath     string
		urlValues   = url.Values{}
		bodyValues  = url.Values{}
		publicParam = make(map[string]interface{})
		payloadMap  = payload.(map[string]interface{})
	)

	publicParam = c.GetPublicParam()

	publicParam["method"] = method
	sign = c.GetSign(publicParam, payloadMap)

	serverUrl, err = url.Parse(c.Opts.BaseUrl)
	if err != nil {
		return nil, err
	}
	urlValues.Add("sign", sign)
	for k, v := range publicParam {
		urlValues.Add(k, fmt.Sprint(v))
	}
	serverUrl.RawQuery = urlValues.Encode()
	urlPath = serverUrl.String()

	for k, v := range payloadMap {
		bodyValues.Add(k, fmt.Sprint(v))
	}
	var httpClient = c.DefaultRtaClient.GetHttpClient()

	req, err := http.NewRequest(c.Opts.HttpMethod, urlPath, strings.NewReader(bodyValues.Encode()))
	if err != nil {
		return nil, err
	}

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

func NewClient(opt *rta.Options, appKey, appSecret string) *Client {
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
