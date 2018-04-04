package translator

import (
	"testing"
)

func TestTranslator_Translate(t *testing.T) {
	dao := &YouDao{
		AppKey: "6a0f0aec8e860c65",
		SecKey: "vTrsGcDDmD0X6RIUUpCi0oEGazF30BOz",
	}

	fn := func(tan Translator, str string) string {
		return tan.Translate(str)
	}

	tests := []struct {
		In, Out string
	}{
		{"chinese people", "中国人民"},
		{"how are you?", "你好吗?"},
	}
	//todo 对比结果可以找一个相拟度算法
	for _, ts := range tests {
		if result := fn(dao, ts.In); result != ts.Out {
			t.Errorf("%s 翻译的结果应该是 %s, 却得到了:%s", ts.In, ts.Out, result)
		}
	}
}
