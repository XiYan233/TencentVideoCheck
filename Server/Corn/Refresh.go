package Corn

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Refresh(cookie string) bool {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://access.video.qq.com/user/auth_refresh?vappid=11059694&vsecret=fdf61a6be0aad57132bc5cdf78ac30145b6cd2c1470b0cfe&type=qq&g_tk=&g_vstk=1959117529&g_actk=1469559637&callback=jQuery191040005793960664526_1658814553950&_=1658814553951", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "access.video.qq.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("cookie", cookie)
	req.Header.Set("referer", "https://film.qq.com/vip/my/")
	req.Header.Set("sec-ch-ua", `".Not/A)Brand";v="99", "Google Chrome";v="103", "Chromium";v="103"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "script")
	req.Header.Set("sec-fetch-mode", "no-cors")
	req.Header.Set("sec-fetch-site", "same-site")
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

	log.Printf("%s\n", bodyText)
	if resp.StatusCode == 200 {
		//log.Printf("Cookie未失效")
		return true
	} else if resp.StatusCode == 401 {
		//log.Printf("Cookie已失效")
		return false
	}
	return true
}
