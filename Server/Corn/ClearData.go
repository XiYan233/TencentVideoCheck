package Corn

import (
	"TencentVideoCheck/Server/Config"
	"database/sql"
	"fmt"
)

func ClearData(cookie string) {
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

	insertDB, err := db.Prepare("UPDATE `user` SET `Barrage` = ?, `Check` = ?, `Download` = ?, `Giving` = ?, `Watch` = ? WHERE `Cookie`=?")
	if err != nil {
		fmt.Println(err)
	}
	_, err = insertDB.Exec(0, 0, 0, 0, 0, cookie)
	if err != nil {
		fmt.Printf("修改数据出错：%v\n", err)
	}
}
