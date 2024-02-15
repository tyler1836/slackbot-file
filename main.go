package main

import (
	"fmt"

	"github.com/slack-go/slack"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
	}
	botToken := viper.GetString("SLACK_BOT_TOKEN")
	if botToken == "" {
		fmt.Println("Token not set")
	}
	channelId := viper.GetString("CHANNEL_ID")
	if channelId == "" {
		fmt.Println("App token not set")
	}

	api := slack.New(botToken)
	channelArr := []string{channelId}
	fileArr := []string{"zip.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Println("Error uploading file")
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}
