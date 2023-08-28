package vip

import (
	"github.com/yesgs/rta"
	"log"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func Test_f1(t *testing.T) {
	rtaOption := rta.Options{
		BaseUrl:    "https://127.0.0.1:9999/v1",
		HttpMethod: http.MethodPost,
		HttpClient: &http.Client{
			Timeout: 1000 * time.Millisecond,
			Transport: &http.Transport{
				DisableKeepAlives:     true,
				DisableCompression:    false,
				MaxIdleConns:          1024,
				MaxIdleConnsPerHost:   1024,
				MaxConnsPerHost:       1024,
				IdleConnTimeout:       time.Second * 60,
				ResponseHeaderTimeout: time.Second * 60,
				ForceAttemptHTTP2:     false,
			},
		},
	}

	cli := NewClient(&rtaOption)
	var tokenList = []string{"a", "a"}
	var didType = 0
	var did = "0009e6508ad46de8eb472b447563c61b"
	var id = strconv.FormatInt(time.Now().UnixNano(), 10)
	py := NewPlatformRequest(id, tokenList, didType, did)
	reqBody, err := cli.ConvertRequest(py)
	_ = reqBody
	log.Println(string(reqBody.([]byte)), err)

	respBody, err := cli.Ask(reqBody)
	log.Println(string(respBody), err)

	//err = cli.ConvertResponse(respBody, &platformResp)
	//log.Println(platformResp, err)
}
