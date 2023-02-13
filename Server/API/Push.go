package API

import (
	"TencentVideoCheck/Server/Setting"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func Push(cookie string, notice string, noticeToken string) {

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

	//INSERT INTO `user` (`Cookie`, `Barrage`, `Check`, `Download`, `Giving`, `Obtained`, `UserInfo`, `Notice`, `NoticeToken`) VALUES ('1', NULL, NULL, NULL, NULL, NULL, NULL, '2', '2');
	insertDB, err := db.Prepare("INSERT INTO `user` (`Cookie`, `Barrage`, `Check`, `Download`, `Giving`,`Watch`, `Obtained`, `UserInfo`, `Notice`, `NoticeToken`, `AddTime`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		fmt.Println(err)
	}
	_, err = insertDB.Exec(cookie, 0, 0, 0, 0, 0, "0", "无", notice, noticeToken, time.Now().Unix()) //time.Now().Unix() 获取当前时间戳
	if err != nil {
		fmt.Printf("插入数据出错：%v\n", err)
	}

}
