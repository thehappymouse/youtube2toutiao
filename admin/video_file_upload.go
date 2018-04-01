package admin

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"net/url"

	"bytes"
	"mime/multipart"
	"os"

	"github.com/gpmgo/gopm/modules/log"
)

func NewUploadRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,fr;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", "i.snssdk.com")
	req.Header.Set("Origin", "https://mp.toutiao.com")
	req.Header.Set("Proxy-Connection", "keep-alive")
	req.Header.Set("Referer", "https://mp.toutiao.com/profile_v3/xigua/upload-video")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")

	return req, err
}

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
	Mid       string        `json:"mid"`
}

//"e": "开始上传",
//"code": 0,
//"ref_id": "4e5e9ccd864647a08757877b8f25d5c2",
//"at": 1522552532251,
//"timestamp": 1522552541943,
//"f": "大圣归来.mp4",
//"url": "http://mp.toutiao.com/profile_v3/xigua/upload-video",
//"ubs": 0,
//"fs": 30780659,
//"log_id": null,
//"via": null,
//"text": "",
//"ua": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36",
//"cookie": "UM_distinctid=160490778164c9-05b9ef66bcee37-17396d56-13c680-1604907781752d; _ba=BA0.2-20180304-5110e-tnDHPUa2kposxvyLOJcn; _ga=GA1.2.830181196.1520146009; slaask-token-4ead732e8206531cc37f8ed618f08986=user-180713; uuid=\"w:d62fc43823c84299a706278b3d59c28f\"; _ga=GA1.3.830181196.1520146009; __tea_sdk__ssid=0688f82f-eef4-48ac-9e50-a404c8f4b4ff; tt_webid=6530933786749961741; __tea_sdk__user_unique_id=94107185706; login_flag=54f5a3dc3817e33a1ca3a82cb0f57ddb; _mp_auth_key=e6156cbc46c854b33b30b7a179f260ea; __utma=68473421.830181196.1520146009.1522552138.1522552138.1; __utmc=68473421; __utmz=68473421.1522552138.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __utmt=1; __utmb=68473421.3.10.1522552138",
//"log_data": [],
//"username": "大力出奇击",
//"mid": 1594097155488776

const startUrl = "http://i.snssdk.com/video/fedata/1/pgc/%s"

// 开始上传
func VideoLogStart(v *VideoFile, api *VideoApiData) {

	//{
	//"e":"开始上传", "code":0, "ref_id":"5a8097484845408481553c68e4a0d8b1", "at":1522504753845, "timestamp":1522504758903, "f":"大圣归来.mp4", "url":"http://mp.toutiao.com/profile_v3/xigua/upload-video", "ubs":0, "fs":30780659, "log_id":null, "via":null, "text":"", "ua":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36", "cookie":"UM_distinctid=160490778164c9-05b9ef66bcee37-17396d56-13c680-1604907781752d; _ba=BA0.2-20180304-5110e-tnDHPUa2kposxvyLOJcn; _ga=GA1.2.830181196.1520146009; slaask-token-4ead732e8206531cc37f8ed618f08986=user-180713; uuid=\"w:d62fc43823c84299a706278b3d59c28f\"; _ga=GA1.3.830181196.1520146009; __tea_sdk__ssid=0688f82f-eef4-48ac-9e50-a404c8f4b4ff; tt_webid=6530933786749961741; __tea_sdk__user_unique_id=94107185706; login_flag=54f5a3dc3817e33a1ca3a82cb0f57ddb; _mp_auth_key=e6156cbc46c854b33b30b7a179f260ea; ptcn_no=ff9e08c00277347be6cbfc1637a72197", "log_data":[], "username":"大力出奇击", "mid":1594097155488776
	//}
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
	ld.UserName = "大力出奇击"
	ld.Mid = "1594097155488776"

	logdata, _ := json.Marshal(ld)
	params := url.Values{}
	params.Set("log", string(logdata))
	data := params.Encode()
	fmt.Println(data)

	req, err := NewUploadRequest(http.MethodPost, fmt.Sprintf(startUrl, api.UploadID), strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Length", string(len(data)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	//	:
	//
	//Referer:
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err.Error())
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("开始上传的上报结果", string(body))
}

