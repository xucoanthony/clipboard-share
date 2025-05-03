package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

// start sample : .\server.exe -p 8080
func main() {

	// read the coming arguments
	port := ":8080"
	if len(os.Args) < 2 {
		fmt.Println("[INFO] No port specified, using default port", port)
	} else {

		if os.Args[1] != "-p" {
			fmt.Println("[Error] Invalid argument, please use -p to specify the port")
			return
		}
		if len(os.Args) < 3 {
			fmt.Println("[Error] Please provide the port")
			return
		}
		if os.Args[2] == "" {
			fmt.Println("[Error] Please provide the port")
			return
		}
		port = ":" + os.Args[2]
		fmt.Println("[INFO] Using port", port)

	}

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
	_ = router.Run(port)

}
