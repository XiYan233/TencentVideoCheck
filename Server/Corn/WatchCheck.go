package Corn

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type WatchCheckStruct struct {
	Ret   int    `json:"ret"`
	Msg   string `json:"msg"`
	Score int    `json:"score"`
}

func WatchCheck(cookie string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://vip.video.qq.com/fcgi-bin/comm_cgi?name=spp_MissionFaHuo&cmd=4&task_id=1&_=1582997048625&callback=观看60分钟签到", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "vip.video.qq.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("cookie", cookie)
	req.Header.Set("sec-ch-ua", `".Not/A)Brand";v="99", "Google Chrome";v="103", "Chromium";v="103"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//傻逼腾讯返回的不是标准json
	temp1 := strings.Index(string(bodyText), "{")
	bodyText = bodyText[temp1:]
	temp2 := strings.Index(string(bodyText), ")")
	bodyText = bodyText[:temp2]

	//fmt.Printf("%s\n", bodyText)

	var givingCheckStruct GivingCheckStruct
	err = json.Unmarshal(bodyText, &givingCheckStruct)
	if err != nil {
		log.Fatal(err)
	}

	if givingCheckStruct.Ret == -2003 {
		log.Printf("观看60分钟签到未完成，Ret[%v]\n", givingCheckStruct.Ret)
	} else if givingCheckStruct.Ret == 0 {
		log.Printf("观看60分钟签到成功，获得了%v点V力值\n", givingCheckStruct.Score)
	} else if givingCheckStruct.Ret == -2002 {
		log.Printf("重复领取，Ret[%v]\n", givingCheckStruct.Ret)
	}
}
