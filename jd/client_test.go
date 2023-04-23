package jd

import (
	"crypto/tls"
	"github.com/yesgs/rta"
	"github.com/yesgs/rta/jd/domain/response"
	"log"
	"net/http"
	"testing"
	"time"
)

func Test_f1(t *testing.T) {
	rtaOption := rta.Options{
		BaseUrl:    "https://api.jd.com/routerjson",
		HttpMethod: http.MethodPost,
		HttpClient: &http.Client{
			Timeout: 1 * time.Second,
			Transport: &http.Transport{
				DisableKeepAlives:     true,
				DisableCompression:    false,
				MaxIdleConns:          10240,
				MaxIdleConnsPerHost:   10240,
				MaxConnsPerHost:       81920,
				IdleConnTimeout:       time.Second * 60,
				ResponseHeaderTimeout: time.Second * 60,
				//WriteBufferSize:       0,
				//ReadBufferSize:        0,
				ForceAttemptHTTP2: false,
				TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			},
		},
	}

	var appKey = "aa"
	var appSecret = "bb"

	cli := NewClient(&rtaOption, appKey, appSecret)
	py := Request{
		DeviceType: 131072,
		DeviceId:   "58D2C273ECA672EB3A65134BA3285672",
	}

	reqBody, err := cli.ConvertRequest(py)
	_ = reqBody

	respBody, err := cli.Ask(reqBody)

	platformResp := response.UnionOpenUserRegisterValidateRowResponse{}

	err = cli.ConvertResponse(respBody, &platformResp)

	log.Println(platformResp, err)
}
