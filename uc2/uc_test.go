package uc2

import (
	"github.com/yesgs/rta"
	"log"
	"net/http"
	"testing"
)

func TestUc(t *testing.T) {
	var channel = "channel"

	ucCli := NewClient(&rta.Options{
		BaseUrl:    "https://rta.uc.cn/rta/open/" + channel,
		HttpMethod: "GET",
		HttpClient: &http.Client{
			Transport: nil,
		},
	})

	var secret = "300fb6cd0eb5cf001eb7e2ac127c3680"
	var oaidMd5 = "300fb6cd0eb5cf001eb7e2ac127c3680"
	var caid = ""

	req := NewPlatformRequest(channel, secret, oaidMd5, caid)

	resp, err := ucCli.Ask(req)

	log.Println(string(resp), err)

}
