package Corn

import (
	"TencentVideoCheck/Server/Setting"
	"database/sql"
	"fmt"
	"github.com/robfig/cron"
	"log"
	"strconv"
	"time"
)

func CronTask() {

	// 每5秒执行一次
	// */5 * * * * ?

	// 每天23：30执行一次
	// 0 30 23 * * ?

	// 每天中午12：00执行一次
	// 0 0 12 * * ?

	refreshCron := cron.New()
	checkCron := cron.New()
	sendNotice := cron.New()

	//检查Cookie是否失效
	err := refreshCron.AddFunc("0 0 12 * * ?", func() {
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

		rows, err := db.Query("SELECT * FROM user")
		if err != nil {
			log.Fatalf("查询数据库出错：", err)
			return
		}

		for rows.Next() {
			var cookie string
			var barrage int
			var check int
			var download int
			var giving int
			var watch int
			var obtained string
			var userInfo string
			var notice string
			var noticeToken string
			var addTime int64
			err = rows.Scan(&cookie, &barrage, &check, &download, &giving, &watch, &obtained, &userInfo, &notice, &noticeToken, &addTime)
			if err != nil {
				log.Fatalf("遍历数据库出错：", err)
				return
			}
			//查询打印结果集
			//fmt.Println(cookie)
			_, _, _ = Refresh(cookie)
			// 2505600 30天
			if time.Now().Unix()-addTime < 2505600 {
				log.Printf("Cookie未失效")
			} else {
				msg := "<font color=\\\"warning\\\">腾讯视频签到通知</font>\n" + "> Cookie状态已失效，请重新提交Cookie"
				Notice(cookie, notice, noticeToken, msg)
				log.Printf("Cookie已失效,正在删除Cookie")
				deleteDB, err := db.Prepare("DELETE FROM `user` WHERE `Cookie`=?")
				if err != nil {
					fmt.Println(err)
				}
				_, err = deleteDB.Exec(cookie)
				if err != nil {
					fmt.Printf("删除数据出错：%v\n", err)
				}
			}
		}
	})
	if err != nil {
		log.Printf("检查Cookie定时任务出错：%v\n", err)
		return
	}

	//自动领取已完成任务产生的V力值
	checkCron.AddFunc("0 30 23 * * ?", func() {
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

		rows, err := db.Query("SELECT * FROM user")
		if err != nil {
			log.Fatalf("查询数据库出错：", err)
			return
		}

		for rows.Next() {
			var cookie string
			var barrage int
			var check int
			var download int
			var giving int
			var watch int
			var obtained string
			var userInfo string
			var notice string
			var noticeToken string
			var addTime int64
			err = rows.Scan(&cookie, &barrage, &check, &download, &giving, &watch, &obtained, &userInfo, &notice, &noticeToken, &addTime)
			if err != nil {
				log.Fatalf("遍历数据库出错：", err)
				return
			}
			//查询打印结果集
			//fmt.Println(cookie)

			//先刷新vusession有效期
			_, vusession, accessToken := Refresh(cookie)
			//弹幕签到
			BarrageCheck(cookie, vusession, accessToken)
			//签到
			Check(cookie, vusession, accessToken)
			//下载签到
			DownloadCheck(cookie, vusession, accessToken)
			//赠送签到
			GivingCheck(cookie, vusession, accessToken)
			//观看60分钟签到
			WatchCheck(cookie, vusession, accessToken)
			//本月获得V力值
			Obtained(cookie, vusession, accessToken)
			//用户信息
			UserInfo(cookie, vusession, accessToken)
		}
	})
	//发送通知
	sendNotice.AddFunc("0 35 23 * * ?", func() {
		//n 运行日志：\n' + resultContent + '\n 会员信息查询日志: \n > ' + vip_info
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

		rows, err := db.Query("SELECT * FROM user")
		if err != nil {
			log.Fatalf("查询数据库出错：", err)
			return
		}

		for rows.Next() {
			var cookie string
			var barrage int
			var check int
			var download int
			var giving int
			var watch int
			var obtained string
			var userInfo string
			var notice string
			var noticeToken string
			var addTime int64
			err = rows.Scan(&cookie, &barrage, &check, &download, &giving, &watch, &obtained, &userInfo, &notice, &noticeToken, &addTime)
			if err != nil {
				log.Fatalf("遍历数据库出错：", err)
				return
			}
			//查询打印结果集
			//fmt.Println(cookie)

			msg := "<font color=\\\"warning\\\">腾讯视频签到通知</font>\n" + "> 用户信息：" + userInfo + "\n >" + obtained + "\n" +
				"> 今日共获得" + strconv.Itoa((barrage + check + download + giving + watch)) + "点V力值\n\n > 任务详情：\n" +
				"发送弹幕任务获得：" + strconv.Itoa(barrage) + "点V力值\n" +
				"签到任务获得：" + strconv.Itoa(check) + "点V力值\n" +
				"下载任务获得：" + strconv.Itoa(download) + "点V力值\n" +
				"赠送任务获得：" + strconv.Itoa(giving) + "点V力值\n" +
				"观看60分钟任务获得：" + strconv.Itoa(watch) + "点V力值\n"
			Notice(cookie, notice, noticeToken, msg)

		}
	})

	refreshCron.Start()
	checkCron.Start()
	sendNotice.Start()
	select {}
}
