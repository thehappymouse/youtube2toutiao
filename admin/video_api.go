package admin

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gpmgo/gopm/modules/log"
)

const video_api = "http://mp.toutiao.com/video/video_api/"

type VideoApiData struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reason    string `json:"reason"`
	UploadUrl string `json:"upload_url"`
	UploadID  string `json:"upload_id"`
}

type VideoApiResp struct {
	Url     string `json:"url"`
	Message string `json:"message"`
	Data    string `json:"data"`
	ApiData *VideoApiData
}

// 获取视频的上传路径
func VideoApi() *VideoApiData {
	data := "json_data=%7B%22api%22%3A%22chunk_upload_info%22%7D"
	req, err := NewTiaoRequest(http.MethodPost,
		video_api, strings.NewReader(data))
	req.Header.Set("Content-Length", string(len(data)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	if err != nil {
		log.Error("make request err:", err)
		return nil
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err.Error())
	}
	defer resp.Body.Close()

	apiResp := &VideoApiResp{ApiData: &VideoApiData{}}
	json.NewDecoder(resp.Body).Decode(apiResp)

	// 需要再次处理
	json.Unmarshal([]byte(apiResp.Data), apiResp.ApiData)

	return apiResp.ApiData
}
