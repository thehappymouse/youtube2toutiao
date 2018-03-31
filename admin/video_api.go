package admin

import (
	"net/http"
	"github.com/gpmgo/gopm/modules/log"
	"fmt"
	"io/ioutil"
	"strings"
	"encoding/json"
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

//func (r *VideoApiResp) UnmarshalJSON(body []byte) error {
//
//	json.Unmarshal(body, r)
//
//	r.ApiData = &VideoApiData{}
//	//if err == nil {
//		err := json.Unmarshal([]byte(r.Data), r.ApiData)
//	//}
//	return err
//	//return nil
//}

// 获取视频的上传路径
func VideoApi() {
	data := "json_data=%7B%22api%22%3A%22chunk_upload_info%22%7D"
	req, err := NewTiaoRequest(http.MethodPost,
		video_api, strings.NewReader(data))
	req.Header.Set("Content-Length", string(len(data)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	if err != nil {
		log.Error("make request err:", err)
		return
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err.Error())
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(resp.StatusCode, string(body))
	apiResp := &VideoApiResp{ApiData: &VideoApiData{}}

	//apiResp.UnmarshalJSON(body)
	json.Unmarshal(body, apiResp)
	json.Unmarshal([]byte(apiResp.Data), apiResp.ApiData)

	fmt.Println(apiResp)
	//fmt.Println(apiResp.Data.UploadUrl)
}
