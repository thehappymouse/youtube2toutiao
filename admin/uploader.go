package admin

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"net/url"

	"github.com/gpmgo/gopm/modules/log"
)

// 上传文件的返回结果
type VideoUploadResponse struct {
	Bytes       int    `json:"bytes"`
	Code        int    `json:"code"`
	Crc32       int    `json:"crc_32"`
	ExpectBytes int    `json:"expect_bytes"`
	Message     string `json:"message"`
	PosterUri   string `json:"poster_uri"`
	PosterUrl   string `json:"poster_url"`
	Reason      string `json:"reason"`
	Url         string `json:"url"`
	StartTime   int64  `json:"-"`
	EndTime     int64  `json:"-"`
}

// 上传前，后，取消都要发送的消息
type logData struct {
	E     string `json:"e"`
	Code  int    `json:"code"`
	RefID string `json:"ref_id"`
	// 上传开始时间
	TimeStart int64 `json:"ts,omitempty"`
	// 上传结束时间
	TimeEnd   int64         `json:"te,omitempty"`
	At        int64         `json:"at"`
	Timestamp int64         `json:"timestamp"`
	F         string        `json:"f"`
	Url       string        `json:"url"`
	Ubs       int           `json:"ubs"`
	FileSize  int64         `json:"fs"`
	LogID     interface{}   `json:"log_id"`
	Via       interface{}   `json:"via"`
	Text      string        `json:"text"`
	UserAgent string        `json:"ua"`
	Cookie    string        `json:"cookie"`
	LogData   []interface{} `json:"log_data"`
	UserName  string        `json:"username"`
	Mid       int64         `json:"mid"`
}

const logUrl = "http://i.snssdk.com/video/fedata/1/pgc/%s"

// 开始上传
func VideoLogStart(v *VideoFile, api *VideoApiData) {

	nowtime := time.Now().Unix()
	ld := logData{}
	ld.E = "开始上传"
	ld.RefID = api.UploadID
	ld.At = nowtime
	ld.Timestamp = nowtime
	ld.F = v.Info.Name()
	ld.Url = "http://mp.toutiao.com/profile_v3/xigua/upload-video"
	ld.FileSize = v.Info.Size()
	ld.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36"
	ld.Cookie = string(getCookie())
	ld.LogData = []interface{}{}
	ld.UserName = User.DisplayName
	ld.Mid = User.ID

	logdata, _ := json.Marshal(ld)
	params := url.Values{}
	params.Set("log", string(logdata))
	data := params.Encode()
	fmt.Println(data)

	req, err := NewLogRequest(http.MethodPost, fmt.Sprintf(logUrl, api.UploadID), data)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Length", string(len(data)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	doRequest2(req, func(reader io.ReadCloser) {
		body, _ := ioutil.ReadAll(reader)
		log.Warn("开始上传的上报结果: %s", string(body))
	})
}

// 上传文件
func VideoUpload(v *VideoFile, api *VideoApiData) *VideoUploadResponse {
	uploadResp := &VideoUploadResponse{}
	log.Warn("上传文件")
	uploadResp.StartTime = time.Now().Unix()

	req, err := NewUploadFileRequest(v, api.UploadUrl)
	if err != nil {
		panic(err)
	}
	doRequest(req, uploadResp)
	uploadResp.EndTime = time.Now().Unix()
	log.Warn("上传完成")
	return uploadResp

}

// 文件上传成功后的上报
func VideoLogSueecss(response *VideoUploadResponse, api *VideoApiData, v *VideoFile) {
	nowtime := time.Now().Unix()
	ld := logData{}
	ld.E = "文件上传成功"
	ld.TimeStart = response.StartTime
	ld.TimeEnd = response.EndTime
	ld.RefID = api.UploadID
	ld.At = nowtime
	ld.Timestamp = nowtime
	ld.F = v.Info.Name()
	ld.Url = "http://mp.toutiao.com/profile_v3/xigua/upload-video"
	ld.FileSize = v.Info.Size()
	ld.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36"
	ld.Cookie = string(getCookie())
	ld.LogData = []interface{}{response}
	ld.UserName = User.DisplayName
	ld.Mid = User.ID

	logdata, _ := json.Marshal(ld)
	params := url.Values{}
	params.Set("log", string(logdata))
	data := params.Encode()
	req, err := NewLogRequest(http.MethodPost, fmt.Sprintf(logUrl, api.UploadID), data)
	if err != nil {
		panic(err)
	}

	doRequest2(req, func(reader io.ReadCloser) {
		body, _ := ioutil.ReadAll(reader)
		log.Warn("上传完成的上报结果: %s", string(body))
	})
}
