package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func handleSlack(text string) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	apiToken := os.Getenv("SLACK_API_TOKEN")

	channelID := os.Getenv("SLACK_CHANNEL_ID") // 送信するチャンネルのIDを指定してください。

	api := slack.New(apiToken)
	_, _, err = api.PostMessage(channelID, slack.MsgOptionText(text, false))

	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	} else {
		fmt.Println("Message sent successfully")
	}
}
