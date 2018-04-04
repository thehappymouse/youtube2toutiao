package tools

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
)

func Md5(file string) (error, string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err, ""
	}

	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)

	return nil, hex.EncodeToString(cipherStr)
}
