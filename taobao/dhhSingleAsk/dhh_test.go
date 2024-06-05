package dhhSingleAsk

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

var channel = "2200803434774"
var appKey = "32560985"
var appSecret = "df49ecf0b667bedd74010ef458937d66"
var adspaceId = "1775194"

var imei string = ""
var oaid string = "649188098240f775f4014174bd0de821"
var idfa string = ""
var caid string

func getRtaOption() rta.Options {
	rtaOption := rta.Options{
		BaseUrl:    "http://gw.api.taobao.com/router/rest",
		HttpMethod: http.MethodPost,
		HttpClient: &http.Client{
			Timeout: 1 * time.Second,
			Transport: &http.Transport{
				DisableKeepAlives:     false,
				DisableCompression:    false,
				MaxIdleConns:          10240,
				MaxIdleConnsPerHost:   10240,
				MaxConnsPerHost:       81920,
				IdleConnTimeout:       time.Second * 60,
				ResponseHeaderTimeout: time.Second * 60,
				//WriteBufferSize:       0,
				//ReadBufferSize:        0,
				ForceAttemptHTTP2: true,
				TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			},
		},
	}

	return rtaOption
}

func TestDHHSingle(t *testing.T) {
	rtaOption := rta.Options{
		BaseUrl:    "http://gw.api.taobao.com/router/rest",
		HttpMethod: http.MethodPost,
		HttpClient: &http.Client{
			Timeout: 1 * time.Second,
			Transport: &http.Transport{
				DisableKeepAlives:     false,
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

	req := NewPlatformRequest(adspaceId, channel, imei, oaid, idfa, caid)

	taoBaoClient := taobao.NewClient(&rtaOption, appKey, appSecret)

	cli := NewClient(taoBaoClient)

	reqBody, err := cli.ConvertRequest(req)
	_ = reqBody

	platformResp := response.TaobaoUsergrowthDhhDeliveryAskResponse{}
	err = rta.MakeRequest(cli, req, &platformResp)
	log.Println(err, err)

}
