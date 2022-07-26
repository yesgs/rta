package rta

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type DefaultRtaClient struct {
	Opts *Options
}

func (t *DefaultRtaClient) GetHttpClient() *http.Client {
	t.Opts.Mutex.Lock()
	defer t.Opts.Mutex.Unlock()
	return t.Opts.HttpClient
}

func (t *DefaultRtaClient) ConvertRequest(body interface{}) (interface{}, error) {
	return json.Marshal(body)
}

func (t *DefaultRtaClient) ResponseHasBusinessError(body interface{}) error {
	return nil
}

func (t *DefaultRtaClient) ConvertResponse(body []byte, output interface{}) error {
	if output == nil {
		return errors.New("output is nil")
	}

	err := json.Unmarshal(body, output)
	if err != nil {
		return err
	}

	return nil
}

func (t *DefaultRtaClient) Ask(payload interface{}) (data []byte, err error) {
	var (
		req        *http.Request
		url        = t.Opts.BaseUrl
		httpClient = t.GetHttpClient()
	)

	switch t.Opts.HttpMethod {
	case http.MethodGet:
		url += "?" + string(payload.([]byte))
		req, err = http.NewRequest(http.MethodGet, url, nil)
	case http.MethodPost:
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewReader(payload.([]byte)))
	}

	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	return data, err
}

func NewPlatformRequest(opt *Options) ClientInterface {
	opt.Init()
	return &DefaultRtaClient{
		Opts: opt,
	}
}
