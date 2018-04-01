package admin

import (
	"net/http"
	"net/url"

	"encoding/json"

	"github.com/gpmgo/gopm/modules/log"
)

const checkUrl = "http://mp.toutiao.com/video/video_uniq_api/"

type Md5Resp struct {
	Message string `json:"message"`
	Data    string `json:"data"`
	IsUniq  bool   `json:"is_uniq"`
}

func Md5Check(md5 string) *Md5Resp {
	u, _ := url.Parse(checkUrl)

	q := u.Query()
	q.Set("md5", md5)
	u.RawQuery = q.Encode()

	request, err := NewTiaoRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		panic(err)
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	if response.StatusCode != http.StatusOK {
		log.Error("%s %s %s", u.String(), "result code ==> ", response.StatusCode)
		return nil
	}
	md5resp := &Md5Resp{}
	der := json.NewDecoder(response.Body)
	der.Decode(md5resp)
	return md5resp
}
