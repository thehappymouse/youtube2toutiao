package admin

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"dali.cc/toutiao/downloader"
	"github.com/gpmgo/gopm/modules/log"
)

// 头条通用的响应结构
// Data一般存放的返回数据
type CommonResult struct {
	Url     string `json:"url"`
	Code    int    `json:"code"`
	Now     int    `json:"now"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
	//Data    json.RawMessage `json:"data"`
	Data interface{} `json:"data"`
}

func getCookie() []byte {
	data, err := ioutil.ReadFile("cookie.txt")
	if err != nil {
		log.Error("读取 cookie.txt 文件失败 :", err)
		os.Exit(0)
	}
	return data
}

// 创建访问头条的普通请求
func NewTiaoRequest(method, url string, body string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,fr;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", string(getCookie()))
	req.Header.Set("Host", "mp.toutiao.com")
	req.Header.Set("Referer", "https://mp.toutiao.com/profile_v3/xigua/content-manage")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")

	if method == http.MethodPost && body != "" {
		req.Header.Set("Content-Length", string(len(body)))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	}

	return req, err
}

// 创建上传文件的请求
func NewUploadFileRequest(v *downloader.VideoFile, url string) (*http.Request, error) {
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	// use the body_writer to write the Part headers to the buffer
	_, err := body_writer.CreateFormFile("video_file", v.FilePath)
	if err != nil {
		fmt.Println("error writing to buffer")
		return nil, err
	}

	// the file data will be the second part of the body
	fh, err := os.Open(v.FilePath)
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
		fmt.Printf("Error Stating file: %s", v.FilePath)
		return nil, err
	}
	req, err := http.NewRequest("POST", url, request_reader)
	if err != nil {
		return nil, err
	}

	// Set headers for multipart, and Content Length
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.ContentLength = fi.Size() + int64(body_buf.Len()) + int64(close_buf.Len())
	return req, err
}

// 创建 Log 类的请求， 都是POST
func NewLogRequest(method, url, data string) (*http.Request, error) {

	req, err := http.NewRequest(method, url, strings.NewReader(data))

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,fr;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", "i.snssdk.com")
	req.Header.Set("Origin", "https://mp.toutiao.com")
	req.Header.Set("Proxy-Connection", "keep-alive")
	req.Header.Set("Referer", "https://mp.toutiao.com/profile_v3/xigua/upload-video")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")

	if http.MethodPost == method {
		req.Header.Set("Content-Length", string(len(data)))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	}

	return req, err
}
