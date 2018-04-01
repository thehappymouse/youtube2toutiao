package admin

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gpmgo/gopm/modules/log"
)

func getCookie() []byte {
	data, err := ioutil.ReadFile("cookie.txt")
	if err != nil {
		log.Error("读取 cookie.txt 文件失败 :", err)
		os.Exit(0)
	}
	return data
}

func NewTiaoRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,fr;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", string(getCookie()))
	req.Header.Set("Host", "mp.toutiao.com")
	req.Header.Set("Referer", "https://mp.toutiao.com/profile_v3/xigua/content-manage")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")

	return req, err
}
