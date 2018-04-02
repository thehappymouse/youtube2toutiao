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
	videofile := admin.NewVideoFile("/Users/apple/Downloads/toutiaohao/W_sCTMEcb9SSY_zh.mp4")

	md5Resp := admin.Md5Check(videofile.Md5)
	if !md5Resp.IsUniq {
		log.Warn("Video Already use in [%s]", md5Resp.Data)
		return
	}

	videoapi := admin.VideoApi()

	log.Warn("video/api: %s", videoapi)

	admin.VideoLogStart(&videofile, videoapi)

	uploadResponse := admin.VideoUpload(&videofile, videoapi)

	log.Warn("uploadResponse: %s", uploadResponse)

	admin.VideoLogSueecss(uploadResponse, videoapi, &videofile)

	form := admin.ArticleForm{
		ArticleAdType: 3,
		Title:         videofile.Info.Name(),
		Abstract:      videofile.Info.Name(),
		Tag:           "video_music",
		Content:       `<p>{!-- PGC_VIDEO:{"sp":"toutiao","vid":"%s","vu":"%[1]s","thumb_url":"%s","src_thumb_uri":"%[2]s","vname":"%s"} --}</p>`,
		TimerTime:     "2018-04-02 17:29",
		ArticleLabel:  "动漫;唯美MV",
		ArticleType:   1,
		Save:          1,
	}
	form.Content = fmt.Sprintf(form.Content, videoapi.UploadID, uploadResponse.PosterUri, videofile.Info.Name())
	fmt.Println(form.Content)
	admin.ArticlePost(form)
}
