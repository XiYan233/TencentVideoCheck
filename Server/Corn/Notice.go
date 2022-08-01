package Corn

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Notice(cookie string, notice string, noticeToken string, msg string) {
	if notice == "close" {
		log.Print("未开启通知\n")
	} else if notice == "WeChat" {

		client := &http.Client{}
		var data = strings.NewReader(`{
		   "msgtype": "markdown",
			"markdown": {
				"content": "` + msg + `"
			}}`)
		req, err := http.NewRequest("POST", "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+noticeToken, data)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bodyText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("发送企业通知：%s\n", bodyText)
	}
	//清除今日签到数据
	ClearData(cookie)
}
