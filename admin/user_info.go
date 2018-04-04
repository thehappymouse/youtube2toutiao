package admin

import (
	"net/http"

	"encoding/json"

	"dali.cc/toutiao/tools"
	"github.com/gpmgo/gopm/modules/log"
)

// 获取用户信息

const infoUrl = "http://mp.toutiao.com/get_media_info/"

type UserInfo struct {
	HttpsAvatarUrl string `json:"https_avatar_url"`
	ID             int64  `json:"id"`
	DisplayName    string `json:"display_name"`
}

type MediaInfo struct {
	Media *UserInfo `json:"media"`
}

// 当前用户
var User *UserInfo

func LoadUserInfo() {
	req, err := NewTiaoRequest(http.MethodGet, infoUrl, "")
	if err != nil {
		panic(err)
	}
	result := &CommonResult{}
	tools.DoRequestJson(req, result)
	media := &MediaInfo{}

	// todo 多余，需要找到更好的方法
	tmpb, _ := json.Marshal(result.Data)
	json.Unmarshal(tmpb, media)

	User = media.Media
	log.Warn("用户信息：%v", User)
}
