package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"dali.cc/toutiao/admin"
	"github.com/gpmgo/gopm/modules/log"
)

func ContentList(req *http.Request) {
	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Error("http get error: ", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode, (string(body)))
}

func main() {
	admin.LoadUserInfo()
	videofile := admin.NewVideoFile("/Users/apple/Downloads/toutiaohao/W_sCTMEcb9SSY_zh.mp4")

	md5Resp := admin.Md5Check(videofile.Md5)
	if !md5Resp.IsUniq {
		log.Warn("Video Already use in [%s]", md5Resp.Data)
		return
	}

	videoapi := admin.VideoApi()

	log.Warn("video/api: %v", videoapi)

	admin.VideoLogStart(&videofile, videoapi)

	uploadResponse := admin.VideoUpload(&videofile, videoapi)

	log.Warn("uploadResponse: %v", uploadResponse)

	admin.VideoLogSueecss(uploadResponse, videoapi, &videofile)

	admin.ArticlePost(videofile, videoapi, uploadResponse)
}
