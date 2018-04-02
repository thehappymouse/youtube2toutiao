package admin

import (
	"net/http"

	"fmt"

	"encoding/json"

	"github.com/gpmgo/gopm/modules/log"
)

const apiurl = "http://mp.toutiao.com/video/video_api/"

type VideoApiData struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reason    string `json:"reason"`
	UploadUrl string `json:"upload_url"`
	UploadID  string `json:"upload_id"`
}

// 获取视频的上传路径
func VideoApi() *VideoApiData {
	data := "json_data=%7B%22api%22%3A%22chunk_upload_info%22%7D"
	req, err := NewTiaoRequest(http.MethodPost,
		apiurl, data)

	if err != nil {
		log.Error("make request err:", err)
		return nil
	}

	result := &CommonResult{}
	doRequest(req, result)

	apidata := &VideoApiData{}
	// 需要再次处理
	json.Unmarshal([]byte(fmt.Sprintf("%s", result.Data)), apidata)

	return apidata
}
