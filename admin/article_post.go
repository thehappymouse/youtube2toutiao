package admin

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"time"

	"dali.cc/toutiao/downloader"
	"dali.cc/toutiao/tools"
	"github.com/gpmgo/gopm/modules/log"
)

// 提交作品
const posturl = "http://mp.toutiao.com/core/article/edit_article_post/?downloader=mp&type=purevideo"

type ArticleForm struct {
	ArticleAdType int    `json:"article_ad_type"`
	Title         string `json:"title"`
	Abstract      string `json:"abstract"`
	// 分类
	Tag           string `json:"tag"`
	ExternLink    string `json:"extern_link"`
	IsFansArticle int    `json:"is_fans_article"`
	Content       string `json:"content"`
	AddThirdTitle int    `json:"add_third_title"`
	TimerStatus   int    `json:"timer_status"`
	// 2018-04-01 09:58
	TimerTime            string `json:"timer_time"`
	RecommendAutoAnalyse int    `json:"recommend_auto_analyse"`
	// 标签，多个以逗号分隔，例如： 大圣归来;唯美MV
	ArticleLabel  string `json:"article_label"`
	FromDiagnosis int    `json:"from_diagnosis"`
	// 和 Tag什么关系
	ArticleType int `json:"article_type"`
	Praise      int `json:"praise"`
	PgcDebut    int `json:"pgc_debut"`
	Save        int `json:"save"`
}

type ArticleResult struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Data    string `json:"data"`
}

// 将结构体转换成 FormQuery
func struct2form(f interface{}) string {

	params := url.Values{}
	t := reflect.TypeOf(f)
	v := reflect.ValueOf(f)
	for k := 0; k < t.NumField(); k++ {
		field := t.Field(k)
		value := fmt.Sprintf("%v", v.Field(k).Interface())
		params.Set(field.Tag.Get("json"), value)
	}
	return params.Encode()
}

func ArticlePost(videofile downloader.VideoFile, videoapi *VideoApiData, uploadResponse *VideoUploadResponse) {

	form := ArticleForm{
		ArticleAdType: 3,
		Title:         videofile.Title,
		Abstract:      videofile.Desc,
		Tag:           "video_animation",
		Content:       `<p>{!-- PGC_VIDEO:{"sp":"toutiao","vid":"%s","vu":"%[1]s","thumb_url":"%s","src_thumb_uri":"%[2]s","vname":"%s"} --}</p>`,
		TimerTime:     time.Now().Format(("2006-01-02 15:04")),
		ArticleLabel:  "动漫;搞笑;",
		ArticleType:   1,
		Save:          1,
	}
	form.Content = fmt.Sprintf(form.Content, videoapi.UploadID, uploadResponse.PosterUri, videofile.Title)
	fmt.Println(form)

	data := struct2form(form)

	log.Warn("PostRawData: %s", data)
	req, err := NewTiaoRequest(http.MethodPost, posturl, data)
	if err != nil {
		panic(err)
	}

	result := &ArticleResult{}
	tools.DoRequestJson(req, result)
	log.Warn("作品提交结果: %v", result)
}