// 上传文件
func VideoUpload(v *VideoFile, api *VideoApiData) *VideoUploadResponse {
	uploadResp := &VideoUploadResponse{}

	uploadResp.StartTime = time.Now().Unix()
	resp, err := upload(v, api)
	if err != nil {
		panic(err)
	}
	uploadResp.EndTime = time.Now().Unix()

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	decoder.Decode(uploadResp)

	return uploadResp

}
func upload(v *VideoFile, api *VideoApiData) (*http.Response, error) {
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	// use the body_writer to write the Part headers to the buffer
	_, err := body_writer.CreateFormFile("video_file", v.Info.Name())
	if err != nil {
		fmt.Println("error writing to buffer")
		return nil, err
	}

	// the file data will be the second part of the body
	fh, err := os.Open(v.LocalPath)
	if err != nil {
		fmt.Println("error opening file")
		return nil, err
	}
	// need to know the boundary to properly close the part myself.
	boundary := body_writer.Boundary()
	//close_string := fmt.Sprintf("\r\n--%s--\r\n", boundary)
	close_buf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))

	// use multi-reader to defer the reading of the file data until
	// writing to the socket buffer.
	request_reader := io.MultiReader(body_buf, fh, close_buf)
	fi, err := fh.Stat()
	if err != nil {
		fmt.Printf("Error Stating file: %s", v.LocalPath)
		return nil, err
	}
	req, err := http.NewRequest("POST", api.UploadUrl, request_reader)
	if err != nil {
		return nil, err
	}

	// Set headers for multipart, and Content Length
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.ContentLength = fi.Size() + int64(body_buf.Len()) + int64(close_buf.Len())

	return http.DefaultClient.Do(req)
}

// 文件上传成功后的上报
func VideoLogSueecss(response *VideoUploadResponse, api *VideoApiData, v *VideoFile) {
	//{
	//"e":"开始上传", "code":0, "ref_id":"5a8097484845408481553c68e4a0d8b1", "at":1522504753845, "timestamp":1522504758903, "f":"大圣归来.mp4", "url":"http://mp.toutiao.com/profile_v3/xigua/upload-video", "ubs":0, "fs":30780659, "log_id":null, "via":null, "text":"", "ua":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36", "cookie":"UM_distinctid=160490778164c9-05b9ef66bcee37-17396d56-13c680-1604907781752d; _ba=BA0.2-20180304-5110e-tnDHPUa2kposxvyLOJcn; _ga=GA1.2.830181196.1520146009; slaask-token-4ead732e8206531cc37f8ed618f08986=user-180713; uuid=\"w:d62fc43823c84299a706278b3d59c28f\"; _ga=GA1.3.830181196.1520146009; __tea_sdk__ssid=0688f82f-eef4-48ac-9e50-a404c8f4b4ff; tt_webid=6530933786749961741; __tea_sdk__user_unique_id=94107185706; login_flag=54f5a3dc3817e33a1ca3a82cb0f57ddb; _mp_auth_key=e6156cbc46c854b33b30b7a179f260ea; ptcn_no=ff9e08c00277347be6cbfc1637a72197", "log_data":[], "username":"大力出奇击", "mid":1594097155488776
	//}
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
	ld.UserName = "大力出奇击"
	ld.Mid = "1594097155488776"

	logdata, _ := json.Marshal(ld)
	params := url.Values{}
	params.Set("log", string(logdata))
	data := params.Encode()
	fmt.Println(data)
	req, err := NewUploadRequest(http.MethodPost, fmt.Sprintf(startUrl, api.UploadID), strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Length", string(len(data)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err.Error())
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("开始上传的上报结果", string(body))
}
