package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"dali.cc/toutiao/admin"
	"github.com/gpmgo/gopm/modules/log"
)

func ContentList(req *http.Request)  {
	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Error("http get error: " , err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode, (string(body)))
}
func main() {
	admin.VideoApi()
}
