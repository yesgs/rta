package consumerMatch

import (
	"crypto/tls"
	"github.com/yesgs/rta"
	"github.com/yesgs/rta/taobao"
	"github.com/yesgs/rta/taobao/response"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestConsumerMatch(t *testing.T) {
	rtaOption := rta.Options{
		BaseUrl:    "http://gw.api.taobao.com/router/rest",
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

	var adZoneId int64 = 115033550310
	var deviceType = "OAID"
	var deviceValue = "94f0580f3a3fbe5f3a3bf5dedcb8574a"
	var strategyIdList = "87434415627"

	req := NewPlatformRequest(adZoneId, deviceType, deviceValue, strategyIdList, "")

	var appKey = "env(appKey)"
	var appSecret = "env(appSecret)"

	taoBaoClient := taobao.NewClient(&rtaOption, appKey, appSecret)

	cli := NewClient(taoBaoClient)

	reqBody, err := cli.ConvertRequest(req)
	_ = reqBody

	respBody, err := cli.Ask(reqBody)

	platformResp := response.TaobaoTbkRtaConsumerMatchResponse{}

	err = cli.ConvertResponse(respBody, &platformResp)

	log.Println(platformResp, err)

}
