package talkingData

import (
	"github.com/yesgs/rta"
	"log"
	"net/http"
	"testing"
	"time"
)

func Test_f1(t *testing.T) {
	rtaOption := rta.Options{
		BaseUrl:    "http://127.0.0.1:8080/dmp-ad-frontend-server/frontend/get",
		HttpMethod: http.MethodPost,
		HttpClient: &http.Client{
			Timeout: 3 * time.Second,
			Transport: &http.Transport{
				DisableKeepAlives:     true,
				DisableCompression:    false,
				MaxIdleConns:          1024,
				MaxIdleConnsPerHost:   1024,
				MaxConnsPerHost:       8192,
				IdleConnTimeout:       time.Second * 60,
				ResponseHeaderTimeout: time.Second * 60,
				ForceAttemptHTTP2:     false,
			},
		},
	}

	cli := NewClient(&rtaOption)
	py := Request{
		MediaId:      "674",
		CampaignId:   []string{"xdfGIMnq"},
		DeviceIdType: 78,
		DeviceId:     "58D2C273ECA672EB3A65134BA3285672",
	}

	reqBody, err := cli.ConvertRequest(py)
	_ = reqBody
	log.Println(string(reqBody.([]byte)), err)

	respBody, err := cli.Ask(reqBody)
	log.Println(string(respBody), err)

	//err = cli.ConvertResponse(respBody, &platformResp)
	//log.Println(platformResp, err)
}
