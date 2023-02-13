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
)

type ObtainedStruct struct {
	LscoreMonth            int    `json:"lscore_month"`
	LscoreMonthLimit       int    `json:"lscore_month_limit"`
	Msg                    string `json:"msg"`
	RankDesc               string `json:"rank_desc"`
	Ret                    int    `json:"ret"`
	UngiveCscore           int    `json:"ungive_cscore"`
	UngiveLscore           int    `json:"ungive_lscore"`
	WillExpireUngiveCscore int    `json:"will_expire_ungive_cscore"`
	WillExpireUngiveLscore int    `json:"will_expire_ungive_lscore"`
}

func Obtained(cookie string, vusession string, accessToken string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://vip.video.qq.com/fcgi-bin/comm_cgi?name=spp_vscore_user_mashup&type=2&otype=xjson", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "vip.video.qq.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("cookie", cookie+";vqq_vusession="+vusession+";"+"vqq_access_token="+accessToken+";")
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
	//fmt.Printf("%s\n", bodyText)

	var obtainedStruct ObtainedStruct
	err = json.Unmarshal(bodyText, &obtainedStruct)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		log.Printf("本月已获得%v点V力值，%v", obtainedStruct.LscoreMonth, obtainedStruct.RankDesc)
		obtained := "本月已获得" + strconv.Itoa(obtainedStruct.LscoreMonth) + "点V力值，" + obtainedStruct.RankDesc
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

		insertDB, err := db.Prepare("UPDATE `user` SET `Obtained`=? WHERE `Cookie`=?")
		if err != nil {
			fmt.Println(err)
		}
		_, err = insertDB.Exec(obtained, cookie)
		if err != nil {
			fmt.Printf("修改数据出错：%v\n", err)
		}
	}
}
