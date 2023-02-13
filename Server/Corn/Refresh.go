package Corn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type RefreshStruct struct {
	Errcode         int    `json:"errcode"`
	Msg             string `json:"msg"`
	Vuserid         int    `json:"vuserid"`
	Vusession       string `json:"vusession"`
	NextRefreshTime int    `json:"next_refresh_time"`
	AccessToken     string `json:"access_token"`
	Head            string `json:"head"`
	Nick            string `json:"nick"`
}

func Refresh(cookie string) (bool, string, string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://access.video.qq.com/user/auth_refresh?vappid=11059694&vsecret=fdf61a6be0aad57132bc5cdf78ac30145b6cd2c1470b0cfe&type=qq&g_tk=&g_vstk=907466502&g_actk=642025383&callback=jQuery19109260084763662189_2676280956094&_=2676280956095", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "access.video.qq.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", cookie)
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://film.qq.com/vip/my/")
	req.Header.Set("sec-ch-ua", `"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "script")
	req.Header.Set("sec-fetch-mode", "no-cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Cookies())

	//傻逼腾讯返回的不是标准json
	temp1 := strings.Index(string(bodyText), "{")
	bodyText = bodyText[temp1:]
	temp2 := strings.Index(string(bodyText), ")")
	bodyText = bodyText[:temp2]

	log.Printf("%s\n", bodyText)

	var refreshStruct RefreshStruct
	err = json.Unmarshal(bodyText, &refreshStruct)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		//log.Printf("Cookie未失效")
		log.Printf("Vusession：%v\n", refreshStruct.Vusession)
		log.Printf("AccessToken：%v\n", refreshStruct.AccessToken)
		return true, refreshStruct.Vusession, refreshStruct.AccessToken
	} else if resp.StatusCode == 401 {
		//log.Printf("Cookie已失效")
		return false, "", ""
	}
	return true, refreshStruct.Vusession, refreshStruct.AccessToken
}
