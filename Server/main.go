package main

import (
	"TencentVideoCheck/Server/API"
	"TencentVideoCheck/Server/Corn"
	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()
	api.POST("/push", func(c *gin.Context) {
		cookies := c.PostForm("cookie")
		notice := c.PostForm("notice")
		noticeToken := c.PostForm("noticetoken")
		if cookies == "" {
			c.JSON(200, gin.H{
				"msg": "Cookie不能为空",
			})
		} else if notice == "" {
			c.JSON(200, gin.H{
				"msg": "通知类型不能为空",
			})
		}
		if cookies != "" && notice != "" {
			API.Push(cookies, notice, noticeToken)
			c.JSON(200, gin.H{
				"msg": "添加成功",
			})
		}

	})

	api.LoadHTMLGlob("./templates/*")
	api.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"title": "腾讯视频自动签到",
		})
	})
	go Corn.CronTask()
	api.Run(":8765")
}
