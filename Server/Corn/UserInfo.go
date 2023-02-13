package Corn

import (
	"TencentVideoCheck/Server/Setting"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type UserInfoStruct struct {
	BeginTime string `json:"beginTime"`
	EndTime   string `json:"endTime"`
	Level     int    `json:"level"`
	Score     int    `json:"score"`
}

func UserInfo(cookie string, vusession string, accessToken string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://vip.video.qq.com/fcgi-bin/comm_cgi?name=payvip&cmd=1&otype=json&getannual=1&geticon=1&getsvip=1&callback=jQuery191008113049860183752_1658307109042&uin=0&t=1&getadpass=0&g_tk=&g_vstk=212673112&g_actk=1469559637&_=1658307109048", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "vip.video.qq.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("cookie", cookie+";vqq_vusession="+vusession+";"+"vqq_access_token="+accessToken+";")
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
	//fmt.Printf("%s\n", bodyText)

	//傻逼腾讯返回的不是标准json
	temp1 := strings.Index(string(bodyText), "{")
	bodyText = bodyText[temp1:]
	temp2 := strings.Index(string(bodyText), ")")
	bodyText = bodyText[:temp2]

	var userInfoStruct UserInfoStruct
	err = json.Unmarshal(bodyText, &userInfoStruct)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		log.Printf("会员开通时间%v，到期时间%v，当前会员等级为%v级，总共获得了%v点V力值", userInfoStruct.BeginTime, userInfoStruct.EndTime, userInfoStruct.Level, userInfoStruct.Score)

		userInfo := "会员开通时间" + userInfoStruct.BeginTime + "，到期时间" + userInfoStruct.EndTime + "，当前会员等级为" + strconv.Itoa(userInfoStruct.Level) + "级，总共获得了" + strconv.Itoa(userInfoStruct.Score) + "点V力值"
		dsn := Setting.GetDsn()
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
		err = db.Ping()
		if err != nil {
			fmt.Printf("连接数据库出错：%v\n", err)
			return
		}

		insertDB, err := db.Prepare("UPDATE `user` SET `UserInfo`=? WHERE `Cookie`=?")
		if err != nil {
			fmt.Println(err)
		}
		_, err = insertDB.Exec(userInfo, cookie)
		if err != nil {
			fmt.Printf("修改数据出错：%v\n", err)
		}
	}
}
