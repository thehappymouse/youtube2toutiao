package translator

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"net/http"
	"net/url"
	"strings"

	"strconv"

	"fmt"

	"dali.cc/toutiao/tools"
	"github.com/gpmgo/gopm/modules/log"
)

const (
	api_url = "http://openapi.youdao.com/api"
)

type YouDao struct {
	AppKey, SecKey string
}
type youdaoResult struct {
	Query       string   `json:"query"`
	Translation []string `json:"translation"`
	ErrorCode   int      `json:"error_code"`
}

func (t *YouDao) builSign(in, salt string) string {

	var buf bytes.Buffer
	buf.WriteString(t.AppKey)
	buf.WriteString(in)
	buf.WriteString(salt)
	buf.WriteString(t.SecKey)

	md5Ctx := md5.New()
	md5Ctx.Write(buf.Bytes())
	cipherStr := md5Ctx.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

func (t *YouDao) Translate(in string) string {
	params := url.Values{}

	log.Warn("Translate:[%s]", in)

	params.Set("q", in)
	params.Set("from", "auto")
	params.Set("to", "zh-CHS")
	params.Set("appKey", t.AppKey)
	params.Set("salt", strconv.Itoa(rand.Intn(10000)))
	params.Set("sign", t.builSign(in, params.Get("salt")))

	req, _ := http.NewRequest(http.MethodPost, api_url, strings.NewReader(params.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	result := youdaoResult{}

	tools.DoRequestJson(req, &result)
	fmt.Println("--------", result)
	return result.Translation[0]
}
