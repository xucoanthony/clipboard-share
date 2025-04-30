package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	println("start")

	// 设置 Gin 的运行模式为 Release 模式
	gin.SetMode(gin.ReleaseMode)

	var text string
	text = "This is a default clipboard content from clipboard-share server"

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		// 处理请求 获取参数
		c.JSON(200, gin.H{
			"message": "You've successfully created the Clipboard-Share Server",
		})
	})

	router.POST("/setClipboard", func(c *gin.Context) {
		// 处理请求 获取参数
		var postData map[string]interface{}
		if err := c.ShouldBind(&postData); err != nil {
			return
		}

		if postData["clipboard"] == nil {
			c.JSON(400, gin.H{
				"message": "Clipboard content is required, check if you lost the key 'clipboard'",
			})
			return
		}

		text = postData["clipboard"].(string)

		c.JSON(200, gin.H{
			"message":   "Clipboard content set successfully",
			"clipboard": text,
		})

		return
	})

	router.GET("/getClipboard", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message":   "Clipboard content retrieved successfully",
			"clipboard": text,
		})

		return
	})

	_ = router.Run() // 监听并在 0.0.0.0:8080 上启动服务

}
