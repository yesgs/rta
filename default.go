package rta

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type DefaultRtaClient struct {
	Opts *Options
}

func (c *DefaultRtaClient) GetHttpClient() *http.Client {
	c.Opts.Mutex.Lock()
	defer c.Opts.Mutex.Unlock()
	return c.Opts.HttpClient
}

func (c *DefaultRtaClient) ConvertRequest(body interface{}) (interface{}, error) {
	return json.Marshal(body)
}

func (c *DefaultRtaClient) ResponseHasBusinessError(body interface{}) error {
	return nil
}

func (c *DefaultRtaClient) ConvertResponse(body []byte, output interface{}) error {
	if output == nil {
		return errors.New("output is nil")
	}

	err := json.Unmarshal(body, output)
	if err != nil {
		return err
	}

	return nil
}

func (c *DefaultRtaClient) Ask(payload interface{}) (data []byte, err error) {
	var (
		req        *http.Request
		url        = c.Opts.BaseUrl
		httpClient = c.GetHttpClient()
	)

	switch c.Opts.HttpMethod {
	case http.MethodGet:
		switch payload.(type) {
		case []byte:
			url += "?" + string(payload.([]byte))
		case string:
			url += "?" + payload.(string)
		default:
			url += "?" + ""
		}
		req, err = http.NewRequest(http.MethodGet, url, nil)
	case http.MethodPost:
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewReader(payload.([]byte)))
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	return data, err
}

func NewPlatformRequest(opt *Options) ClientInterface {
	opt.Init()
	return &DefaultRtaClient{
		Opts: opt,
	}
}
