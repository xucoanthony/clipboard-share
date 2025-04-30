package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	println("start")

	// Set the Gin mode to release mode
	gin.SetMode(gin.ReleaseMode)

	var text string
	text = "This is a default clipboard content from clipboard-share server"

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		// Hello World
		c.JSON(200, gin.H{
			"message": "You've successfully created the Clipboard-Share Server",
		})
	})

	// Set the clipboard content
	router.POST("/setClipboard", func(c *gin.Context) {
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

	// Get the clipboard content
	router.GET("/getClipboard", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message":   "Clipboard content retrieved successfully",
			"clipboard": text,
		})

		return
	})

	//Start the server on port 8080, you can change it to any port you want like below:
	//_ = router.Run(":8080")
	_ = router.Run()

}
