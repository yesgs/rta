package rta

import (
	"net/http"
	"sync"
	"time"
)

type ClientInterface interface {
	ConvertRequest(body interface{}) (interface{}, error)
	Ask(payload interface{}) ([]byte, error)
	ConvertResponse(body []byte, output interface{}) error
	ResponseHasBusinessError(body interface{}) error
}

type Options struct {
	BaseUrl    string
	HttpMethod string
	RtaClient  ClientInterface
	HttpClient *http.Client
	Mutex      *sync.Mutex
}

func (o *Options) Init() {
	if len(o.HttpMethod) == 0 {
		o.HttpMethod = http.MethodPost
	}
	if o.HttpClient == nil {
		o.HttpClient = NewHttpClient()
	}
	if o.Mutex == nil {
		o.Mutex = &sync.Mutex{}
	}
}

func NewHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:      false,
			DisableCompression:     false,
			MaxIdleConns:           0,
			MaxIdleConnsPerHost:    0,
			MaxConnsPerHost:        0,
			IdleConnTimeout:        0,
			ResponseHeaderTimeout:  0,
			MaxResponseHeaderBytes: 0,
			WriteBufferSize:        0,
			ReadBufferSize:         0,
			ForceAttemptHTTP2:      false,
		},
		Timeout: time.Second * 1,
	}
}

func MakeRequest(cli ClientInterface, platformRequest, platformResponse interface{}) error {
	reqBody, err := cli.ConvertRequest(platformRequest)
	if err != nil {
		return err
	}
	respBody, err := cli.Ask(reqBody)
	if err != nil {
		return err
	}

	err = cli.ConvertResponse(respBody, platformResponse)
	if err != nil {
		return err
	}
	return nil
}
