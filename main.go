package main

import (
	"flag"
	"fmt"
	"os"

	"dali.cc/toutiao/admin"
	"dali.cc/toutiao/downloader"
	"dali.cc/toutiao/tools"
	"dali.cc/toutiao/translator"
	"github.com/gpmgo/gopm/modules/log"
)

var dao = &translator.YouDao{
	AppKey: "6a0f0aec8e860c65",
	SecKey: "vTrsGcDDmD0X6RIUUpCi0oEGazF30BOz",
}
var id = flag.String("url", "", "请输入一个 youtube 地址")

func main() {
	flag.Parse()
	ok, video := downloader.Download(*id)
	if !ok {
		log.Error("下载源文件出错")
		os.Exit(1)
	}

	// 分析标题和内容
	// todo 可以扩展string 自由调用吗

	fmt.Println(video.Title)

	video.Title = translator.Translate(dao, video.Title)
	video.Desc = translator.Translate(dao, video.Desc)
	video.Title = tools.CutByUtf8(video.Title, 30)
	video.Desc = tools.CutByUtf8(video.Desc, 300)

	// 文件准备完成
	admin.LoadUserInfo()

	md5Resp := admin.Md5Check(video.Md5)
	if !md5Resp.IsUniq {
		log.Warn("Video Already use in [%s]", md5Resp.Data)
		return
	}

	videoapi := admin.VideoApi()

	log.Warn("video/api: %v", videoapi)

	admin.VideoLogStart(&video, videoapi)

	uploadResponse := admin.VideoUpload(&video, videoapi)

	log.Warn("uploadResponse: %v", uploadResponse)

	admin.VideoLogSueecss(uploadResponse, videoapi, &video)

	admin.ArticlePost(video, videoapi, uploadResponse)
}
