package dhh

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

var channel = "22200803434774"
var adspaceId = "11842023"
var appKey = "332560985"
var appSecret = "ddf49ecf0b667bedd74010ef458937d66"
var imei string = "16284e0ef0874518c5e4de782022d25e"
var oaid string = "000001bcdf2ace177bc6fbb92e5d2893"
var idfa string = "1cbfbd8a27c5c0221b3c30b0f18260703"
var caid string

func TestDHH(t *testing.T) {
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

	var imeiList []string = []string{imei}
	var oaidList []string = []string{oaid}
	var idfaList []string = []string{idfa}
	var caidList []string

	req := NewPlatformRequest(adspaceId, channel, imeiList, oaidList, idfaList, caidList)

	taoBaoClient := taobao.NewClient(&rtaOption, appKey, appSecret)

	cli := NewClient(taoBaoClient)

	reqBody, err := cli.ConvertRequest(req)
	_ = reqBody

	platformResp := response.TaobaoUsergrowthDhhDeliveryBatchaskResponse{}
	err = rta.MakeRequest(cli, req, &platformResp)
	log.Println(err, err)

	//respBody, err := cli.Ask(reqBody)

}

func TestDHHSingle(t *testing.T) {
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

	req := NewPlatformRequestSingleQuery(adspaceId, channel, imei, oaid, idfa, caid)

	taoBaoClient := taobao.NewClient(&rtaOption, appKey, appSecret)

	cli := NewClient(taoBaoClient)

	reqBody, err := cli.ConvertRequest(req)
	_ = reqBody

	platformResp := response.TaobaoUsergrowthDhhDeliveryAskResponse{}
	err = rta.MakeRequest(cli, req, &platformResp)
	log.Println(err, err)

}
