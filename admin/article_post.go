package admin

import "net/http"
import (
	"fmt"
	"net/url"
	"strings"

	"reflect"

	"encoding/json"

	"github.com/gpmgo/gopm/modules/log"
)

// 提交作品

const posturl = "http://mp.toutiao.com/core/article/edit_article_post/?source=mp&type=purevideo"

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

func ArticlePost(f ArticleForm) {
	// 将 f 压入params
	params := url.Values{}
	t := reflect.TypeOf(f)
	v := reflect.ValueOf(f)
	for k := 0; k < t.NumField(); k++ {
		field := t.Field(k)
		value := fmt.Sprintf("%v", v.Field(k).Interface())
		params.Set(field.Tag.Get("json"), value)
	}
	data := params.Encode()
	log.Warn("PostRawData: %s", data)
	req, err := NewTiaoRequest(http.MethodPost, posturl, strings.NewReader(data))
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
	r := &ArticleResult{}
	der := json.NewDecoder(resp.Body)
	der.Decode(r)
	log.Warn("节上传结果: %v", r)
}
