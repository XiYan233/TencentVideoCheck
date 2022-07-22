package Corn

import (
	"TencentVideoCheck/Server/Config"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type BarrageCheckStruct struct {
	Ret   int    `json:"ret"`
	Msg   string `json:"msg"`
	Score int    `json:"score"`
}

func BarrageCheck(cookie string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://vip.video.qq.com/fcgi-bin/comm_cgi?name=spp_MissionFaHuo&cmd=4&task_id=3&_=1582368319252&callback=弹幕签到请求", nil)
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

	var barrageCheckStruct BarrageCheckStruct
	err = json.Unmarshal(bodyText, &barrageCheckStruct)
	if err != nil {
		log.Fatal(err)
	}

	if barrageCheckStruct.Ret == -2003 {
		log.Printf("弹幕签到未完成或重复领取，Ret[%v]\n", barrageCheckStruct.Ret)
	} else if barrageCheckStruct.Ret == 0 {
		log.Printf("弹幕签到成功，获得了%v点V力值\n", barrageCheckStruct.Score)

		dsn := Config.GetDsn()
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

		insertDB, err := db.Prepare("UPDATE `user` SET `Barrage`=? WHERE `Cookie`=?")
		if err != nil {
			fmt.Println(err)
		}
		_, err = insertDB.Exec(barrageCheckStruct.Score, cookie)
		if err != nil {
			fmt.Printf("修改数据出错：%v\n", err)
		}

	}
}
