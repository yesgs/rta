package autohome

import (
	"github.com/yesgs/rta"
	"log"
	"net/http"
	"testing"
)

func TestAutoHome(t *testing.T) {
	var channel = "aaa"

	cli := NewClient(&rta.Options{
		BaseUrl:    "http://uag-rta-active.autohome.com.cn/rta-pull-active/active/req/get",
		HttpMethod: "GET",
		HttpClient: &http.Client{
			Transport: nil,
		},
	})

	var idType = "OAID_MD5"
	var id = "300fb6cd0eb5cf001eb7e2ac127c3680"

	req := NewPlatformRequest(channel, idType, id)

	resp, err := cli.Ask(req)

	log.Println(string(resp), err)

}
