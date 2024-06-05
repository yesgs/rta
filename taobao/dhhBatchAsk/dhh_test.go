package dhhBatchAsk

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

var imei string = "16284e0ef0874518c5e4de782022d25e"
var oaid string = "0723292b8dccd03301cc990288569686"
var idfa string = "1cbfbd8a27c5c0221b3c30b0f18260703"
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

func TestDHH(t *testing.T) {
	rtaOption := getRtaOption()

	var imeiList []string = []string{imei}
	var oaidList []string = []string{oaid}
	var idfaList []string = []string{idfa}
	var caidList []string

	oaidList = getOaidList()

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

func TestDHH2(t *testing.T) {
	rtaOption := getRtaOption()

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

func getOaidList() []string {
	var data = make([]string, 0, 10)
	data = append(data, "0723284761ac3443edf45c30dccae195")
	data = append(data, "072328640c9548b56834f3aab0e33351")
	data = append(data, "0723289441a08111e981c854e3c48959")
	data = append(data, "072328aff33d2ebb7ad1dc08ccb20d73")
	data = append(data, "072328c7f532e74917ebcf4d551974e6")
	data = append(data, "072328ecf72905ef3a4a442cce47a8e0")
	data = append(data, "072329176f3c3cce58a862b2a1a1e161")
	data = append(data, "0723292070f394c1cfc5aa59cfb79eeb")
	data = append(data, "0723292b8dccd03301cc990288569686")
	data = append(data, "072329d38ef84a14448375ecd2de2c75")
	data = append(data, "072329d615907e70af5b353148e1cc1d")
	data = append(data, "07232a2679aca93c39b542dc023ae8cd")
	data = append(data, "07232a28dc6ae45a80f2aa58ac48a4d7")
	data = append(data, "07232a291a0c57fddfdb2af2ae325d5d")
	return data
}
