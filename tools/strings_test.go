package tools

import (
	"strings"
	"testing"
)

func TestTrimUrl(t *testing.T) {
	var s Describe
	s = `https://youtu.be/lvObk195BQU
The most funny Memorable and interesting moments from the cartoon about the minions in one short film
https://youtu.be/lvObk195BQU`

	r := s.TrimUrl()
	if strings.Index(string(r), "http://") != -1 {
		t.Errorf("%s 中依然包含url信息", r)
	}
}

func TestFileNameOnly(t *testing.T) {
	tests := []struct {
		source, right string
	}{
		{"test abc.a .mp4", "test abc.a "},

		{"中文文件 abc. 你怎么办 .mp4", "中文文件 abc. 你怎么办 "},
	}

	for _, ts := range tests {
		if r := FileNameOnly(ts.source); r != ts.right {
			t.Errorf("取文件名部分错误，想要: %s, 得到: %s", ts.right, r)
		}
	}
}
