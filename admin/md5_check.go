package admin

import (
	"net/http"
	"net/url"

	"dali.cc/toutiao/tools"
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

	request, err := NewTiaoRequest(http.MethodGet, u.String(), "")
	if err != nil {
		panic(err)
	}

	md5resp := &Md5Resp{}
	tools.DoRequestJson(request, md5resp)
	return md5resp
}
