package main

import (
	"fmt"
	"os"

	"github.com/Anatolij-Grigorjev/tele-go-chi/telegram"
)

func main() {

	tgBotToken := os.Getenv("BOT_TOKEN")

	botClient, err := telegram.NewTgClient(tgBotToken)
	exitOnError(err)

	tgUpdates, err := botClient.OpenUpdatesChannel()
	exitOnError(err)

	for update := range tgUpdates {
		if update.Message != nil {
			err := botClient.EchoMessage(update.Message)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Update came with no message")
		}
	}
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
