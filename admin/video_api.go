package admin

import (
	"encoding/json"
	"net/http"

	"github.com/gpmgo/gopm/modules/log"
)

const apiurl = "http://mp.toutiao.com/video/apiurl/"

type VideoApiData struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reason    string `json:"reason"`
	UploadUrl string `json:"upload_url"`
	UploadID  string `json:"upload_id"`
}

type VideoApiResult struct {
	Url     string `json:"url"`
	Message string `json:"message"`
	Data    string `json:"data"`
	ApiData *VideoApiData
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

	apiResp := &VideoApiResult{ApiData: &VideoApiData{}}
	doRequest(req, apiResp)

	// 需要再次处理
	json.Unmarshal([]byte(apiResp.Data), apiResp.ApiData)

	return apiResp.ApiData
}
