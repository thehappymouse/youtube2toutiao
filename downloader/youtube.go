package downloader

import (
	"strings"

	"io/ioutil"

	"os"

	"dali.cc/toutiao/tools"
	"github.com/gpmgo/gopm/modules/log"
)

// 在打印信息中查找文件名
func ParseFileName(ss []string) (title string) {
	key := "Destination"
	for _, s := range ss {
		if index := strings.Index(s, key); index != -1 {
			title = s[(index + len(key) + 1):]
			// 可以只取文件名称
			title = strings.TrimSpace(title)
			break
		}
		// already 下载
		already := "has already been downloaded"
		if index := strings.Index(s, already); index != -1 {
			title = s[len("[download]"):index]
			title = strings.TrimSpace(title)
			break
		}
	}
	return title
}

var DownloadCommand = "./downloader/bin/youtube.sh"

func Download(url string) (ok bool, v VideoFile) {

	params := []string{url}
	ok, ss := tools.ExecCommand(DownloadCommand, params)
	if !ok {
		log.Warn("%v", ss)
		return
	}
	filename := ParseFileName(ss)
	basename := tools.FileNameOnly(filename)

	desc, err := ioutil.ReadFile(basename + ".description")
	if err != nil {
		log.Error("打开描述文件[%s]失败", err)
	}

	_, md5 := tools.Md5(filename)

	v.FilePath = filename
	v.Title = basename
	v.Desc = tools.TrimUrl(string(desc))
	v.Md5 = md5

	info, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}
	v.FileSize = info.Size()

	return true, v
}
