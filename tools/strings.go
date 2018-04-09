package tools

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

type Describe string

func (s Describe) String() string {
	return string(s)
}

// 删除字符串中的 url 数据
func (s Describe) TrimUrl() Describe {
	reg := regexp.MustCompile(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)
	r := reg.ReplaceAllString(s.String(), "")
	r = strings.Trim(r, "\n")
	return s
}

// 删除字符串中的 url 数据
func TrimUrl(s string) string {
	reg := regexp.MustCompile(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)
	r := reg.ReplaceAllString(s, "")
	r = strings.Trim(r, "\n")
	return r
}

// 只取文件的名称部分，去掉后缀
func FileNameOnly(s string) string {
	start := strings.LastIndex(s, "/")
	if start == -1 {
		start = 0
	}
	end := strings.LastIndex(s, ".")
	return s[start:end]
}

// 如果长了，减掉
func CutByUtf8(s string, length int) string {
	if utf8.RuneCountInString(s) > length {
		b := []rune(s)
		s = string(b[0:length-3]) + "..."
	}
	return s
}
