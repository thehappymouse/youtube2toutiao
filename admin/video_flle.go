package admin

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
)

type VideoFile struct {
	LocalPath string
	Info      os.FileInfo
	Md5       string
}

func NewVideoFile(path string) VideoFile {
	v := VideoFile{}
	v.LocalPath = path
	info, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	v.Info = info
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)

	v.Md5 = hex.EncodeToString(cipherStr)
	return v
}
