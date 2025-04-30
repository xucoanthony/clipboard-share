package main

import (
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/go-resty/resty/v2"
	"os"
	"time"
)

func main() {
	// read the configuration file config.json
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println("[Error] Failed to open config.json: ", err)
		return
	}
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			fmt.Println("[Error] Failed to close config.json: ", err)
		}
	}(configFile)

	var config Config
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		fmt.Println("[Error] Failed to decode config.json: ", err)
		return
	}

	// Apply the configuration
	deviceName := config.DeviceName
	updateFrequency := config.UpdateFrequencySecond
	remoteServerHost := config.RemoteServer.Host
	remoteServerPort := config.RemoteServer.Port
	remoteUrl := fmt.Sprintf("http://%s:%d", remoteServerHost, remoteServerPort)

	fmt.Println("[Info] Device Name:", deviceName)
	fmt.Println("[Info] Update Frequency:", updateFrequency, "seconds")
	fmt.Println("[Info] Remote Server:", remoteUrl)

	var client *resty.Client
	client = resty.New()

	oldLocalClipboard := ""

	for {

		newLocalClipboard := getLocalClipboard()

		// 先判断本地的剪贴板内容是否相同
		if newLocalClipboard != oldLocalClipboard {
			// 如果不同，则更新远程剪贴板
			setRemoteClipboard(client, remoteUrl, newLocalClipboard)
			fmt.Println("[Info] Updated remote clipboard from local clipboard")
		} else {
			// 如果相同，则获取远程剪贴板内容进行判断
			remoteClipboard := getRemoteClipboard(client, remoteUrl)
			if remoteClipboard != newLocalClipboard {
				// 如果远程剪贴板内容不同，则更新本地剪贴板
				setLocalClipboard(remoteClipboard)
				fmt.Println("[Info] Updated local clipboard from remote clipboard")
			} else {
				// 如果相同，则不做任何操作
				fmt.Println("[Info] No changes in clipboard")
			}
		}

		oldLocalClipboard = newLocalClipboard
		time.Sleep(time.Duration(updateFrequency) * time.Second)
	}
}

// Get clipboard content
func getLocalClipboard() string {

	text := ""
	// Check if clipboard is supported
	if !clipboard.Unsupported {
		_text, err := clipboard.ReadAll()
		if err != nil {
			fmt.Println("[Error] ", err)
		}
		text = _text
	} else {
		fmt.Println("[Error] Your system does not support clipboard operations")
	}

	return text

}

// Set clipboard content
func setLocalClipboard(text string) {
	if !clipboard.Unsupported {
		err := clipboard.WriteAll(text)
		if err != nil {
			fmt.Println("[Error] ", err)
		}
	} else {
		fmt.Println("[Error] Your system does not support clipboard operations")
	}
}

// Get remote clipboard content
func getRemoteClipboard(client *resty.Client, baseUrl string) string {
	resp, err := client.R().Get(baseUrl + "/getClipboard")
	if err != nil {
		fmt.Println("[Error] ", err)
		return ""
	}
	if resp.StatusCode() != 200 {
		fmt.Println("[Error] ", resp.Status())
		return ""
	}
	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		fmt.Println("[Error] ", err)
		return ""
	}
	return result["clipboard"].(string)
}

// Set remote clipboard content
func setRemoteClipboard(client *resty.Client, baseUrl string, text string) {
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{"clipboard": text}).
		Post(baseUrl + "/setClipboard")
	if err != nil {
		fmt.Println("[Error] ", err)
		return
	}
	if resp.StatusCode() != 200 {
		fmt.Println("[Error] ", resp.Status())
		return
	}
}
