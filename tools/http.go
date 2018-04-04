package tools

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gpmgo/gopm/modules/log"
)

// 执行一个请求操作，执行回调
func DoReqeustByFn(req *http.Request, callback func(reader io.ReadCloser)) {
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Error("%s %s %s", req.URL, "result code ==> ", response.StatusCode)
	}
	callback(response.Body)
}

// 取得一个请的结果，尝试转换成 json 对象
func DoRequestJson(req *http.Request, r interface{}) {
	DoReqeustByFn(req, func(reader io.ReadCloser) {
		data, _ := ioutil.ReadAll(reader)
		json.Unmarshal(data, r)
	})
}
