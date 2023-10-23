package uc

import (
	"encoding/json"
	"github.com/yesgs/rta"
	"log"
	"net/http"
	"testing"
)

func TestUc(t *testing.T) {

	ucCli := NewClient(&rta.Options{
		BaseUrl:    "https://ugp-cloud.uc.cn/rta/aaa/v1",
		HttpMethod: "POST",
		HttpClient: &http.Client{
			Transport: nil,
		},
	})

	var channel = "300fb6cd0eb5cf001eb7e2ac127c3680"
	var platform = "ANDROID"
	var didType = "OAID_MD5"
	var did = "300fb6cd0eb5cf001eb7e2ac127c3680"

	req := NewPlatformRequest(channel, platform, didType, did)

	body, _ := json.Marshal(req)
	resp, err := ucCli.Ask(body)

	log.Println(string(resp), err)

}
